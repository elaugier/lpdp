package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//JudgingPanelMembersController ...
type JudgingPanelMembersController struct{}

//Get ...
func (u JudgingPanelMembersController) Get(c *gin.Context) {

	log := logs.GetInstance()

	log.Println("try to retieve 'id' in url path")

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var judgingpanelmember models.JudgingPanelMember

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&judgingpanelmember).Error; err != nil {
		log.Println("judgingpanelmember not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "judgingpanelmember not found.",
		})
	} else {
		log.Println("judgingpanelmember returned")
		c.JSON(http.StatusOK, judgingpanelmember)
	}

}

//List ...
func (u JudgingPanelMembersController) List(c *gin.Context) {

	log := logs.GetInstance()

	var judgingpanelmembers []models.JudgingPanelMember

	conn := db.GetInstance()
	if err := conn.Find(&judgingpanelmembers).Error; err != nil {
		log.Println("judgingpanelmember not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "judgingpanelmember not found.",
		})
	} else {
		log.Println("judgingpanelmember returned")
		c.JSON(http.StatusOK, judgingpanelmembers)
	}

}

//Add ...
func (u JudgingPanelMembersController) Add(c *gin.Context) {

	var judgingpanelmember models.JudgingPanelMember

	if err := c.ShouldBindJSON(&judgingpanelmember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&judgingpanelmember); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn := db.GetInstance()
	conn.Save(&judgingpanelmember)

	c.JSON(http.StatusCreated, judgingpanelmember)

}

//Modify ...
func (u JudgingPanelMembersController) Modify(c *gin.Context) {

	log := logs.GetInstance()

	var judgingpanelmember models.JudgingPanelMember

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).First(&judgingpanelmember).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("judgingpanelmember not found.")
	}

	if err := c.ShouldBindJSON(&judgingpanelmember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&judgingpanelmember); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&judgingpanelmember)

	c.JSON(http.StatusOK, judgingpanelmember)
}

//Remove ...
func (u JudgingPanelMembersController) Remove(c *gin.Context) {

	log := logs.GetInstance()

	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}

	var judgingpanelmember models.JudgingPanelMember

	conn := db.GetInstance()

	if err = conn.Where("id = ?", id).Delete(&judgingpanelmember).Error; err != nil {
		log.Println("judgingpanelmember not found.")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "judgingpanelmember not found.",
		})
	} else {
		log.Println("judgingpanelmember returned")
		c.JSON(http.StatusOK, judgingpanelmember)
	}

}
