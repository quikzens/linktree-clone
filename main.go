package main

import (
	"linktree-clone/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	SetMiddlewares(router)
	SetRoutes(router)

	log.Fatal(router.Run(config.ServerAddress))
}
