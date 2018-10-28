package graphqllpdp

import (
	"github.com/graphql-go/graphql"
)

func alertInterface() *graphql.Interface {
	return graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "Alert",
		Description: "Alert for LPDP Post or Comment",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "alert id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of alert",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of alert",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of alert",
			},
			"Type": &graphql.Field{
				Type:        graphql.String,
				Description: "type of alert",
			},
			"Details": &graphql.Field{
				Type:        graphql.String,
				Description: "Details of alert",
			},
			"PostRef": &graphql.Field{
				Type:        graphql.String,
				Description: "post targeted by the alert",
			},
			"CommentRef": &graphql.Field{
				Type:        graphql.String,
				Description: "comment targeted of alert",
			},
			"UserRef": &graphql.Field{
				Type:        graphql.String,
				Description: "owner of alert",
			},
		},
	})
}
