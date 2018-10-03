package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//AlertActionsController ...
type AlertActionsController struct{}

//Get ...
func (u AlertActionsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var alertaction models.AlertAction

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&alertaction).Error; err != nil {
		log.Println("alert action not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "alert action not found.",
		})
	} else {
		log.Println("alert action returned")
		c.JSON(http.StatusOK, alertaction)
	}

}

//List ...
func (u AlertActionsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var alertactions []models.AlertAction

	conn := db.GetInstance()
	if err := conn.Find(&alertactions).Error; err != nil {
		log.Println("alert action not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "alert action not found.",
		})
	} else {
		log.Println("alert action returned")
		c.JSON(http.StatusOK, alertactions)
	}

}

//Add ...
func (u AlertActionsController) Add(c *gin.Context) {

	var alertactions models.AlertAction

	if err := c.ShouldBindJSON(&alertactions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&alertactions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&alertactions)

	c.JSON(http.StatusCreated, alertactions)

}

//Modify ...
func (u AlertActionsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var alertactions models.AlertAction

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&alertactions).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("alert action not found.")
	}

	if err := c.ShouldBindJSON(&alertactions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&alertactions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&alertactions)

	c.JSON(http.StatusOK, alertactions)
}

//Remove ...
func (u AlertActionsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var alertaction models.AlertAction

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&alertaction).Error; err != nil {
		log.Println("alert action not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "alert action not found.",
		})
	} else {
		log.Println("alert action returned")
		c.JSON(http.StatusOK, alertaction)
	}

}
