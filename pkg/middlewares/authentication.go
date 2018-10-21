package middlewares

import (
	"fmt"
	"log"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/elaugier/lpdp/pkg/config"
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
		} else {
			logger.Printf("get Authorization header: '%s'", token)
			token = strings.Replace(token, "Bearer", "", -1)
			logger.Printf("'Bearer removed' : '%s'", token)
			jwtToken, err := jwt.Parse(token, func(ptoken *jwt.Token) (interface{}, error) {
				if _, ok := ptoken.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", ptoken.Header["alg"])
				}
				configuration, err := config.Get()
				if err != nil {
					logger.Fatal(err)
				}
				return []byte(configuration.GetString("jwt:secret")), nil
			})

			if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
				for i, s := range claims {
					logger.Printf("token claims : '%s' = '%s'", i, s)
				}
			} else {
				logger.Println(err)
			}
		}
		c.Next()
		logger.Println("Exit from Authentication middleware")
	}
}
