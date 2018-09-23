package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//WarningTemplatesController ...
type WarningTemplatesController struct{}

//Get ...
func (u WarningTemplatesController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u WarningTemplatesController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u WarningTemplatesController) Add(c *gin.Context) {
	var json models.WarningTemplate
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u WarningTemplatesController) Modify(c *gin.Context) {
	var json models.WarningTemplate
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u WarningTemplatesController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
