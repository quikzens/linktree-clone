package user

import (
	"errors"

	"linktree-clone/util"

	"github.com/gin-gonic/gin"
)

func VerifyAuth(c *gin.Context) {
	token, err := c.Cookie("auth")
	if err != nil {
		util.SendUnauthorized(c, err)
		return
	}

	payload, err := util.VerifyToken(token)
	if err != nil {
		if errors.Is(err, util.ErrExpiredToken) {
			c.SetCookie("token", "", 0, "", "", true, true)
		}

		util.SendUnauthorized(c, err)
		return
	}

	c.Set("user", payload)
	c.Next()
}
