package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//TagsController ...
type TagsController struct{}

//Get ...
func (u TagsController) Get(c *gin.Context) {
	log := logs.GetInstance()
	log.Println("try to retieve 'id' in url path")
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(400, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}
	var tag models.Tag
	conn := db.GetInstance()
	if err = conn.Where("id = ?", id).First(&tag).Error; err != nil {
		log.Println("tag not found.")
		c.JSON(404, gin.H{
			"msg": "tag not found.",
		})
	} else {
		log.Println("tag returned")
		c.JSON(200, tag)
	}
}

//List ...
func (u TagsController) List(c *gin.Context) {
	log := logs.GetInstance()
	var tags []models.Tag
	conn := db.GetInstance()
	if err := conn.Find(&tags).Error; err != nil {
		log.Println("tag not found.")
		c.JSON(404, gin.H{
			"msg": "tag not found.",
		})
	} else {
		log.Println("tag returned")
		c.JSON(200, tags)
	}
}

//Add ...
func (u TagsController) Add(c *gin.Context) {
	var json models.Tag
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u TagsController) Modify(c *gin.Context) {
	var json models.Tag
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u TagsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
