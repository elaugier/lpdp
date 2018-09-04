package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//CoauthorsController ...
type CoauthorsController struct{}

//Get ...
func (ca CoauthorsController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (ca CoauthorsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (ca CoauthorsController) Add(c *gin.Context) {
	var json models.CoAuthor
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (ca CoauthorsController) Modify(c *gin.Context) {
	var json models.CoAuthor
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (ca CoauthorsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
