package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//ActivitiesController ...
type ActivitiesController struct{}

//Get ...
func (u ActivitiesController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var activity models.Activity

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&activity).Error; err != nil {
		log.Println("activity not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "activity not found.",
		})
	} else {
		log.Println("activity returned")
		c.JSON(http.StatusOK, activity)
	}

}

//List ...
func (u ActivitiesController) List(c *gin.Context) {

	log := logs.GetInstance()

	var activities []models.Activity

	conn := db.GetInstance()
	if err := conn.Find(&activities).Error; err != nil {
		log.Println("activity not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "activity not found.",
		})
	} else {
		log.Println("activity returned")
		c.JSON(http.StatusOK, activities)
	}

}

//Add ...
func (u ActivitiesController) Add(c *gin.Context) {

	var activity models.Activity

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&activity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&activity)

	c.JSON(http.StatusCreated, activity)

}

//Modify ...
func (u ActivitiesController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var activity models.Activity

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path",
		})
	}

	conn := db.GetInstance()

	if err := conn.Where("id = ?", id).First(&activity).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("activity not found")
	}

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&activity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&activity)

	c.JSON(http.StatusOK, activity)
}

//Remove ...
func (u ActivitiesController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var activity models.Activity

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&activity).Error; err != nil {
		log.Println("activity not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "activity not found.",
		})
	} else {
		log.Println("activity returned")
		c.JSON(http.StatusOK, activity)
	}

}
