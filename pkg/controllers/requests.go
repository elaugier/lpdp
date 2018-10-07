package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//RequestsController ...
type RequestsController struct{}

//Get ...
func (u RequestsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var request models.Request

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&request).Error; err != nil {
		log.Println("request not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "request not found.",
		})
	} else {
		log.Println("request returned")
		c.JSON(http.StatusOK, request)
	}

}

//List ...
func (u RequestsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var requests []models.Request

	conn := db.GetInstance()
	if err := conn.Find(&requests).Error; err != nil {
		log.Println("request not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "request not found.",
		})
	} else {
		log.Println("request returned")
		c.JSON(http.StatusOK, requests)
	}

}

//Add ...
func (u RequestsController) Add(c *gin.Context) {

	var request models.Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&request)

	c.JSON(http.StatusCreated, request)

}

//Modify ...
func (u RequestsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var request models.Request

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&request).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("request not found.")
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&request)

	c.JSON(http.StatusOK, request)
}

//Remove ...
func (u RequestsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var request models.Request

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&request).Error; err != nil {
		log.Println("request not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "request not found.",
		})
	} else {
		log.Println("request returned")
		c.JSON(http.StatusOK, request)
	}

}
