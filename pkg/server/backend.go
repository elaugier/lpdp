package server

import (
	"log"
	"net/http"

	"github.com/elaugier/lpdp/pkg/controllers"
	"github.com/elaugier/lpdp/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

//BackendRouter ...
func BackendRouter(logger *log.Logger) http.Handler {
	logger.Println("Create new backend router")
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middlewares.Identification(logger))
	router.Use(middlewares.RequestID(logger))

	//corsConfig := cors.DefaultConfig()
	//router.Use(cors.New(corsConfig))

	router.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server Backend",
			},
		)
	})

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
