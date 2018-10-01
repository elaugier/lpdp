package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/logs"

	"github.com/google/uuid"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
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
		c.JSON(400, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}
	var alertaction models.AlertAction
	conn := db.GetInstance()
	if err = conn.Where("id = ?", id).First(&alertaction).Error; err != nil {
		log.Println("alert not found.")
		c.JSON(404, gin.H{
			"msg": "alert not found.",
		})
	} else {
		log.Println("alert returned")
		c.JSON(200, alertaction)
	}
}

//List ...
func (u AlertActionsController) List(c *gin.Context) {
	log := logs.GetInstance()
	var alertactions []models.AlertAction
	conn := db.GetInstance()
	if err := conn.Find(&alertactions).Error; err != nil {
		log.Println("action not found.")
		c.JSON(404, gin.H{
			"msg": "action not found.",
		})
	} else {
		log.Println("action returned")
		c.JSON(200, alertactions)
	}
}

//Add ...
func (u AlertActionsController) Add(c *gin.Context) {
	var json models.AlertAction
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u AlertActionsController) Modify(c *gin.Context) {
	var json models.AlertAction
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u AlertActionsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
