package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//RequestID ...
func RequestID(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Println("Enter in RequestID middleware")
		requestID, _ := uuid.NewUUID()
		c.Request.Header.Set("X-Request-ID", requestID.String())
		c.Header("X-Request-ID", requestID.String())
		c.Next()
		logger.Println("Exit from RequestID middleware")
	}
}
