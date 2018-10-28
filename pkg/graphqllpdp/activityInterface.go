package graphqllpdp

import (
	"github.com/graphql-go/graphql"
)

func activityInterface() *graphql.Interface {
	return graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "Activity",
		Description: "Activity for LPDP Member",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "activity id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of activity",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of activity",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of activity",
			},
			"Message": &graphql.Field{
				Type:        graphql.String,
				Description: "message of activity",
			},
			"Type": &graphql.Field{
				Type:        graphql.String,
				Description: "type of activity",
			},
			"OwnerRef": &graphql.Field{
				Type:        graphql.String,
				Description: "owner of activity",
			},
		},
	})

}
