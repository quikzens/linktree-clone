package link

import (
	"context"
	"linktree-clone/db"
	"linktree-clone/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/mgo.v2/bson"
)

func CreateLink(c *gin.Context) {
	payload, _ := c.Get("user")
	userPayload, _ := payload.(*util.UserPayload)
	userEmail := userPayload.Email

	newLinkID := uuid.NewString()
	newLink := link{
		ID:        newLinkID,
		Title:     "Title",
		Url:       "Url",
		IsActive:  false,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: 0,
	}
	_, err := db.LinkColl.InsertOne(context.TODO(), newLink)
	if err != nil {
		util.SendServerError(c, err)
		return
	}

	var user userLinks
	err = db.UserColl.FindOne(context.TODO(), bson.M{"email": userEmail}).Decode(&user)
	if err != nil {
		util.SendServerError(c, err)
		return
	}

	updatedLinks := []string{newLinkID}
	updatedLinks = append(updatedLinks, user.Links...)
	_, err = db.UserColl.UpdateOne(context.TODO(), bson.M{"email": userEmail}, bson.M{
		"$set": updateUserLinks{
			Links:     updatedLinks,
			UpdatedAt: time.Now().Unix(),
		},
	})
	if err != nil {
		util.SendServerError(c, err)
		return
	}

	util.SendSuccess(c, newLink)
}
