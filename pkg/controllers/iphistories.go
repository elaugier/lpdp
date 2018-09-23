package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//IPHistoriesController ...
type IPHistoriesController struct{}

//Get ...
func (u IPHistoriesController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u IPHistoriesController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u IPHistoriesController) Add(c *gin.Context) {
	var json models.IPHistory
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u IPHistoriesController) Modify(c *gin.Context) {
	var json models.IPHistory
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u IPHistoriesController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
