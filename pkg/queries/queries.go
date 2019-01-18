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
			Type:    graphql.NewList(models.AchievementT),
			Resolve: db.GetAchievementQuery,
		},
		"Activity": &graphql.Field{
			Type:    graphql.NewList(models.ActivityT),
			Resolve: db.GetActivityQuery,
		},
		"Alert": &graphql.Field{
			Type:    graphql.NewList(models.AlertT),
			Resolve: db.GetAlertQuery,
		},
		"AlertAction": &graphql.Field{
			Type:    graphql.NewList(models.AlertActionT),
			Resolve: db.GetAlertActionQuery,
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
		"Section": &graphql.Field{
			Type:    graphql.NewList(models.SectionT),
			Resolve: db.GetSectionQuery,
		},
		"Tag": &graphql.Field{
			Type:    graphql.NewList(models.TagT),
			Resolve: db.GetTagQuery,
		},
		"Warning": &graphql.Field{
			Type: graphql.NewList(models.WarningT),
		},
		"WarningAction": &graphql.Field{
			Type: graphql.NewList(models.WarningActionT),
		},
		"WarningTemplate": &graphql.Field{
			Type: graphql.NewList(models.WarningTemplateT),
		},
	}
}
