package main

import (
	"linktree-clone/domain/link"
	"linktree-clone/domain/user"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.GET("/login/:provider", user.Login)
	auth.GET("/callback/:provider", user.Callback)
	auth.GET("/logout", user.Logout)

	api := router.Group("/api")

	api.GET("/user/auth", user.VerifyAuth, user.CheckAuth)
	api.GET("/user/:username", user.GetUser)

	api.POST("/link", user.VerifyAuth, link.CreateLink)
}
