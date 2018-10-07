package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//VotesController ...
type VotesController struct{}

//Get ...
func (u VotesController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var vote models.Vote

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&vote).Error; err != nil {
		log.Println("vote not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "vote not found.",
		})
	} else {
		log.Println("vote returned")
		c.JSON(http.StatusOK, vote)
	}

}

//List ...
func (u VotesController) List(c *gin.Context) {

	log := logs.GetInstance()

	var votes []models.Vote

	conn := db.GetInstance()
	if err := conn.Find(&votes).Error; err != nil {
		log.Println("vote not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "vote not found.",
		})
	} else {
		log.Println("vote returned")
		c.JSON(http.StatusOK, votes)
	}

}

//Add ...
func (u VotesController) Add(c *gin.Context) {

	var vote models.Vote

	if err := c.ShouldBindJSON(&vote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&vote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&vote)

	c.JSON(http.StatusCreated, vote)

}

//Modify ...
func (u VotesController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var vote models.Vote

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&vote).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("vote not found.")
	}

	if err := c.ShouldBindJSON(&vote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&vote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&vote)

	c.JSON(http.StatusOK, vote)
}

//Remove ...
func (u VotesController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var vote models.Vote

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&vote).Error; err != nil {
		log.Println("vote not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "vote not found.",
		})
	} else {
		log.Println("vote returned")
		c.JSON(http.StatusOK, vote)
	}

}
