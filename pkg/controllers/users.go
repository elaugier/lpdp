package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"

	"github.com/google/uuid"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//UsersController ...
type UsersController struct {
}

//Get ...
func (u UsersController) Get(c *gin.Context) {
	log := logs.GetInstance()
	log.Println("try to retieve 'id' in url path")
	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(400, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}
	var user models.User
	conn := db.GetInstance()
	if err = conn.Where("id = ?", id).First(&user).Error; err != nil {
		log.Println("user not found.")
		c.JSON(404, gin.H{
			"msg": "user not found.",
		})
	} else {
		log.Println("user returned")
		c.JSON(200, user)
	}
}

//List ...
func (u UsersController) List(c *gin.Context) {
	log := logs.GetInstance()
	var users []models.User
	conn := db.GetInstance()
	if err := conn.Find(&users).Error; err != nil {
		log.Println("user not found.")
		c.JSON(404, gin.H{
			"msg": "user not found.",
		})
	} else {
		log.Println("user returned")
		c.JSON(200, users)
	}
}

//Add ...
func (u UsersController) Add(c *gin.Context) {
	var json models.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u UsersController) Modify(c *gin.Context) {
	var json models.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u UsersController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
