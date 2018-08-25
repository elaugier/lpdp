package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

//Identification ...
func Identification(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Println("Enter in Identification middleware")
		c.Next()
		logger.Println("Exit from Identification middleware")
	}
}
