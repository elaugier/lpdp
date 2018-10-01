package controllers

import (
	"net/http"

	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//BadIPAddressesController ...
type BadIPAddressesController struct{}

//Get ...
func (u BadIPAddressesController) Get(c *gin.Context) {
	log := logs.GetInstance()
	log.Println("try to retieve 'id' in url path")
	id, err := uuid.Parse(c.Params.ByName("id"))
	if err != nil {
		log.Println("Cannot get 'id' in url path.")
		c.JSON(400, gin.H{
			"msg": "Cannot get 'id' in url path.",
		})
	}
	var i models.BadIPAddress
	conn := db.GetInstance()
	if err = conn.Where("id = ?", id).First(&i).Error; err != nil {
		log.Println("ip not found.")
		c.JSON(404, gin.H{
			"msg": "ip not found.",
		})
	} else {
		log.Println("ip returned")
		c.JSON(200, i)
	}
}

//List ...
func (u BadIPAddressesController) List(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

//Add ...
func (u BadIPAddressesController) Add(c *gin.Context) {
	var json models.BadIPAddress
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Modify ...
func (u BadIPAddressesController) Modify(c *gin.Context) {
	var json models.BadIPAddress
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Working!")
}

//Remove ...
func (u BadIPAddressesController) Remove(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
