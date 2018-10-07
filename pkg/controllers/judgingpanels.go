package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//JudgingPanelsController ...
type JudgingPanelsController struct{}

//Get ...
func (u JudgingPanelsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var judgingpanel models.JudgingPanel

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&judgingpanel).Error; err != nil {
		log.Println("judgingpanel not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "judgingpanel not found.",
		})
	} else {
		log.Println("judgingpanel returned")
		c.JSON(http.StatusOK, judgingpanel)
	}

}

//List ...
func (u JudgingPanelsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var judgingpanels []models.JudgingPanel

	conn := db.GetInstance()
	if err := conn.Find(&judgingpanels).Error; err != nil {
		log.Println("judgingpanel not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "judgingpanel not found.",
		})
	} else {
		log.Println("judgingpanel returned")
		c.JSON(http.StatusOK, judgingpanels)
	}

}

//Add ...
func (u JudgingPanelsController) Add(c *gin.Context) {

	var judgingpanel models.JudgingPanel

	if err := c.ShouldBindJSON(&judgingpanel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&judgingpanel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&judgingpanel)

	c.JSON(http.StatusCreated, judgingpanel)

}

//Modify ...
func (u JudgingPanelsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var judgingpanel models.JudgingPanel

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&judgingpanel).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("judgingpanel not found.")
	}

	if err := c.ShouldBindJSON(&judgingpanel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&judgingpanel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&judgingpanel)

	c.JSON(http.StatusOK, judgingpanel)
}

//Remove ...
func (u JudgingPanelsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var judgingpanel models.JudgingPanel

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&judgingpanel).Error; err != nil {
		log.Println("judgingpanel not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "judgingpanel not found.",
		})
	} else {
		log.Println("judgingpanel returned")
		c.JSON(http.StatusOK, judgingpanel)
	}

}
