package graphqllpdp

import (
	"github.com/graphql-go/graphql"
)

func badgeTypeInterface() *graphql.Interface {
	return graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "BadgeType",
		Description: "Badge Type",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "badge type id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of badge type",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of badge type",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of badge type",
			},
			"Title": &graphql.Field{
				Type:        graphql.String,
				Description: "title of badge type",
			},
		},
	})
}
