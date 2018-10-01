package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//AlertsController ...
type AlertsController struct{}

//Get ...
func (u AlertsController) Get(c *gin.Context) {
	log := logs.GetInstance()
	log.Println("try to retieve 'id' in url path")
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(400, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}
	var alert models.Alert
	conn := db.GetInstance()
	if err = conn.Where("id = ?", id).First(&alert).Error; err != nil {
		log.Println("alert not found.")
		c.JSON(404, gin.H{
			"msg": "alert not found.",
		})
	} else {
		log.Println("alert returned")
		c.JSON(200, alert)
	}
}

//List ...
func (u AlertsController) List(c *gin.Context) {
	log := logs.GetInstance()
	var alert []models.Alert
	conn := db.GetInstance()
	if err := conn.Find(&alert).Error; err != nil {
		log.Println("alert not found.")
		c.JSON(404, gin.H{
			"msg": "alert not found.",
		})
	} else {
		log.Println("alert returned")
		c.JSON(200, alert)
	}
}

//Add ...
func (u AlertsController) Add(c *gin.Context) {
	var json models.Alert
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u AlertsController) Modify(c *gin.Context) {
	var json models.Alert
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u AlertsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
