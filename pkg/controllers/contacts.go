package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//ContactsController ...
type ContactsController struct{}

//Get ...
func (u ContactsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var contact models.Contact

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&contact).Error; err != nil {
		log.Println("contact not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "contact not found.",
		})
	} else {
		log.Println("contact returned")
		c.JSON(http.StatusOK, contact)
	}

}

//List ...
func (u ContactsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var contacts []models.Contact

	conn := db.GetInstance()
	if err := conn.Find(&contacts).Error; err != nil {
		log.Println("contact not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "contact not found.",
		})
	} else {
		log.Println("contact returned")
		c.JSON(http.StatusOK, contacts)
	}

}

//Add ...
func (u ContactsController) Add(c *gin.Context) {

	var contact models.Contact

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&contact)

	c.JSON(http.StatusCreated, contact)

}

//Modify ...
func (u ContactsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var contact models.Contact

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&contact).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("contact not found.")
	}

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&contact)

	c.JSON(http.StatusOK, contact)
}

//Remove ...
func (u ContactsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var contact models.Contact

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&contact).Error; err != nil {
		log.Println("contact not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "contact not found.",
		})
	} else {
		log.Println("contact returned")
		c.JSON(http.StatusOK, contact)
	}

}
