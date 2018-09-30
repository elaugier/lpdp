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
		c.Next()
		logger.Println("Exit from Authentication middleware")
	}
}
