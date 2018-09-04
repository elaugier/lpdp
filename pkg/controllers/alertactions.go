package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//AlertActionsController ...
type AlertActionsController struct{}

//Get ...
func (u AlertActionsController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u AlertActionsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u AlertActionsController) Add(c *gin.Context) {
	var json models.AlertAction
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u AlertActionsController) Modify(c *gin.Context) {
	var json models.AlertAction
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u AlertActionsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
