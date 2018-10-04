package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//AlertController ...
type AlertController struct{}

//Get ...
func (u AlertController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var alert models.Alert

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&alert).Error; err != nil {
		log.Println("alert not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "alert not found.",
		})
	} else {
		log.Println("alert returned")
		c.JSON(http.StatusOK, alert)
	}

}

//List ...
func (u AlertController) List(c *gin.Context) {

	log := logs.GetInstance()

	var alert []models.Alert

	conn := db.GetInstance()
	if err := conn.Find(&alert).Error; err != nil {
		log.Println("alert not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "alert not found.",
		})
	} else {
		log.Println("alert returned")
		c.JSON(http.StatusOK, alert)
	}

}

//Add ...
func (u AlertController) Add(c *gin.Context) {

	var alert models.Alert

	if err := c.ShouldBindJSON(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&alert); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&alert)

	c.JSON(http.StatusCreated, alert)

}

//Modify ...
func (u AlertController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var alert models.Alert

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&alert).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("alert not found.")
	}

	if err := c.ShouldBindJSON(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&alert); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&alert)

	c.JSON(http.StatusOK, alert)
}

//Remove ...
func (u AlertController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var alert models.Alert

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&alert).Error; err != nil {
		log.Println("alert not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "alert not found.",
		})
	} else {
		log.Println("alert returned")
		c.JSON(http.StatusOK, alert)
	}

}