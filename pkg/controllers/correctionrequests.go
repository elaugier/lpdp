package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//CorrectionRequestsController ...
type CorrectionRequestsController struct{}

//Get ...
func (u CorrectionRequestsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var correctionrequest models.CorrectionRequest

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&correctionrequest).Error; err != nil {
		log.Println("correctionrequest not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "correctionrequest not found.",
		})
	} else {
		log.Println("correctionrequest returned")
		c.JSON(http.StatusOK, correctionrequest)
	}

}

//List ...
func (u CorrectionRequestsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var correctionrequests []models.CorrectionRequest

	conn := db.GetInstance()
	if err := conn.Find(&correctionrequests).Error; err != nil {
		log.Println("correctionrequest not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "correctionrequest not found.",
		})
	} else {
		log.Println("correctionrequest returned")
		c.JSON(http.StatusOK, correctionrequests)
	}

}

//Add ...
func (u CorrectionRequestsController) Add(c *gin.Context) {

	var correctionrequest models.CorrectionRequest

	if err := c.ShouldBindJSON(&correctionrequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&correctionrequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&correctionrequest)

	c.JSON(http.StatusCreated, correctionrequest)

}

//Modify ...
func (u CorrectionRequestsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var correctionrequest models.CorrectionRequest

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&correctionrequest).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("correctionrequest not found.")
	}

	if err := c.ShouldBindJSON(&correctionrequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&correctionrequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&correctionrequest)

	c.JSON(http.StatusOK, correctionrequest)
}

//Remove ...
func (u CorrectionRequestsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var correctionrequest models.CorrectionRequest

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&correctionrequest).Error; err != nil {
		log.Println("correctionrequest not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "correctionrequest not found.",
		})
	} else {
		log.Println("correctionrequest returned")
		c.JSON(http.StatusOK, correctionrequest)
	}

}
