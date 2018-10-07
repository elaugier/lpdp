package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//LicensesController ...
type LicensesController struct{}

//Get ...
func (u LicensesController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var license models.License

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&license).Error; err != nil {
		log.Println("license not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "license not found.",
		})
	} else {
		log.Println("license returned")
		c.JSON(http.StatusOK, license)
	}

}

//List ...
func (u LicensesController) List(c *gin.Context) {

	log := logs.GetInstance()

	var licenses []models.License

	conn := db.GetInstance()
	if err := conn.Find(&licenses).Error; err != nil {
		log.Println("license not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "license not found.",
		})
	} else {
		log.Println("license returned")
		c.JSON(http.StatusOK, licenses)
	}

}

//Add ...
func (u LicensesController) Add(c *gin.Context) {

	var license models.License

	if err := c.ShouldBindJSON(&license); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&license); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&license)

	c.JSON(http.StatusCreated, license)

}

//Modify ...
func (u LicensesController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var license models.License

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&license).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("license not found.")
	}

	if err := c.ShouldBindJSON(&license); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&license); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&license)

	c.JSON(http.StatusOK, license)
}

//Remove ...
func (u LicensesController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var license models.License

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&license).Error; err != nil {
		log.Println("license not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "license not found.",
		})
	} else {
		log.Println("license returned")
		c.JSON(http.StatusOK, license)
	}

}
