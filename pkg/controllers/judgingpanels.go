package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//JudgingPanelsController ...
type JudgingPanelsController struct{}

//Get ...
func (u JudgingPanelsController) Get(c *gin.Context) {
	log := logs.GetInstance()
	log.Println("try to retieve 'id' in url path")
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(400, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}
	var i models.JudgingPanel
	conn := db.GetInstance()
	if err = conn.Where("id = ?", id).First(&i).Error; err != nil {
		log.Println("judging panel not found.")
		c.JSON(404, gin.H{
			"msg": "judging panel not found.",
		})
	} else {
		log.Println("judging panel returned")
		c.JSON(200, i)
	}
}

//List ...
func (u JudgingPanelsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u JudgingPanelsController) Add(c *gin.Context) {
	var json models.JudgingPanel
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u JudgingPanelsController) Modify(c *gin.Context) {
	var json models.JudgingPanel
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u JudgingPanelsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
