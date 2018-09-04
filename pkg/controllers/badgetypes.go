package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//BadgeTypesController ...
type BadgeTypesController struct{}

//Get ...
func (u BadgeTypesController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u BadgeTypesController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u BadgeTypesController) Add(c *gin.Context) {
	var json models.BadgeType
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u BadgeTypesController) Modify(c *gin.Context) {
	var json models.BadgeType
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u BadgeTypesController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
