package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//PaymentsController ...
type PaymentsController struct{}

//Get ...
func (u PaymentsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var payment models.Payment

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&payment).Error; err != nil {
		log.Println("payment not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "payment not found.",
		})
	} else {
		log.Println("payment returned")
		c.JSON(http.StatusOK, payment)
	}

}

//List ...
func (u PaymentsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var payments []models.Payment

	conn := db.GetInstance()
	if err := conn.Find(&payments).Error; err != nil {
		log.Println("payment not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "payment not found.",
		})
	} else {
		log.Println("payment returned")
		c.JSON(http.StatusOK, payments)
	}

}

//Add ...
func (u PaymentsController) Add(c *gin.Context) {

	var payment models.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&payment)

	c.JSON(http.StatusCreated, payment)

}

//Modify ...
func (u PaymentsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var payment models.Payment

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&payment).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("payment not found.")
	}

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&payment)

	c.JSON(http.StatusOK, payment)
}

//Remove ...
func (u PaymentsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var payment models.Payment

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&payment).Error; err != nil {
		log.Println("payment not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "payment not found.",
		})
	} else {
		log.Println("payment returned")
		c.JSON(http.StatusOK, payment)
	}

}
