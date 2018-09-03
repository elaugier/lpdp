package controllers

import (
	"net/http"

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
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u ContestsController) Modify(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u ContestsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
