package main

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func SetMiddlewares(router *gin.Engine) {
	// Serve Frontend Client
	uiFS, _ := fs.Sub(UI, "_ui/build")
	staticServer := static.Serve("/", embedFileSystem{
		FileSystem: http.FS(uiFS),
		indexes:    true,
	})

	router.Use(staticServer)
	router.NoRoute(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet &&
			!strings.ContainsRune(c.Request.URL.Path, '.') &&
			(!strings.HasPrefix(c.Request.URL.Path, "/api/") || !strings.HasPrefix(c.Request.URL.Path, "/auth/")) {
			c.Request.URL.Path = "/"
			staticServer(c)
		}
	})
}
