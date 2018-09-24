package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/logs"

	"github.com/google/uuid"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//UsersController ...
type UsersController struct {
}

//Get ...
func (u UsersController) Get(c *gin.Context) {
	id := c.Param("id")
	uid, _ := uuid.Parse(id)
	logs.Inst.L.Println("get uuid %u", uid)
	//db.Users.GetUser(uid)
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u UsersController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u UsersController) Add(c *gin.Context) {
	var json models.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u UsersController) Modify(c *gin.Context) {
	var json models.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u UsersController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
