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
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(400, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}
	var i models.ContestRound
	conn := db.GetInstance()
	if err = conn.Where("id = ?", id).First(&i).Error; err != nil {
		log.Println("contest round not found.")
		c.JSON(404, gin.H{
			"msg": "contest round not found.",
		})
	} else {
		log.Println("contest round returned")
		c.JSON(200, i)
	}
}

//List ...
func (u ContestRoundsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u ContestRoundsController) Add(c *gin.Context) {
	var json models.ContestRound
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u ContestRoundsController) Modify(c *gin.Context) {
	var json models.ContestRound
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u ContestRoundsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
