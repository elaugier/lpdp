package controllers

import (
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"

	"github.com/elaugier/lpdp/pkg/config"
	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//LoginController ...
type LoginController struct{}

//Get ...
func (u LoginController) Get(c *gin.Context) {
	logger := logs.GetInstance()
	logger.Println("Enter Get from LoginController")
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u LoginController) List(c *gin.Context) {
	logger := logs.GetInstance()
	logger.Println("Enter List from LoginController")
	c.String(http.StatusOK, "Working!")
}

//Login ...
func (u LoginController) Login(c *gin.Context) {
	logger := logs.GetInstance()
	logger.Println("Enter Login from LoginController")
	var json models.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db.GetInstance()
	count := 0
	db.Where("Email = ? AND Password = ?", json.User, json.Password).Count(&count)
	if count == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	var user models.User
	if err := db.Where("Email = ? AND Password = ?", json.User, json.Password).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "bad user or bad password"})
		return
	}

	// 'iat' => $issuedAt,         // Issued at: time when the token was generated
	// 'jti' => $tokenId,          // Json Token Id: an unique identifier for the token
	// 'iss' => $serverName,       // Issuer
	// 'nbf' => $notBefore,        // Not before
	// 'exp' => $expire,           // Expire
	// 'data' => [
	// 	'id' => $element->id,
	// 	'email' => $element->email,
	// 	'pseudo' => $element->pseudo,
	// 	'accountType' => $accountTypes[$element->accountType],
	// 	'memberRole' => $memberRoles[$element->memberRole],
	// ]

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"jti": uuid.New().String(),
	})
	configuration, err := config.Get()
	if err != nil {
		logger.Fatal(err)
	}
	tokenString, err := token.SignedString(configuration.GetString("jwt:secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error on token genration"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in", "token": tokenString})
}

//Logout ...
func (u LoginController) Logout(c *gin.Context) {
	logger := logs.GetInstance()
	logger.Println("Enter Logout from LoginController")
	c.String(http.StatusOK, "Working!")
}
