package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//BooksController ...
type BooksController struct{}

//Get ...
func (u BooksController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var book models.Book

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&book).Error; err != nil {
		log.Println("book not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "book not found.",
		})
	} else {
		log.Println("book returned")
		c.JSON(http.StatusOK, book)
	}

}

//List ...
func (u BooksController) List(c *gin.Context) {

	log := logs.GetInstance()

	var books []models.Book

	conn := db.GetInstance()
	if err := conn.Find(&books).Error; err != nil {
		log.Println("book not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "book not found.",
		})
	} else {
		log.Println("book returned")
		c.JSON(http.StatusOK, books)
	}

}

//Add ...
func (u BooksController) Add(c *gin.Context) {

	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&book)

	c.JSON(http.StatusCreated, book)

}

//Modify ...
func (u BooksController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var book models.Book

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&book).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("book not found.")
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&book)

	c.JSON(http.StatusOK, book)
}

//Remove ...
func (u BooksController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var book models.Book

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&book).Error; err != nil {
		log.Println("book not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "book not found.",
		})
	} else {
		log.Println("book returned")
		c.JSON(http.StatusOK, book)
	}

}
