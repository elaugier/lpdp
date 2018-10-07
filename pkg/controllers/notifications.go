package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//NotificationsController ...
type NotificationsController struct{}

//Get ...
func (u NotificationsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var notification models.Notification

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&notification).Error; err != nil {
		log.Println("notification not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "notification not found.",
		})
	} else {
		log.Println("notification returned")
		c.JSON(http.StatusOK, notification)
	}

}

//List ...
func (u NotificationsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var notifications []models.Notification

	conn := db.GetInstance()
	if err := conn.Find(&notifications).Error; err != nil {
		log.Println("notification not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "notification not found.",
		})
	} else {
		log.Println("notification returned")
		c.JSON(http.StatusOK, notifications)
	}

}

//Add ...
func (u NotificationsController) Add(c *gin.Context) {

	var notification models.Notification

	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&notification); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&notification)

	c.JSON(http.StatusCreated, notification)

}

//Modify ...
func (u NotificationsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var notification models.Notification

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&notification).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("notification not found.")
	}

	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&notification); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&notification)

	c.JSON(http.StatusOK, notification)
}

//Remove ...
func (u NotificationsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var notification models.Notification

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&notification).Error; err != nil {
		log.Println("notification not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "notification not found.",
		})
	} else {
		log.Println("notification returned")
		c.JSON(http.StatusOK, notification)
	}

}
