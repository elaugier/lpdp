package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//BookPartsController ...
type BookPartsController struct{}

//Get ...
func (u BookPartsController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u BookPartsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u BookPartsController) Add(c *gin.Context) {
	var json models.BookPart
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u BookPartsController) Modify(c *gin.Context) {
	var json models.BookPart
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u BookPartsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
