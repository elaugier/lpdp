package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//CorrectionRequestActionsController ...
type CorrectionRequestActionsController struct{}

//Get ...
func (u CorrectionRequestActionsController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u CorrectionRequestActionsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u CorrectionRequestActionsController) Add(c *gin.Context) {
	var json models.CorrectionRequestAction
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u CorrectionRequestActionsController) Modify(c *gin.Context) {
	var json models.CorrectionRequestAction
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u CorrectionRequestActionsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
