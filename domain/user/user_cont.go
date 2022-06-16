package user

import (
	"fmt"
	"linktree-clone/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
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

	user, err := provider.GetUser(creds)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error when trying to get user from %s: %s", provider, err))
		c.Abort()
		return
	}

	token, _, err := util.CreateToken(&util.UserPayload{
		ID:        uuid.NewString(),
		Name:      user.Name(),
		Email:     user.Email(),
		AvatarUrl: user.AvatarURL(),
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(24 * time.Hour),
	})
	if err != nil {
		util.SendServerError(c, err)
		return
	}

	c.SetCookie("auth", token, 60*60*24, "/", "", false, true)
	c.Status(http.StatusTemporaryRedirect)
	c.Header("Location", "/")
}
