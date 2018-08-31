package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//BadIPAddressesController ...
type BadIPAddressesController struct{}

//Get ...
func (u BadIPAddressesController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u BadIPAddressesController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u BadIPAddressesController) Add(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u BadIPAddressesController) Modify(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u BadIPAddressesController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
