package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//BadgesController ...
type BadgesController struct{}

//Get ...
func (u BadgesController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var badge models.Badge

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&badge).Error; err != nil {
		log.Println("badge not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "badge not found.",
		})
	} else {
		log.Println("badge returned")
		c.JSON(http.StatusOK, badge)
	}

}

//List ...
func (u BadgesController) List(c *gin.Context) {

	log := logs.GetInstance()

	var badges []models.Badge

	conn := db.GetInstance()
	if err := conn.Find(&badges).Error; err != nil {
		log.Println("badge not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "badge not found.",
		})
	} else {
		log.Println("badge returned")
		c.JSON(http.StatusOK, badges)
	}

}

//Add ...
func (u BadgesController) Add(c *gin.Context) {

	var badge models.Badge

	if err := c.ShouldBindJSON(&badge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&badge); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&badge)

	c.JSON(http.StatusCreated, badge)

}

//Modify ...
func (u BadgesController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var badge models.Badge

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&badge).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("badge not found.")
	}

	if err := c.ShouldBindJSON(&badge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&badge); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&badge)

	c.JSON(http.StatusOK, badge)
}

//Remove ...
func (u BadgesController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var badge models.Badge

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&badge).Error; err != nil {
		log.Println("badge not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "badge not found.",
		})
	} else {
		log.Println("badge returned")
		c.JSON(http.StatusOK, badge)
	}

}
