package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/logs"

	"github.com/google/uuid"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//AchievementsController ...
type AchievementsController struct{}

//Get ...
func (u AchievementsController) Get(c *gin.Context) {
	log := logs.GetInstance()
	log.Println("try to retieve 'id' in url path")
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(400, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}
	var achievement models.Achievement
	conn := db.GetInstance()
	if err = conn.Where("id = ?", id).First(&achievement).Error; err != nil {
		log.Println("achievement not found.")
		c.JSON(404, gin.H{
			"msg": "achievement not found.",
		})
	} else {
		log.Println("achievement returned")
		c.JSON(200, achievement)
	}
}

//List ...
func (u AchievementsController) List(c *gin.Context) {
	log := logs.GetInstance()
	var achievements []models.Achievement
	conn := db.GetInstance()
	if err := conn.Find(&achievements).Error; err != nil {
		log.Println("achievement not found.")
		c.JSON(404, gin.H{
			"msg": "achievement not found.",
		})
	} else {
		log.Println("achievement returned")
		c.JSON(200, achievements)
	}
}

//Add ...
func (u AchievementsController) Add(c *gin.Context) {
	var json models.Achievement
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u AchievementsController) Modify(c *gin.Context) {
	var json models.Achievement
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u AchievementsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
