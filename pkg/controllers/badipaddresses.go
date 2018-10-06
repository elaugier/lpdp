package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//BadIPAddressesController ...
type BadIPAddressesController struct{}

//Get ...
func (u BadIPAddressesController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var badipaddress models.BadIPAddress

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&badipaddress).Error; err != nil {
		log.Println("ip address not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "ip address not found.",
		})
	} else {
		log.Println("ip address returned")
		c.JSON(http.StatusOK, badipaddress)
	}

}

//List ...
func (u BadIPAddressesController) List(c *gin.Context) {

	log := logs.GetInstance()

	var badipaddresses []models.BadIPAddress

	conn := db.GetInstance()
	if err := conn.Find(&badipaddresses).Error; err != nil {
		log.Println("ip address not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "ip address not found.",
		})
	} else {
		log.Println("ip address returned")
		c.JSON(http.StatusOK, badipaddresses)
	}

}

//Add ...
func (u BadIPAddressesController) Add(c *gin.Context) {

	var badipaddress models.BadIPAddress

	if err := c.ShouldBindJSON(&badipaddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&badipaddress); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&badipaddress)

	c.JSON(http.StatusCreated, badipaddress)

}

//Modify ...
func (u BadIPAddressesController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var badipaddress models.BadIPAddress

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&badipaddress).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("ip address not found.")
	}

	if err := c.ShouldBindJSON(&badipaddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&badipaddress); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&badipaddress)

	c.JSON(http.StatusOK, badipaddress)
}

//Remove ...
func (u BadIPAddressesController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var badipaddress models.BadIPAddress

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&badipaddress).Error; err != nil {
		log.Println("ip address not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "ip address not found.",
		})
	} else {
		log.Println("ip address returned")
		c.JSON(http.StatusOK, badipaddress)
	}

}
