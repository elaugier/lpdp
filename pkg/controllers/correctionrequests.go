package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//CorrectionRequestsController ...
type CorrectionRequestsController struct{}

//Get ...
func (u CorrectionRequestsController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u CorrectionRequestsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u CorrectionRequestsController) Add(c *gin.Context) {
	var json models.CorrectionRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u CorrectionRequestsController) Modify(c *gin.Context) {
	var json models.CorrectionRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u CorrectionRequestsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
