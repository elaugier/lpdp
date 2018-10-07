package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//CorrectionRequestActionsController ...
type CorrectionRequestActionsController struct{}

//Get ...
func (u CorrectionRequestActionsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var correctionrequestaction models.CorrectionRequestAction

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&correctionrequestaction).Error; err != nil {
		log.Println("correctionrequestaction not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "correctionrequestaction not found.",
		})
	} else {
		log.Println("correctionrequestaction returned")
		c.JSON(http.StatusOK, correctionrequestaction)
	}

}

//List ...
func (u CorrectionRequestActionsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var correctionrequestactions []models.CorrectionRequestAction

	conn := db.GetInstance()
	if err := conn.Find(&correctionrequestactions).Error; err != nil {
		log.Println("correctionrequestaction not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "correctionrequestaction not found.",
		})
	} else {
		log.Println("correctionrequestaction returned")
		c.JSON(http.StatusOK, correctionrequestactions)
	}

}

//Add ...
func (u CorrectionRequestActionsController) Add(c *gin.Context) {

	var correctionrequestaction models.CorrectionRequestAction

	if err := c.ShouldBindJSON(&correctionrequestaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&correctionrequestaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&correctionrequestaction)

	c.JSON(http.StatusCreated, correctionrequestaction)

}

//Modify ...
func (u CorrectionRequestActionsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var correctionrequestaction models.CorrectionRequestAction

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&correctionrequestaction).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("correctionrequestaction not found.")
	}

	if err := c.ShouldBindJSON(&correctionrequestaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&correctionrequestaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&correctionrequestaction)

	c.JSON(http.StatusOK, correctionrequestaction)
}

//Remove ...
func (u CorrectionRequestActionsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var correctionrequestaction models.CorrectionRequestAction

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&correctionrequestaction).Error; err != nil {
		log.Println("correctionrequestaction not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "correctionrequestaction not found.",
		})
	} else {
		log.Println("correctionrequestaction returned")
		c.JSON(http.StatusOK, correctionrequestaction)
	}

}
