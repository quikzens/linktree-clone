package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	authCookieValue := objx.New(map[string]interface{}{
		"name":       user.Name(),
		"email":      user.Email(),
		"avatar_url": user.AvatarURL(),
	}).MustBase64()

	c.SetCookie("auth", authCookieValue, 60*15, "/", "", false, true)
	c.Status(http.StatusTemporaryRedirect)
	c.Header("Location", "/")
}
