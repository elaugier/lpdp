package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//LikesController ...
type LikesController struct{}

//Get ...
func (u LikesController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var like models.Like

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&like).Error; err != nil {
		log.Println("like not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "like not found.",
		})
	} else {
		log.Println("like returned")
		c.JSON(http.StatusOK, like)
	}

}

//List ...
func (u LikesController) List(c *gin.Context) {

	log := logs.GetInstance()

	var likes []models.Like

	conn := db.GetInstance()
	if err := conn.Find(&likes).Error; err != nil {
		log.Println("like not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "like not found.",
		})
	} else {
		log.Println("like returned")
		c.JSON(http.StatusOK, likes)
	}

}

//Add ...
func (u LikesController) Add(c *gin.Context) {

	var like models.Like

	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&like); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&like)

	c.JSON(http.StatusCreated, like)

}

//Modify ...
func (u LikesController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var like models.Like

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&like).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("like not found.")
	}

	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&like); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&like)

	c.JSON(http.StatusOK, like)
}

//Remove ...
func (u LikesController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var like models.Like

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&like).Error; err != nil {
		log.Println("like not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "like not found.",
		})
	} else {
		log.Println("like returned")
		c.JSON(http.StatusOK, like)
	}

}
