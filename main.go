package main

import (
	"linktree-clone/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

func init() {
	// setup oauth
	gomniauth.SetSecurityKey(config.SecretKey)
	gomniauth.WithProviders(
		google.New(config.GoogleClientID, config.GoogleClientSecret, config.GoogleRedirectUrl),
	)
}

func main() {
	router := gin.Default()

	SetMiddlewares(router)
	SetRoutes(router)

	log.Fatal(router.Run(config.ServerAddress))
}
