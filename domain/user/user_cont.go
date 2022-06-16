package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/gomniauth"
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
