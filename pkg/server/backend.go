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

	activities := new(controllers.ActivitiesController)
	router.GET("/activities/:id", activities.Get)
	router.GET("/activities", activities.List)
	router.POST("/activities", activities.Add)
	router.PUT("/activities/:id", activities.Modify)
	router.DELETE("/activities/:id", activities.Remove)

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

	badges := new(controllers.BadgesController)
	router.GET("/badges/:id", badges.Get)
	router.GET("/badges", badges.List)
	router.POST("/badges", badges.Add)
	router.PUT("/badges/:id", badges.Modify)
	router.DELETE("/badges/:id", badges.Remove)

	badgetypes := new(controllers.BadgeTypesController)
	router.GET("/badgetypes/:id", badgetypes.Get)
	router.GET("/badgetypes", badgetypes.List)
	router.POST("/badgetypes", badgetypes.Add)
	router.PUT("/badgetypes/:id", badgetypes.Modify)
	router.DELETE("/badgetypes/:id", badgetypes.Remove)

	badipaddresses := new(controllers.BadIPAddressesController)
	router.GET("/badipaddresses/:id", badipaddresses.Get)
	router.GET("/badipaddresses", badipaddresses.List)
	router.POST("/badipaddresses", badipaddresses.Add)
	router.PUT("/badipaddresses/:id", badipaddresses.Modify)
	router.DELETE("/badipaddresses/:id", badipaddresses.Remove)

	books := new(controllers.BooksController)
	router.GET("/books/:id", books.Get)
	router.GET("/books", books.List)
	router.POST("/books", books.Add)
	router.PUT("/books/:id", books.Modify)
	router.DELETE("/books/:id", books.Remove)

	bookparts := new(controllers.BookPartsController)
	router.GET("/bookparts/:id", bookparts.Get)
	router.GET("/bookparts", bookparts.List)
	router.POST("/bookparts", bookparts.Add)
	router.PUT("/bookparts/:id", bookparts.Modify)
	router.DELETE("/bookparts/:id", bookparts.Remove)

	coauthors := new(controllers.CoauthorsController)
	router.GET("/coauthors/:id", coauthors.Get)
	router.GET("/coauthors", coauthors.List)
	router.POST("/coauthors", coauthors.Add)
	router.PUT("/coauthors/:id", coauthors.Modify)
	router.DELETE("/coauthors/:id", coauthors.Remove)

	tags := new(controllers.TagsController)
	router.GET("/tags/:id", tags.Get)
	router.GET("/tags", tags.List)
	router.POST("/tags", tags.Add)
	router.PUT("/tags/:id", tags.Modify)
	router.DELETE("/tags/:id", tags.Remove)

	users := new(controllers.UsersController)
	router.GET("/users/:id", users.Get)
	router.GET("/users", users.List)
	router.POST("/users", users.Add)
	router.PUT("/users/:id", users.Modify)
	router.DELETE("/users/:id", users.Remove)

	return router
}
