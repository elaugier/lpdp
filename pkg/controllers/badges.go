package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//BadgesController ...
type BadgesController struct{}

//Get ...
func (u BadgesController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u BadgesController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u BadgesController) Add(c *gin.Context) {
	var json models.Badge
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u BadgesController) Modify(c *gin.Context) {
	var json models.Badge
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u BadgesController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
