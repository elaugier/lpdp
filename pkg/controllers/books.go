package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//BooksController ...
type BooksController struct{}

//Get ...
func (u BooksController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u BooksController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u BooksController) Add(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u BooksController) Modify(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u BooksController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
