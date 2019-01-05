package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"github.com/elaugier/lpdp/pkg/middlewares"
	"github.com/elaugier/lpdp/pkg/mutations"
	"github.com/elaugier/lpdp/pkg/queries"
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

	h, err := Handler()
	if err == nil {
		router.GET("/graphql", h)
		router.POST("/graphql", h)
	}

	return router
}

// Handler ...
func Handler() (gin.HandlerFunc, error) {

	logger := logs.GetInstance()

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "LpdpQuery",
			Fields:      queries.GetRootFields(),
			Description: "List of all available queries in Lpdp GraphQL API.",
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:        "LpdpMutation",
			Fields:      mutations.GetRootFields(),
			Description: "List of all available mutations in Lpdp GraphQL API.",
		}),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		logger.Printf("error : %s", err)
		return nil, errors.New("Error on loading Graphql Schema")
	}

	// Creates a GraphQL-go HTTP handler with the defined schema
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}, nil
}
