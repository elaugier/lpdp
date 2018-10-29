package graphqllpdp

import (
	"github.com/graphql-go/graphql"
)

var (
	graphqlLpdpSchema graphql.Schema
)

// GetSchema ...
func GetSchema() (graphql.Schema, error) {

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"achievements": &graphql.Field{
				Type: graphql.NewList(achievementInterface()),
			},
			"achievement": &graphql.Field{
				Type: achievementInterface(),
			},
			"activities": &graphql.Field{
				Type: graphql.NewList(activityInterface()),
			},
			"activity": &graphql.Field{
				Type: activityInterface(),
			},
			"alerts": &graphql.Field{
				Type: graphql.NewList(alertInterface()),
			},
			"alert": &graphql.Field{
				Type: alertInterface(),
			},
			"alertactions": &graphql.Field{
				Type: graphql.NewList(alertActionInterface()),
			},
			"alertaction": &graphql.Field{
				Type: alertActionInterface(),
			},
			"badges": &graphql.Field{
				Type: graphql.NewList(badgeInterface()),
			},
			"badge": &graphql.Field{
				Type: badgeInterface(),
			},
			"badgeTypes": &graphql.Field{
				Type: graphql.NewList(badgeTypeInterface()),
			},
			"badgeType": &graphql.Field{
				Type: badgeTypeInterface(),
			},
			"badIPAddresses": &graphql.Field{
				Type: graphql.NewList(badIPAddressInterface()),
			},
			"badIPAddress": &graphql.Field{
				Type: badIPAddressInterface(),
			},
			"users": &graphql.Field{
				Type: graphql.NewList(userInterface()),
			},
			"user": &graphql.Field{
				Type: userInterface(),
			},
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createAchievement": &graphql.Field{
				Type: achievementInterface(),
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"owner_ref": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
}