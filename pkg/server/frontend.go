package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//FrontendRouter ...
func FrontendRouter(logger *log.Logger) http.Handler {
	logger.Println("Create new frontend router")
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server FrontEnd",
			},
		)
	})
	return e
}
