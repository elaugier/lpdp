package server

import (
	"github.com/elaugier/lpdp/pkg/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//NewRouter ...
func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	router.Use(cors.New(corsConfig))

	achievements := new(controllers.AchievementsController)
	router.GET("/achievements/:id", achievements.Get)
	router.GET("/achievements", achievements.List)
	router.POST("/achievements", achievements.Add)
	router.PUT("/achievements/:id", achievements.Modify)
	router.DELETE("/achievements/:id", achievements.Remove)

	alerts := new(controllers.AlertsController)
	router.GET("/alerts/:id", alerts.Get)
	router.GET("/alerts", alerts.List)
	router.POST("/alerts", alerts.Add)
	router.PUT("/alerts/:id", alerts.Modify)
	router.DELETE("/alerts/:id", alerts.Remove)

	alertactions := new(controllers.AlertActionsController)
	router.GET("/alertactions/:id", alertactions.Get)
	router.GET("/alertactions", alertactions.List)
	router.POST("/alertactions", alertactions.Add)
	router.PUT("/alertactions/:id", alertactions.Modify)
	router.DELETE("/alertactions/:id", alertactions.Remove)

	users := new(controllers.UsersController)
	router.GET("/users/:id", users.Get)
	router.GET("/users", users.List)
	router.POST("/users", users.Add)
	router.PUT("/users/:id", users.Modify)
	router.DELETE("/users/:id", users.Remove)

	return router
}
