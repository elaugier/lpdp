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

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var section models.Section

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&section).Error; err != nil {
		log.Println("section not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "section not found.",
		})
	} else {
		log.Println("section returned")
		c.JSON(http.StatusOK, section)
	}

}

//List ...
func (u SectionsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var sections []models.Section

	conn := db.GetInstance()
	if err := conn.Find(&sections).Error; err != nil {
		log.Println("section not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "section not found.",
		})
	} else {
		log.Println("section returned")
		c.JSON(http.StatusOK, sections)
	}

}

//Add ...
func (u SectionsController) Add(c *gin.Context) {

	var section models.Section

	if err := c.ShouldBindJSON(&section); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&section); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&section)

	c.JSON(http.StatusCreated, section)

}

//Modify ...
func (u SectionsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var section models.Section

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&section).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("section not found.")
	}

	if err := c.ShouldBindJSON(&section); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&section); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&section)

	c.JSON(http.StatusOK, section)
}

//Remove ...
func (u SectionsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var section models.Section

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&section).Error; err != nil {
		log.Println("section not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "section not found.",
		})
	} else {
		log.Println("section returned")
		c.JSON(http.StatusOK, section)
	}

}
