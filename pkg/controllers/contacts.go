package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ContactsController ...
type ContactsController struct{}

//Get ...
func (u ContactsController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u ContactsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u ContactsController) Add(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u ContactsController) Modify(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u ContactsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
