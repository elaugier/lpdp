package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//IPHistoriesController ...
type IPHistoriesController struct{}

//Get ...
func (u IPHistoriesController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var iphistory models.IPHistory

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&iphistory).Error; err != nil {
		log.Println("iphistory not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "iphistory not found.",
		})
	} else {
		log.Println("iphistory returned")
		c.JSON(http.StatusOK, iphistory)
	}

}

//List ...
func (u IPHistoriesController) List(c *gin.Context) {

	log := logs.GetInstance()

	var iphistorys []models.IPHistory

	conn := db.GetInstance()
	if err := conn.Find(&iphistorys).Error; err != nil {
		log.Println("iphistory not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "iphistory not found.",
		})
	} else {
		log.Println("iphistory returned")
		c.JSON(http.StatusOK, iphistorys)
	}

}

//Add ...
func (u IPHistoriesController) Add(c *gin.Context) {

	var iphistory models.IPHistory

	if err := c.ShouldBindJSON(&iphistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&iphistory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&iphistory)

	c.JSON(http.StatusCreated, iphistory)

}

//Modify ...
func (u IPHistoriesController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var iphistory models.IPHistory

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&iphistory).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("iphistory not found.")
	}

	if err := c.ShouldBindJSON(&iphistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&iphistory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&iphistory)

	c.JSON(http.StatusOK, iphistory)
}

//Remove ...
func (u IPHistoriesController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var iphistory models.IPHistory

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&iphistory).Error; err != nil {
		log.Println("iphistory not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "iphistory not found.",
		})
	} else {
		log.Println("iphistory returned")
		c.JSON(http.StatusOK, iphistory)
	}

}
