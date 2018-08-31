package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//TagsController ...
type TagsController struct{}

//Get ...
func (u TagsController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u TagsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u TagsController) Add(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u TagsController) Modify(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u TagsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
