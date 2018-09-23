package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//JudgingPanelMembersController ...
type JudgingPanelMembersController struct{}

//Get ...
func (u JudgingPanelMembersController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//List ...
func (u JudgingPanelMembersController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u JudgingPanelMembersController) Add(c *gin.Context) {
	var json models.JudgingPanelMember
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u JudgingPanelMembersController) Modify(c *gin.Context) {
	var json models.JudgingPanelMember
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u JudgingPanelMembersController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
