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
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(400, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}
	var i models.Vote
	conn := db.GetInstance()
	if err = conn.Where("id = ?", id).First(&i).Error; err != nil {
		log.Println("vote not found.")
		c.JSON(404, gin.H{
			"msg": "vote not found.",
		})
	} else {
		log.Println("vote returned")
		c.JSON(200, i)
	}
}

//List ...
func (u VotesController) List(c *gin.Context) {
	log := logs.GetInstance()
	var votes []models.Vote
	conn := db.GetInstance()
	if err := conn.Find(&votes).Error; err != nil {
		log.Println("vote not found.")
		c.JSON(404, gin.H{
			"msg": "vote not found.",
		})
	} else {
		log.Println("vote returned")
		c.JSON(200, votes)
	}
}

//Add ...
func (u VotesController) Add(c *gin.Context) {
	var json models.Vote
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u VotesController) Modify(c *gin.Context) {
	var json models.Vote
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u VotesController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
