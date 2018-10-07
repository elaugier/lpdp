package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//UsersController ...
type UsersController struct{}

//Get ...
func (u UsersController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var user models.User

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&user).Error; err != nil {
		log.Println("user not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "user not found.",
		})
	} else {
		log.Println("user returned")
		c.JSON(http.StatusOK, user)
	}

}

//List ...
func (u UsersController) List(c *gin.Context) {

	log := logs.GetInstance()

	var users []models.User

	conn := db.GetInstance()
	if err := conn.Find(&users).Error; err != nil {
		log.Println("user not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "user not found.",
		})
	} else {
		log.Println("user returned")
		c.JSON(http.StatusOK, users)
	}

}

//Add ...
func (u UsersController) Add(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&user)

	c.JSON(http.StatusCreated, user)

}

//Modify ...
func (u UsersController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var user models.User

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("user not found.")
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&user)

	c.JSON(http.StatusOK, user)
}

//Remove ...
func (u UsersController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var user models.User

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&user).Error; err != nil {
		log.Println("user not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "user not found.",
		})
	} else {
		log.Println("user returned")
		c.JSON(http.StatusOK, user)
	}

}
