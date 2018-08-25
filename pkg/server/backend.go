package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//BackendRouter ...
func BackendRouter(logger *log.Logger) http.Handler {
	logger.Println("Create new backend router")
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server Backend",
			},
		)
	})

	return e
}
