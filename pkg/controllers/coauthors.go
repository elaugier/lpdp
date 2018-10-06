package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//CoAuthorsController ...
type CoAuthorsController struct{}

//Get ...
func (u CoAuthorsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var coauthor models.CoAuthor

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&coauthor).Error; err != nil {
		log.Println("coauthor not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "coauthor not found.",
		})
	} else {
		log.Println("coauthor returned")
		c.JSON(http.StatusOK, coauthor)
	}

}

//List ...
func (u CoAuthorsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var coauthors []models.CoAuthor

	conn := db.GetInstance()
	if err := conn.Find(&coauthors).Error; err != nil {
		log.Println("coauthor not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "coauthor not found.",
		})
	} else {
		log.Println("coauthor returned")
		c.JSON(http.StatusOK, coauthors)
	}

}

//Add ...
func (u CoAuthorsController) Add(c *gin.Context) {

	var coauthor models.CoAuthor

	if err := c.ShouldBindJSON(&coauthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&coauthor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&coauthor)

	c.JSON(http.StatusCreated, coauthor)

}

//Modify ...
func (u CoAuthorsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var coauthor models.CoAuthor

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&coauthor).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("coauthor not found.")
	}

	if err := c.ShouldBindJSON(&coauthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&coauthor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&coauthor)

	c.JSON(http.StatusOK, coauthor)
}

//Remove ...
func (u CoAuthorsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var coauthor models.CoAuthor

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&coauthor).Error; err != nil {
		log.Println("coauthor not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "coauthor not found.",
		})
	} else {
		log.Println("coauthor returned")
		c.JSON(http.StatusOK, coauthor)
	}

}
