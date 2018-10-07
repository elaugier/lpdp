package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//WarningsController ...
type WarningsController struct{}

//Get ...
func (u WarningsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var warning models.Warning

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&warning).Error; err != nil {
		log.Println("warning not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "warning not found.",
		})
	} else {
		log.Println("warning returned")
		c.JSON(http.StatusOK, warning)
	}

}

//List ...
func (u WarningsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var warnings []models.Warning

	conn := db.GetInstance()
	if err := conn.Find(&warnings).Error; err != nil {
		log.Println("warning not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "warning not found.",
		})
	} else {
		log.Println("warning returned")
		c.JSON(http.StatusOK, warnings)
	}

}

//Add ...
func (u WarningsController) Add(c *gin.Context) {

	var warning models.Warning

	if err := c.ShouldBindJSON(&warning); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&warning); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&warning)

	c.JSON(http.StatusCreated, warning)

}

//Modify ...
func (u WarningsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var warning models.Warning

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&warning).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("warning not found.")
	}

	if err := c.ShouldBindJSON(&warning); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&warning); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&warning)

	c.JSON(http.StatusOK, warning)
}

//Remove ...
func (u WarningsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var warning models.Warning

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&warning).Error; err != nil {
		log.Println("warning not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "warning not found.",
		})
	} else {
		log.Println("warning returned")
		c.JSON(http.StatusOK, warning)
	}

}
