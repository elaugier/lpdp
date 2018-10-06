package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//ContestRoundsController ...
type ContestRoundsController struct{}

//Get ...
func (u ContestRoundsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var contestround models.ContestRound

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&contestround).Error; err != nil {
		log.Println("contestround not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "contestround not found.",
		})
	} else {
		log.Println("contestround returned")
		c.JSON(http.StatusOK, contestround)
	}

}

//List ...
func (u ContestRoundsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var contestrounds []models.ContestRound

	conn := db.GetInstance()
	if err := conn.Find(&contestrounds).Error; err != nil {
		log.Println("contestround not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "contestround not found.",
		})
	} else {
		log.Println("contestround returned")
		c.JSON(http.StatusOK, contestrounds)
	}

}

//Add ...
func (u ContestRoundsController) Add(c *gin.Context) {

	var contestround models.ContestRound

	if err := c.ShouldBindJSON(&contestround); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&contestround); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&contestround)

	c.JSON(http.StatusCreated, contestround)

}

//Modify ...
func (u ContestRoundsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var contestround models.ContestRound

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&contestround).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("contestround not found.")
	}

	if err := c.ShouldBindJSON(&contestround); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&contestround); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&contestround)

	c.JSON(http.StatusOK, contestround)
}

//Remove ...
func (u ContestRoundsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var contestround models.ContestRound

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&contestround).Error; err != nil {
		log.Println("contestround not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "contestround not found.",
		})
	} else {
		log.Println("contestround returned")
		c.JSON(http.StatusOK, contestround)
	}

}
