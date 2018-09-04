package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//AchievementsController ...
type AchievementsController struct{}

//Get ...
func (u AchievementsController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u AchievementsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
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
