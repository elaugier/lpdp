package graphqllpdp

import (
	"github.com/graphql-go/graphql"
)

var (
	graphqlLpdpSchema graphql.Schema
)

// GetSchema ...
func GetSchema() (graphql.Schema, error) {
	achievementInterface := graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "Achievement",
		Description: "Achievement for LPDP Member",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "achievement id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of achievement",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of achievement",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of achievement",
			},
		},
	})
	userInterface := graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "User",
		Description: "Member of LPDP",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "user id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of user",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of user",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of user",
			},
			"Email": &graphql.Field{
				Type:        graphql.String,
				Description: "email address of user",
			},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"achievement": &graphql.Field{
				Type: achievementInterface,
			},
			"user": &graphql.Field{
				Type: userInterface,
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
