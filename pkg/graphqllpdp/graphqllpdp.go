package graphqllpdp

import (
	"github.com/graphql-go/graphql"
)

var (
	graphqlLpdpSchema graphql.Schema
)

// GetSchema ...
func GetSchema() (graphql.Schema, error) {
	userInterface := graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "User",
		Description: "User of LPDP",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "user id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of user",
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{})
}
