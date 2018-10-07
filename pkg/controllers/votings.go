package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//VotingsController ...
type VotingsController struct{}

//Get ...
func (u VotingsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var voting models.Voting

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&voting).Error; err != nil {
		log.Println("voting not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "voting not found.",
		})
	} else {
		log.Println("voting returned")
		c.JSON(http.StatusOK, voting)
	}

}

//List ...
func (u VotingsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var votings []models.Voting

	conn := db.GetInstance()
	if err := conn.Find(&votings).Error; err != nil {
		log.Println("voting not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "voting not found.",
		})
	} else {
		log.Println("voting returned")
		c.JSON(http.StatusOK, votings)
	}

}

//Add ...
func (u VotingsController) Add(c *gin.Context) {

	var voting models.Voting

	if err := c.ShouldBindJSON(&voting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&voting); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&voting)

	c.JSON(http.StatusCreated, voting)

}

//Modify ...
func (u VotingsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var voting models.Voting

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&voting).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("voting not found.")
	}

	if err := c.ShouldBindJSON(&voting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&voting); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&voting)

	c.JSON(http.StatusOK, voting)
}

//Remove ...
func (u VotingsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var voting models.Voting

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&voting).Error; err != nil {
		log.Println("voting not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "voting not found.",
		})
	} else {
		log.Println("voting returned")
		c.JSON(http.StatusOK, voting)
	}

}
