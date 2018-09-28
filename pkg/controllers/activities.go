package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/logs"

	"github.com/google/uuid"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//ActivitiesController ...
type ActivitiesController struct{}

//Get ...
func (u ActivitiesController) Get(c *gin.Context) {
	log := logs.GetInstance()
	log.Println("try to retieve 'id' in url path")
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(400, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}
	var i models.Activity
	conn := db.GetInstance()
	if err = conn.Where("id = ?", id).First(&i).Error; err != nil {
		log.Println("activity not found.")
		c.JSON(404, gin.H{
			"msg": "activity not found.",
		})
	} else {
		log.Println("activity returned")
		c.JSON(200, i)
	}
}

//List ...
func (u ActivitiesController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u ActivitiesController) Add(c *gin.Context) {
	var json models.Activity
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u ActivitiesController) Modify(c *gin.Context) {
	var json models.Activity
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u ActivitiesController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
