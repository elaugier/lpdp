package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ActivitiesController ...
type ActivitiesController struct{}

//Get ...
func (u ActivitiesController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u ActivitiesController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u ActivitiesController) Add(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u ActivitiesController) Modify(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u ActivitiesController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
