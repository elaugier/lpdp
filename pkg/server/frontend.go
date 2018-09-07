package server

import (
	"log"
	"net/http"
	"path"
	"path/filepath"

	"github.com/elaugier/lpdp/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

//FrontendRouter ...
func FrontendRouter(logger *log.Logger) http.Handler {
	logger.Println("Create new frontend router")
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.Identification(logger))
	router.Use(middlewares.RequestID(logger))
	//router.Static("/", "./www/")
	router.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		logger.Printf("dir => %s | file => %s", dir, file)
		ext := filepath.Ext(file)
		logger.Printf("ext => %s", ext)
		if file == "" || ext == "" {
			c.File("./www/index.html")
		} else {
			// strings.Split(file, "?")
			c.File("./www" + path.Join(dir, file))
		}

	})
	return router
}
