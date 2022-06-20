package user

import (
	"context"
	"fmt"
	"linktree-clone/db"
	"linktree-clone/util"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func Login(c *gin.Context) {
	providerParam := c.Param("provider")

	provider, err := gomniauth.Provider(providerParam)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error when trying to get provider %s: %s", provider, err))
		c.Abort()
		return
	}

	loginUrl, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error when trying to GetBeginAuthURL for %s:%s", provider, err))
		c.Abort()
		return
	}

	c.Status(http.StatusTemporaryRedirect)
	c.Header("Location", loginUrl)
}

func Callback(c *gin.Context) {
	providerParam := c.Param("provider")

	provider, err := gomniauth.Provider(providerParam)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error when trying to get provider %s: %s", provider, err))
		c.Abort()
		return
	}

	creds, err := provider.CompleteAuth(objx.MustFromURLQuery(c.Request.URL.RawQuery))
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error when trying to complete auth for %s: %s", provider, err))
		c.Abort()
		return
	}

	userData, err := provider.GetUser(creds)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error when trying to get user from %s: %s", provider, err))
		c.Abort()
		return
	}

	// if it's a new user (check by email address):
	// create new document in db to keep it's links
	var userDB user
	userPayload := util.UserPayload{
		ID:        uuid.NewString(),
		Name:      userData.Name(),
		Email:     userData.Email(),
		AvatarUrl: userData.AvatarURL(),
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(24 * time.Hour),
	}

	err = db.UserColl.FindOne(context.TODO(), bson.M{"email": userData.Email()}).Decode(&userDB)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			newUsername := strings.ReplaceAll(strings.ToLower(userData.Name()), " ", "_") + util.RandomString(10)
			_, err := db.UserColl.InsertOne(context.TODO(), user{
				ID:        uuid.NewString(),
				Username:  newUsername,
				Email:     userData.Email(),
				Links:     []string{},
				AvatarURL: userData.AvatarURL(),
				CreatedAt: time.Now().Unix(),
				UpdatedAt: 0,
			})
			if err != nil {
				util.SendServerError(c, err)
				return
			}
			userPayload.Username = newUsername
		} else {
			util.SendServerError(c, err)
			return
		}
	} else {
		userPayload.Username = userDB.Username
	}

	token, _, err := util.CreateToken(&userPayload)
	if err != nil {
		util.SendServerError(c, err)
		return
	}

	c.SetCookie("auth", token, 60*60*24, "/", "", false, true)
	c.Status(http.StatusTemporaryRedirect)
	c.Header("Location", "/")
}

func Logout(c *gin.Context) {
	c.SetCookie("auth", "", -1, "/", "", false, true)
	c.Status(http.StatusTemporaryRedirect)
	c.Header("Location", "/")
}

// CheckAuth is validate a user session from it's token and return it's authenticated user data
func CheckAuth(c *gin.Context) {
	payload, _ := c.Get("user")
	userPayload, _ := payload.(*util.UserPayload)

	util.SendSuccess(c, gin.H{
		"username":        userPayload.Username,
		"user_name":       userPayload.Name,
		"user_email":      userPayload.Email,
		"user_avatar_url": userPayload.AvatarUrl,
	})
}

func GetUser(c *gin.Context) {
	username := c.Param("username")

	var user user
	err := db.UserColl.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		util.SendServerError(c, err)
		return
	}

	userLinks := []linkResponse{}
	for _, linkID := range user.Links {
		var link linkResponse
		err := db.LinkColl.FindOne(context.TODO(), bson.M{"id": linkID}).Decode(&link)
		if err != nil {
			util.SendServerError(c, err)
			return
		}
		userLinks = append(userLinks, link)
	}

	util.SendSuccess(c, userResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Links:     userLinks,
		AvatarURL: user.AvatarURL,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}
