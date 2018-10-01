package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

//Authentication ...
func Authentication(logger *log.Logger) gin.HandlerFunc {
	logger.Println("Initialization of Authentication middleware")
	return func(c *gin.Context) {
		logger.Println("Enter Authentication middleware")
		token := c.Request.Header.Get("Authorization")
		if len(token) == 0 {
			logger.Println("No Authentication")
		}
		c.Next()
		logger.Println("Exit from Authentication middleware")
	}
}
