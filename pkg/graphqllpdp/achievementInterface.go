package graphqllpdp

import (
	"github.com/graphql-go/graphql"
)

func achievementInterface() *graphql.Interface {
	return graphql.NewInterface(graphql.InterfaceConfig{
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
			"Title": &graphql.Field{
				Type:        graphql.String,
				Description: "title of achievement",
			},
			"OwnerRef": &graphql.Field{
				Type:        graphql.String,
				Description: "owner of achievement",
			},
		},
	})
}
