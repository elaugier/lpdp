package graphqllpdp

import (
	"github.com/graphql-go/graphql"
)

func badIPAddressInterface() *graphql.Interface {
	return graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "BadIPAddress",
		Description: "Bad IP Address",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "bad ip address id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of bad ip address",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of bad ip address",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of bad ip address",
			},
			"IPAddress": &graphql.Field{
				Type:        graphql.String,
				Description: "ip address of bad ip address",
			},
			"Reason": &graphql.Field{
				Type:        graphql.String,
				Description: "reason of bad ip address",
			},
			"Expires": &graphql.Field{
				Type:        graphql.DateTime,
				Description: "date expiration of bad ip address",
			},
		},
	})
}
