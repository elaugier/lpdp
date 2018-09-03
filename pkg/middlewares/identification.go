package middlewares

import (
	"log"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

//Identification ...
func Identification(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionCookie := false
		logger.Println("Enter in Identification middleware")
		cookie, err := c.Request.Cookie("session")
		if err != nil {
			logger.Println("Can't retrieve session cookie")
		} else {
			logger.Printf("Session cookie present : '%s' = '%s'", "session", cookie.String())
			sessionCookie = true
		}
		if !sessionCookie {
			sessionUUID, _ := uuid.NewUUID()
			logger.Printf("Generate new session cookie UUID = %s", sessionUUID.String())
			c.SetCookie("session", sessionUUID.String(), 999, "", "", false, false)
		}
		c.Next()
		logger.Println("Exit from Identification middleware")
	}
}
