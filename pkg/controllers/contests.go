package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//ContestsController ...
type ContestsController struct{}

//Get ...
func (u ContestsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var contest models.Contest

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&contest).Error; err != nil {
		log.Println("contest not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "contest not found.",
		})
	} else {
		log.Println("contest returned")
		c.JSON(http.StatusOK, contest)
	}

}

//List ...
func (u ContestsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var contests []models.Contest

	conn := db.GetInstance()
	if err := conn.Find(&contests).Error; err != nil {
		log.Println("contest not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "contest not found.",
		})
	} else {
		log.Println("contest returned")
		c.JSON(http.StatusOK, contests)
	}

}

//Add ...
func (u ContestsController) Add(c *gin.Context) {

	var contest models.Contest

	if err := c.ShouldBindJSON(&contest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&contest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&contest)

	c.JSON(http.StatusCreated, contest)

}

//Modify ...
func (u ContestsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var contest models.Contest

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&contest).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("contest not found.")
	}

	if err := c.ShouldBindJSON(&contest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&contest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&contest)

	c.JSON(http.StatusOK, contest)
}

//Remove ...
func (u ContestsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var contest models.Contest

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&contest).Error; err != nil {
		log.Println("contest not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "contest not found.",
		})
	} else {
		log.Println("contest returned")
		c.JSON(http.StatusOK, contest)
	}

}
