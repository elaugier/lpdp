package graphqllpdp

import (
	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/graphql-go/graphql"
)

var (
	graphqlLpdpSchema graphql.Schema
)

// GetSchema ...
func GetSchema() (graphql.Schema, error) {

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"Achievement": &graphql.Field{
				Type: graphql.NewList(models.AchievementT),
			},
			"Activity": &graphql.Field{
				Type: graphql.NewList(models.ActivityT),
			},
			"Alert": &graphql.Field{
				Type: graphql.NewList(models.AlertT),
			},
			"AlertAction": &graphql.Field{
				Type: graphql.NewList(models.AlertActionT),
			},
			"Badge": &graphql.Field{
				Type: graphql.NewList(models.BadgeT),
			},
			"BadgeType": &graphql.Field{
				Type: graphql.NewList(models.BadgeTypeT),
			},
			"BadIPAddress": &graphql.Field{
				Type: graphql.NewList(models.BadIPAddressT),
			},
			"Book": &graphql.Field{
				Type: graphql.NewList(models.BookT),
			},
			"BookPart": &graphql.Field{
				Type: graphql.NewList(models.BookPartT),
			},
			"CoAuthor": &graphql.Field{
				Type: graphql.NewList(models.CoAuthorT),
			},
			"User": &graphql.Field{
				Type:    graphql.NewList(models.UserT),
				Resolve: db.GetUserQuery,
			},
			"Tag": &graphql.Field{
				Type:    graphql.NewList(models.TagT),
				Resolve: db.GetTagQuery,
			},
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createAchievement": &graphql.Field{
				Type: models.AchievementT,
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
