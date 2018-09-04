package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//LoginController ...
type LoginController struct{}

//Get ...
func (u LoginController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u LoginController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Login ...
func (u LoginController) Login(c *gin.Context) {
	var json models.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.User != "admin" || json.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

//Logout ...
func (u LoginController) Logout(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
