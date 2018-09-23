package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//JudgingPanelsController ...
type JudgingPanelsController struct{}

//Get ...
func (u JudgingPanelsController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
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
