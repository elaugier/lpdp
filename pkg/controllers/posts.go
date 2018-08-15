package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//PostsController ...
type PostsController struct{}

//Get ...
func (u PostsController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u PostsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u PostsController) Add(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u PostsController) Modify(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u PostsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
