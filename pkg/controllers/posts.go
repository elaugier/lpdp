package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//PostsController ...
type PostsController struct{}

//Get ...
func (u PostsController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var post models.Post

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&post).Error; err != nil {
		log.Println("post not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "post not found.",
		})
	} else {
		log.Println("post returned")
		c.JSON(http.StatusOK, post)
	}

}

//List ...
func (u PostsController) List(c *gin.Context) {

	log := logs.GetInstance()

	var posts []models.Post

	conn := db.GetInstance()
	if err := conn.Find(&posts).Error; err != nil {
		log.Println("post not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "post not found.",
		})
	} else {
		log.Println("post returned")
		c.JSON(http.StatusOK, posts)
	}

}

//Add ...
func (u PostsController) Add(c *gin.Context) {

	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&post)

	c.JSON(http.StatusCreated, post)

}

//Modify ...
func (u PostsController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var post models.Post

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&post).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("post not found.")
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&post)

	c.JSON(http.StatusOK, post)
}

//Remove ...
func (u PostsController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var post models.Post

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&post).Error; err != nil {
		log.Println("post not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "post not found.",
		})
	} else {
		log.Println("post returned")
		c.JSON(http.StatusOK, post)
	}

}
