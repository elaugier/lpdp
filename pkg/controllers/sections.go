package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//SectionsController ...
type SectionsController struct{}

//Get ...
func (u SectionsController) Get(c *gin.Context) {
	log := logs.GetInstance()
	log.Println("try to retieve 'id' in url path")
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(400, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}
	var section models.Section
	conn := db.GetInstance()
	if err = conn.Where("id = ?", id).First(&section).Error; err != nil {
		log.Println("section not found.")
		c.JSON(404, gin.H{
			"msg": "section not found.",
		})
	} else {
		log.Println("section returned")
		c.JSON(200, section)
	}
}

//List ...
func (u SectionsController) List(c *gin.Context) {
	log := logs.GetInstance()
	var sections []models.Section
	conn := db.GetInstance()
	if err := conn.Find(&sections).Error; err != nil {
		log.Println("section not found.")
		c.JSON(404, gin.H{
			"msg": "section not found.",
		})
	} else {
		log.Println("section returned")
		c.JSON(200, sections)
	}
}

//Add ...
func (u SectionsController) Add(c *gin.Context) {
	var json models.Section
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u SectionsController) Modify(c *gin.Context) {
	var json models.Section
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u SectionsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
