package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//ExitReasonsController ...
type ExitReasonsController struct{}

//Get ...
func (u ExitReasonsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var exitreason models.CorrectionRequest

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&exitreason).Error; err != nil {
		log.Println("exitreason not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "exitreason not found.",
		})
	} else {
		log.Println("exitreason returned")
		c.JSON(http.StatusOK, exitreason)
	}

}

//List ...
func (u ExitReasonsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var exitreasons []models.CorrectionRequest

	conn := db.GetInstance()
	if err := conn.Find(&exitreasons).Error; err != nil {
		log.Println("exitreason not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "exitreason not found.",
		})
	} else {
		log.Println("exitreason returned")
		c.JSON(http.StatusOK, exitreasons)
	}

}

//Add ...
func (u ExitReasonsController) Add(c *gin.Context) {

	var exitreason models.CorrectionRequest

	if err := c.ShouldBindJSON(&exitreason); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&exitreason); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&exitreason)

	c.JSON(http.StatusCreated, exitreason)

}

//Modify ...
func (u ExitReasonsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var exitreason models.CorrectionRequest

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&exitreason).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("exitreason not found.")
	}

	if err := c.ShouldBindJSON(&exitreason); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&exitreason); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&exitreason)

	c.JSON(http.StatusOK, exitreason)
}

//Remove ...
func (u ExitReasonsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var exitreason models.CorrectionRequest

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&exitreason).Error; err != nil {
		log.Println("exitreason not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "exitreason not found.",
		})
	} else {
		log.Println("exitreason returned")
		c.JSON(http.StatusOK, exitreason)
	}

}
