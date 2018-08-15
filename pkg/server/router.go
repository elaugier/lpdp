package server

import (
	"github.com/elaugier/lpdp/pkg/controllers"
	"github.com/gin-gonic/gin"
)

//NewRouter ...
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	users := new(controllers.UsersController)
	router.GET("/users/:id", users.Get)
	router.GET("/users", users.List)
	router.POST("/users", users.Add)
	router.PUT("/users/:id", users.Modify)
	router.DELETE("/users/:id", users.Remove)

	return router
}
