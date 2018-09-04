package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//AlertsController ...
type AlertsController struct{}

//Get ...
func (u AlertsController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u AlertsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
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
