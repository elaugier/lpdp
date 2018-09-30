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
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(400, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}
	var v models.Voting
	conn := db.GetInstance()
	if err = conn.Where("id = ?", id).First(&v).Error; err != nil {
		log.Println("voting not found.")
		c.JSON(404, gin.H{
			"msg": "voting not found.",
		})
	} else {
		log.Println("voting returned")
		c.JSON(200, v)
	}
}

//List ...
func (u VotingsController) List(c *gin.Context) {
	log := logs.GetInstance()
	var v []models.Voting
	conn := db.GetInstance()
	if err := conn.Find(&v).Error; err != nil {
		log.Println("voting not found.")
		c.JSON(404, gin.H{
			"msg": "voting not found.",
		})
	} else {
		log.Println("voting returned")
		c.JSON(200, v)
	}
}

//Add ...
func (u VotingsController) Add(c *gin.Context) {
	var json models.Voting
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u VotingsController) Modify(c *gin.Context) {
	var json models.Voting
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u VotingsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
