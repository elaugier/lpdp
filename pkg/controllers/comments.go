package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//CommentsController ...
type CommentsController struct{}

//Get ...
func (u CommentsController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u CommentsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u CommentsController) Add(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u CommentsController) Modify(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u CommentsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
