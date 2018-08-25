package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//BackendRouter ...
func BackendRouter() http.Handler {
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
