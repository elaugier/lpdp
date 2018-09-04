package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//UsersController ...
type UsersController struct{}

//Get ...
func (u UsersController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u UsersController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u UsersController) Add(c *gin.Context) {
	var json models.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u UsersController) Modify(c *gin.Context) {
	var json models.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u UsersController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
