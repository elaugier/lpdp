package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//WarningTemplatesController ...
type WarningTemplatesController struct{}

//Get ...
func (u WarningTemplatesController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var warningtemplate models.WarningTemplate

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&warningtemplate).Error; err != nil {
		log.Println("warningtemplate not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "warningtemplate not found.",
		})
	} else {
		log.Println("warningtemplate returned")
		c.JSON(http.StatusOK, warningtemplate)
	}

}

//List ...
func (u WarningTemplatesController) List(c *gin.Context) {

	log := logs.GetInstance()

	var warningtemplates []models.WarningTemplate

	conn := db.GetInstance()
	if err := conn.Find(&warningtemplates).Error; err != nil {
		log.Println("warningtemplate not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "warningtemplate not found.",
		})
	} else {
		log.Println("warningtemplate returned")
		c.JSON(http.StatusOK, warningtemplates)
	}

}

//Add ...
func (u WarningTemplatesController) Add(c *gin.Context) {

	var warningtemplate models.WarningTemplate

	if err := c.ShouldBindJSON(&warningtemplate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&warningtemplate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&warningtemplate)

	c.JSON(http.StatusCreated, warningtemplate)

}

//Modify ...
func (u WarningTemplatesController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var warningtemplate models.WarningTemplate

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&warningtemplate).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("warningtemplate not found.")
	}

	if err := c.ShouldBindJSON(&warningtemplate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&warningtemplate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&warningtemplate)

	c.JSON(http.StatusOK, warningtemplate)
}

//Remove ...
func (u WarningTemplatesController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var warningtemplate models.WarningTemplate

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&warningtemplate).Error; err != nil {
		log.Println("warningtemplate not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "warningtemplate not found.",
		})
	} else {
		log.Println("warningtemplate returned")
		c.JSON(http.StatusOK, warningtemplate)
	}

}
