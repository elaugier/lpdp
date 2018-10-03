package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//AchievementsController ...
type AchievementsController struct{}

//Get ...
func (u AchievementsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var achievement models.Achievement

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&achievement).Error; err != nil {
		log.Println("achievement not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "achievement not found.",
		})
	} else {
		log.Println("achievement returned")
		c.JSON(http.StatusOK, achievement)
	}

}

//List ...
func (u AchievementsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var achievements []models.Achievement

	conn := db.GetInstance()
	if err := conn.Find(&achievements).Error; err != nil {
		log.Println("achievement not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "achievement not found.",
		})
	} else {
		log.Println("achievement returned")
		c.JSON(http.StatusOK, achievements)
	}

}

//Add ...
func (u AchievementsController) Add(c *gin.Context) {

	var achievement models.Achievement

	if err := c.ShouldBindJSON(&achievement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&achievement); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&achievement)

	c.JSON(http.StatusCreated, achievement)

}

//Modify ...
func (u AchievementsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var achievement models.Achievement

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&achievement).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("achievement not found.")
	}

	if err := c.ShouldBindJSON(&achievement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&achievement); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&achievement)

	c.JSON(http.StatusOK, achievement)
}

//Remove ...
func (u AchievementsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var achievement models.Achievement

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&achievement).Error; err != nil {
		log.Println("achievement not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "achievement not found.",
		})
	} else {
		log.Println("achievement returned")
		c.JSON(http.StatusOK, achievement)
	}

}
