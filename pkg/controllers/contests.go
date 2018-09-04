package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//ContestsController ...
type ContestsController struct{}

//Get ...
func (u ContestsController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u ContestsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u ContestsController) Add(c *gin.Context) {
	var json models.Contest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u ContestsController) Modify(c *gin.Context) {
	var json models.Contest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u ContestsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
