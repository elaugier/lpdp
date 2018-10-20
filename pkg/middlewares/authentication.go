package middlewares

import (
	"fmt"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
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
			jwtToken, err := jwt.Parse(token, func(ptoken *jwt.Token) (interface{}, error) {
				if _, ok := ptoken.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", ptoken.Header["alg"])
				}
				return []byte("ZYF8La5SLTTXemG4Gdrxbb9MwtEdJ4jjL3Dg5mh5i3mwJ2VdG9xY4S94udE3kh435B6Vx2B8iqEj6yz6n5iSvJ4k3iJyVjx842eVk7ZZ88q6DK3Aw39N8Nn47j4dwSy53ffHNE4ASwEVtfbT4X87WTDJAv8bb2842529Fkk8zv6gx7p2k24DBKfUrD64JG3fTbD48m5D7Z3jmyfw8r2qB9P8FJ342sYP5397RJEC8aQD7HbrtskLKF9X5a585LCN"), nil
			})

			if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
				fmt.Println(claims["foo"], claims["nbf"])
			} else {
				fmt.Println(err)
			}
		}
		c.Next()
		logger.Println("Exit from Authentication middleware")
	}
}
