package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//CorrectionRequestActionsController ...
type CorrectionRequestActionsController struct{}

//Get ...
func (u CorrectionRequestActionsController) Get(c *gin.Context) {
	log := logs.GetInstance()
	log.Println("try to retieve 'id' in url path")
	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(400, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}
	var i models.CorrectionRequestAction
	conn := db.GetInstance()
	if err = conn.Where("id = ?", id).First(&i).Error; err != nil {
		log.Println("action not found.")
		c.JSON(404, gin.H{
			"msg": "action not found.",
		})
	} else {
		log.Println("action returned")
		c.JSON(200, i)
	}
}

//List ...
func (u CorrectionRequestActionsController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u CorrectionRequestActionsController) Add(c *gin.Context) {
	var json models.CorrectionRequestAction
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u CorrectionRequestActionsController) Modify(c *gin.Context) {
	var json models.CorrectionRequestAction
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u CorrectionRequestActionsController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
