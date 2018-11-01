package queries

import (
	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/graphql-go/graphql"
)

// GetRootFields returns all the available queries.
func GetRootFields() graphql.Fields {
	return graphql.Fields{
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
	}
}
