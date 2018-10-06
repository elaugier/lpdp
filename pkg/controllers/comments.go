package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//CommentsController ...
type CommentsController struct{}

//Get ...
func (u CommentsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var comment models.Comment

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&comment).Error; err != nil {
		log.Println("comment not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "comment not found.",
		})
	} else {
		log.Println("comment returned")
		c.JSON(http.StatusOK, comment)
	}

}

//List ...
func (u CommentsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var comments []models.Comment

	conn := db.GetInstance()
	if err := conn.Find(&comments).Error; err != nil {
		log.Println("comment not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "comment not found.",
		})
	} else {
		log.Println("comment returned")
		c.JSON(http.StatusOK, comments)
	}

}

//Add ...
func (u CommentsController) Add(c *gin.Context) {

	var comment models.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&comment)

	c.JSON(http.StatusCreated, comment)

}

//Modify ...
func (u CommentsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var comment models.Comment

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&comment).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("comment not found.")
	}

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&comment)

	c.JSON(http.StatusOK, comment)
}

//Remove ...
func (u CommentsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var comment models.Comment

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&comment).Error; err != nil {
		log.Println("comment not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "comment not found.",
		})
	} else {
		log.Println("comment returned")
		c.JSON(http.StatusOK, comment)
	}

}
