package main

import (
	"linktree-clone/domain/user"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.GET("/login/:provider", user.Login)
	auth.GET("/callback/:provider", user.Callback)

	api := router.Group("/api")
	api.GET("/user/auth", user.VerifyAuth, user.CheckAuth)
}
