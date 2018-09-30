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
	router.Use(middlewares.Authentication(logger))

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

	comments := new(controllers.CommentsController)
	router.GET("/comments/:id", comments.Get)
	router.GET("/comments", comments.List)
	router.POST("/comments", comments.Add)
	router.PUT("/comments/:id", comments.Modify)
	router.DELETE("/comments/:id", comments.Remove)

	contacts := new(controllers.ContactsController)
	router.GET("/contacts/:id", contacts.Get)
	router.GET("/contacts", contacts.List)
	router.POST("/contacts", contacts.Add)
	router.PUT("/contacts/:id", contacts.Modify)
	router.DELETE("/contacts/:id", contacts.Remove)

	contests := new(controllers.ContestsController)
	router.GET("/contests/:id", contests.Get)
	router.GET("/contests", contests.List)
	router.POST("/contests", contests.Add)
	router.PUT("/contests/:id", contests.Modify)
	router.DELETE("/contests/:id", contests.Remove)

	contestrounds := new(controllers.ContestRoundsController)
	router.GET("/contestrounds/:id", contestrounds.Get)
	router.GET("/contestrounds", contestrounds.List)
	router.POST("/contestrounds", contestrounds.Add)
	router.PUT("/contestrounds/:id", contestrounds.Modify)
	router.DELETE("/contestrounds/:id", contestrounds.Remove)

	correctionrequests := new(controllers.CorrectionRequestsController)
	router.GET("/correctionrequests/:id", correctionrequests.Get)
	router.GET("/correctionrequests", correctionrequests.List)
	router.POST("/correctionrequests", correctionrequests.Add)
	router.PUT("/correctionrequests/:id", correctionrequests.Modify)
	router.DELETE("/correctionrequests/:id", correctionrequests.Remove)

	correctionrequestactions := new(controllers.CorrectionRequestActionsController)
	router.GET("/correctionrequestactions/:id", correctionrequestactions.Get)
	router.GET("/correctionrequestactions", correctionrequestactions.List)
	router.POST("/correctionrequestactions", correctionrequestactions.Add)
	router.PUT("/correctionrequestactions/:id", correctionrequestactions.Modify)
	router.DELETE("/correctionrequestactions/:id", correctionrequestactions.Remove)

	iphistories := new(controllers.IPHistoriesController)
	router.GET("/iphistories/:id", iphistories.Get)
	router.GET("/iphistories", iphistories.List)
	router.POST("/iphistories", iphistories.Add)
	router.PUT("/iphistories/:id", iphistories.Modify)
	router.DELETE("/iphistories/:id", iphistories.Remove)

	judgingpanelmembers := new(controllers.JudgingPanelMembersController)
	router.GET("/judgingpanelmembers/:id", judgingpanelmembers.Get)
	router.GET("/judgingpanelmembers", judgingpanelmembers.List)
	router.POST("/judgingpanelmembers", judgingpanelmembers.Add)
	router.PUT("/judgingpanelmembers/:id", judgingpanelmembers.Modify)
	router.DELETE("/judgingpanelmembers/:id", judgingpanelmembers.Remove)

	judgingpanels := new(controllers.JudgingPanelsController)
	router.GET("/judgingpanels/:id", judgingpanels.Get)
	router.GET("/judgingpanels", judgingpanels.List)
	router.POST("/judgingpanels", judgingpanels.Add)
	router.PUT("/judgingpanels/:id", judgingpanels.Modify)
	router.DELETE("/judgingpanels/:id", judgingpanels.Remove)

	licenses := new(controllers.LicensesController)
	router.GET("/licenses/:id", licenses.Get)
	router.GET("/licenses", licenses.List)
	router.POST("/licenses", licenses.Add)
	router.PUT("/licenses/:id", licenses.Modify)
	router.DELETE("/licenses/:id", licenses.Remove)

	likes := new(controllers.LikesController)
	router.GET("/likes/:id", likes.Get)
	router.GET("/likes", likes.List)
	router.POST("/likes", likes.Add)
	router.PUT("/likes/:id", likes.Modify)
	router.DELETE("/likes/:id", likes.Remove)

	login := new(controllers.LoginController)
	router.GET("/auth", login.Get)
	router.GET("/auths", login.List)
	router.POST("/auth", login.Login)
	router.DELETE("/auth", login.Logout)

	messages := new(controllers.MessagesController)
	router.GET("/messages/:id", messages.Get)
	router.GET("/messages", messages.List)
	router.POST("/messages", messages.Add)
	router.PUT("/messages/:id", messages.Modify)
	router.DELETE("/messages/:id", messages.Remove)

	notifications := new(controllers.NotificationsController)
	router.GET("/notifications/:id", notifications.Get)
	router.GET("/notifications", notifications.List)
	router.POST("/notifications", notifications.Add)
	router.PUT("/notifications/:id", notifications.Modify)
	router.DELETE("/notifications/:id", notifications.Remove)

	payments := new(controllers.PaymentsController)
	router.GET("/payments/:id", payments.Get)
	router.GET("/payments", payments.List)
	router.POST("/payments", payments.Add)
	router.PUT("/payments/:id", payments.Modify)
	router.DELETE("/payments/:id", payments.Remove)

	posts := new(controllers.PostsController)
	router.GET("/posts/:id", posts.Get)
	router.GET("/posts", posts.List)
	router.POST("/posts", posts.Add)
	router.PUT("/posts/:id", posts.Modify)
	router.DELETE("/posts/:id", posts.Remove)

	ratings := new(controllers.RatingsController)
	router.GET("/ratings/:id", ratings.Get)
	router.GET("/ratings", ratings.List)
	router.POST("/ratings", ratings.Add)
	router.PUT("/ratings/:id", ratings.Modify)
	router.DELETE("/ratings/:id", ratings.Remove)

	reads := new(controllers.ReadsController)
	router.GET("/reads/:id", reads.Get)
	router.GET("/reads", reads.List)
	router.POST("/reads", reads.Add)
	router.PUT("/reads/:id", reads.Modify)
	router.DELETE("/reads/:id", reads.Remove)

	requests := new(controllers.RequestsController)
	router.GET("/requests/:id", requests.Get)
	router.GET("/requests", requests.List)
	router.POST("/requests", requests.Add)
	router.PUT("/requests/:id", requests.Modify)
	router.DELETE("/requests/:id", requests.Remove)

	sections := new(controllers.SectionsController)
	router.GET("/sections/:id", sections.Get)
	router.GET("/sections", sections.List)
	router.POST("/sections", sections.Add)
	router.PUT("/sections/:id", sections.Modify)
	router.DELETE("/sections/:id", sections.Remove)

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

	votes := new(controllers.VotesController)
	router.GET("/votes/:id", votes.Get)
	router.GET("/votes", votes.List)
	router.POST("/votes", votes.Add)
	router.PUT("/votes/:id", votes.Modify)
	router.DELETE("/votes/:id", votes.Remove)

	votings := new(controllers.VotingsController)
	router.GET("/votings/:id", votings.Get)
	router.GET("/votings", votings.List)
	router.POST("/votings", votings.Add)
	router.PUT("/votings/:id", votings.Modify)
	router.DELETE("/votings/:id", votings.Remove)

	warnings := new(controllers.WarningsController)
	router.GET("/warnings/:id", warnings.Get)
	router.GET("/warnings", warnings.List)
	router.POST("/warnings", warnings.Add)
	router.PUT("/warnings/:id", warnings.Modify)
	router.DELETE("/warnings/:id", warnings.Remove)

	warningtemplates := new(controllers.WarningTemplatesController)
	router.GET("/warningtemplates/:id", warningtemplates.Get)
	router.GET("/warningtemplates", warningtemplates.List)
	router.POST("/warningtemplates", warningtemplates.Add)
	router.PUT("/warningtemplates/:id", warningtemplates.Modify)
	router.DELETE("/warningtemplates/:id", warningtemplates.Remove)

	return router
}
