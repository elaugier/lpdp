package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//AlertsController ...
type AlertsController struct{}

//Get ...
func (u AlertsController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u AlertsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u AlertsController) Add(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u AlertsController) Modify(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u AlertsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
