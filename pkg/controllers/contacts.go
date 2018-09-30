package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//ContactsController ...
type ContactsController struct{}

//Get ...
func (u ContactsController) Get(c *gin.Context) {
	log := logs.GetInstance()
	log.Println("try to retieve 'id' in url path")
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(400, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}
	var i models.Contact
	conn := db.GetInstance()
	if err = conn.Where("id = ?", id).First(&i).Error; err != nil {
		log.Println("contact not found.")
		c.JSON(404, gin.H{
			"msg": "contact not found.",
		})
	} else {
		log.Println("contact returned")
		c.JSON(200, i)
	}
}

//List ...
func (u ContactsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u ContactsController) Add(c *gin.Context) {
	var json models.Contact
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u ContactsController) Modify(c *gin.Context) {
	var json models.Contact
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u ContactsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
