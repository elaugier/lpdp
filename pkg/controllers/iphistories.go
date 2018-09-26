package controllers

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

//IPHistoriesController ...
type IPHistoriesController struct{}

//Get ...
func (u IPHistoriesController) Get(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}
	var i models.IPHistory
	conn := db.GetInstance()
	if err = conn.Where("id = ?", id).First(&i).Error; err != nil {
		c.JSON(404, gin.H{
			"msg": "ip history not found.",
		})
	} else {
		c.JSON(200, i)
	}
}

//List ...
func (u IPHistoriesController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u IPHistoriesController) Add(c *gin.Context) {
	var json models.IPHistory
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u IPHistoriesController) Modify(c *gin.Context) {
	var json models.IPHistory
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u IPHistoriesController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
