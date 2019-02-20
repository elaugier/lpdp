package queries

import (
	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/graphql-go/graphql"
)

//DefaultItemsPerPage ...
const DefaultItemsPerPage = 50

// GetRootFields returns all the available queries.
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"Achievement": &graphql.Field{
			Type: graphql.NewList(models.AchievementT),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
				"itemsPerPage": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: DefaultItemsPerPage,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: db.GetAchievementQuery,
		},
		"Activity": &graphql.Field{
			Type: graphql.NewList(models.ActivityT),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
				"itemsPerPage": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: DefaultItemsPerPage,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: db.GetActivityQuery,
		},
		"Alert": &graphql.Field{
			Type: graphql.NewList(models.AlertT),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
				"itemsPerPage": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: DefaultItemsPerPage,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: db.GetAlertQuery,
		},
		"AlertAction": &graphql.Field{
			Type: graphql.NewList(models.AlertActionT),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
				"itemsPerPage": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: DefaultItemsPerPage,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: db.GetAlertActionQuery,
		},
		"Badge": &graphql.Field{
			Type: graphql.NewList(models.BadgeT),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
				"itemsPerPage": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: DefaultItemsPerPage,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: db.GetBadgeQuery,
		},
		"BadgeType": &graphql.Field{
			Type: graphql.NewList(models.BadgeTypeT),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
				"itemsPerPage": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: DefaultItemsPerPage,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: db.GetBadgeTypeQuery,
		},
		"BadIPAddress": &graphql.Field{
			Type: graphql.NewList(models.BadIPAddressT),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
				"itemsPerPage": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: DefaultItemsPerPage,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: db.GetBadIPAddressQuery,
		},
		"Book": &graphql.Field{
			Type: graphql.NewList(models.BookT),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
				"itemsPerPage": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: DefaultItemsPerPage,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
		},
		"BookPart": &graphql.Field{
			Type: graphql.NewList(models.BookPartT),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
				"itemsPerPage": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: DefaultItemsPerPage,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
		},
		"CoAuthor": &graphql.Field{
			Type: graphql.NewList(models.CoAuthorT),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
				"itemsPerPage": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: DefaultItemsPerPage,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
		},
		"User": &graphql.Field{
			Type: graphql.NewList(models.UserT),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
				"itemsPerPage": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: DefaultItemsPerPage,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: db.GetUserQuery,
		},
		"Section": &graphql.Field{
			Type: graphql.NewList(models.SectionT),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
				"itemsPerPage": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: DefaultItemsPerPage,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: db.GetSectionQuery,
		},
		"Tag": &graphql.Field{
			Type: graphql.NewList(models.TagT),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
				"itemsPerPage": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: DefaultItemsPerPage,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: db.GetTagQuery,
		},
		"Warning": &graphql.Field{
			Type: graphql.NewList(models.WarningT),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
				"itemsPerPage": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: DefaultItemsPerPage,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
		},
		"WarningAction": &graphql.Field{
			Type: graphql.NewList(models.WarningActionT),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
				"itemsPerPage": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: DefaultItemsPerPage,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
		},
		"WarningTemplate": &graphql.Field{
			Type: graphql.NewList(models.WarningTemplateT),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
				"itemsPerPage": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: DefaultItemsPerPage,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
		},
	}
}
