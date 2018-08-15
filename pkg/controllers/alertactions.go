package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//AlertActionsController ...
type AlertActionsController struct{}

//Get ...
func (u AlertActionsController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u AlertActionsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u AlertActionsController) Add(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u AlertActionsController) Modify(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u AlertActionsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
