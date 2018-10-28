package graphqllpdp

import (
	"github.com/graphql-go/graphql"
)

func badgeInterface() *graphql.Interface {
	return graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "Badge",
		Description: "Badge for LPDP Member",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "badge id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of badge",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of badge",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of badge",
			},
			"Message": &graphql.Field{
				Type:        graphql.String,
				Description: "message of badge",
			},
			"TypeRef": &graphql.Field{
				Type:        graphql.String,
				Description: "type of badge",
			},
			"OwnerRef": &graphql.Field{
				Type:        graphql.String,
				Description: "owner of badge",
			},
		},
	})
}
