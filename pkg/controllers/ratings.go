package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//RatingsController ...
type RatingsController struct{}

//Get ...
func (u RatingsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var rating models.Rating

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&rating).Error; err != nil {
		log.Println("rating not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "rating not found.",
		})
	} else {
		log.Println("rating returned")
		c.JSON(http.StatusOK, rating)
	}

}

//List ...
func (u RatingsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var ratings []models.Rating

	conn := db.GetInstance()
	if err := conn.Find(&ratings).Error; err != nil {
		log.Println("rating not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "rating not found.",
		})
	} else {
		log.Println("rating returned")
		c.JSON(http.StatusOK, ratings)
	}

}

//Add ...
func (u RatingsController) Add(c *gin.Context) {

	var rating models.Rating

	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&rating); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&rating)

	c.JSON(http.StatusCreated, rating)

}

//Modify ...
func (u RatingsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var rating models.Rating

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&rating).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("rating not found.")
	}

	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&rating); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&rating)

	c.JSON(http.StatusOK, rating)
}

//Remove ...
func (u RatingsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var rating models.Rating

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&rating).Error; err != nil {
		log.Println("rating not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "rating not found.",
		})
	} else {
		log.Println("rating returned")
		c.JSON(http.StatusOK, rating)
	}

}
