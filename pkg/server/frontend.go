package server

import (
	"log"
	"net/http"

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
	router.Static("/", "./www/")

	return router
}
