package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//ReadsController ...
type ReadsController struct{}

//Get ...
func (u ReadsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var read models.Read

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&read).Error; err != nil {
		log.Println("read not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "read not found.",
		})
	} else {
		log.Println("read returned")
		c.JSON(http.StatusOK, read)
	}

}

//List ...
func (u ReadsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var reads []models.Read

	conn := db.GetInstance()
	if err := conn.Find(&reads).Error; err != nil {
		log.Println("read not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "read not found.",
		})
	} else {
		log.Println("read returned")
		c.JSON(http.StatusOK, reads)
	}

}

//Add ...
func (u ReadsController) Add(c *gin.Context) {

	var read models.Read

	if err := c.ShouldBindJSON(&read); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&read); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&read)

	c.JSON(http.StatusCreated, read)

}

//Modify ...
func (u ReadsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var read models.Read

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&read).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("read not found.")
	}

	if err := c.ShouldBindJSON(&read); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&read); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&read)

	c.JSON(http.StatusOK, read)
}

//Remove ...
func (u ReadsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var read models.Read

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&read).Error; err != nil {
		log.Println("read not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "read not found.",
		})
	} else {
		log.Println("read returned")
		c.JSON(http.StatusOK, read)
	}

}
