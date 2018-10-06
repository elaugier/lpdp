package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//BookPartsController ...
type BookPartsController struct{}

//Get ...
func (u BookPartsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var bookpart models.BookPart

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&bookpart).Error; err != nil {
		log.Println("bookpart not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "bookpart not found.",
		})
	} else {
		log.Println("bookpart returned")
		c.JSON(http.StatusOK, bookpart)
	}

}

//List ...
func (u BookPartsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var bookparts []models.BookPart

	conn := db.GetInstance()
	if err := conn.Find(&bookparts).Error; err != nil {
		log.Println("bookpart not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "bookpart not found.",
		})
	} else {
		log.Println("bookpart returned")
		c.JSON(http.StatusOK, bookparts)
	}

}

//Add ...
func (u BookPartsController) Add(c *gin.Context) {

	var bookpart models.BookPart

	if err := c.ShouldBindJSON(&bookpart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&bookpart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&bookpart)

	c.JSON(http.StatusCreated, bookpart)

}

//Modify ...
func (u BookPartsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var bookpart models.BookPart

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&bookpart).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("bookpart not found.")
	}

	if err := c.ShouldBindJSON(&bookpart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&bookpart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&bookpart)

	c.JSON(http.StatusOK, bookpart)
}

//Remove ...
func (u BookPartsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var bookpart models.BookPart

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&bookpart).Error; err != nil {
		log.Println("bookpart not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "bookpart not found.",
		})
	} else {
		log.Println("bookpart returned")
		c.JSON(http.StatusOK, bookpart)
	}

}
