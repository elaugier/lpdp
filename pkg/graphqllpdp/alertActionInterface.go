package graphqllpdp

import (
	"github.com/graphql-go/graphql"
)

func alertActionInterface() *graphql.Interface {
	return graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "AlertAction",
		Description: "Alert action for LPDP Alert",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "alert action id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of alert action",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of alert action",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of alert action",
			},
			"Title": &graphql.Field{
				Type:        graphql.String,
				Description: "title of alert action",
			},
			"Details": &graphql.Field{
				Type:        graphql.String,
				Description: "Details of alert action",
			},
			"AlertRef": &graphql.Field{
				Type:        graphql.String,
				Description: "alert targeted by the alert action",
			},
			"UserRef": &graphql.Field{
				Type:        graphql.String,
				Description: "owner of alert action",
			},
		},
	})
}
