package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//BadgeTypesController ...
type BadgeTypesController struct{}

//Get ...
func (u BadgeTypesController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var badgetype models.BadgeType

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&badgetype).Error; err != nil {
		log.Println("badge type not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "badge type not found.",
		})
	} else {
		log.Println("badge type returned")
		c.JSON(http.StatusOK, badgetype)
	}

}

//List ...
func (u BadgeTypesController) List(c *gin.Context) {

	log := logs.GetInstance()

	var badgetypes []models.BadgeType

	conn := db.GetInstance()
	if err := conn.Find(&badgetypes).Error; err != nil {
		log.Println("badge type not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "badge type not found.",
		})
	} else {
		log.Println("badge type returned")
		c.JSON(http.StatusOK, badgetypes)
	}

}

//Add ...
func (u BadgeTypesController) Add(c *gin.Context) {

	var badgetype models.BadgeType

	if err := c.ShouldBindJSON(&badgetype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&badgetype); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&badgetype)

	c.JSON(http.StatusCreated, badgetype)

}

//Modify ...
func (u BadgeTypesController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var badgetype models.BadgeType

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&badgetype).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("badgetype not found.")
	}

	if err := c.ShouldBindJSON(&badgetype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&badgetype); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&badgetype)

	c.JSON(http.StatusOK, badgetype)
}

//Remove ...
func (u BadgeTypesController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var badgetype models.BadgeType

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&badgetype).Error; err != nil {
		log.Println("badge type not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "badge type not found.",
		})
	} else {
		log.Println("badge type returned")
		c.JSON(http.StatusOK, badgetype)
	}

}
