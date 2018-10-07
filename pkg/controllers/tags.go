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

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var tag models.Tag

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&tag).Error; err != nil {
		log.Println("tag not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "tag not found.",
		})
	} else {
		log.Println("tag returned")
		c.JSON(http.StatusOK, tag)
	}

}

//List ...
func (u TagsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var tags []models.Tag

	conn := db.GetInstance()
	if err := conn.Find(&tags).Error; err != nil {
		log.Println("tag not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "tag not found.",
		})
	} else {
		log.Println("tag returned")
		c.JSON(http.StatusOK, tags)
	}

}

//Add ...
func (u TagsController) Add(c *gin.Context) {

	var tag models.Tag

	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&tag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&tag)

	c.JSON(http.StatusCreated, tag)

}

//Modify ...
func (u TagsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var tag models.Tag

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&tag).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("tag not found.")
	}

	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&tag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&tag)

	c.JSON(http.StatusOK, tag)
}

//Remove ...
func (u TagsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var tag models.Tag

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&tag).Error; err != nil {
		log.Println("tag not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "tag not found.",
		})
	} else {
		log.Println("tag returned")
		c.JSON(http.StatusOK, tag)
	}

}
