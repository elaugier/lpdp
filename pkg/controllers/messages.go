package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//MessagesController ...
type MessagesController struct{}

//Get ...
func (u MessagesController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var message models.Message

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&message).Error; err != nil {
		log.Println("message not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "message not found.",
		})
	} else {
		log.Println("message returned")
		c.JSON(http.StatusOK, message)
	}

}

//List ...
func (u MessagesController) List(c *gin.Context) {

	log := logs.GetInstance()

	var messages []models.Message

	conn := db.GetInstance()
	if err := conn.Find(&messages).Error; err != nil {
		log.Println("message not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "message not found.",
		})
	} else {
		log.Println("message returned")
		c.JSON(http.StatusOK, messages)
	}

}

//Add ...
func (u MessagesController) Add(c *gin.Context) {

	var message models.Message

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&message)

	c.JSON(http.StatusCreated, message)

}

//Modify ...
func (u MessagesController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var message models.Message

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&message).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("message not found.")
	}

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&message)

	c.JSON(http.StatusOK, message)
}

//Remove ...
func (u MessagesController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var message models.Message

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&message).Error; err != nil {
		log.Println("message not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "message not found.",
		})
	} else {
		log.Println("message returned")
		c.JSON(http.StatusOK, message)
	}

}
