package graphqllpdp

import (
	"github.com/graphql-go/graphql"
)

func tagInterface() *graphql.Interface {
	return graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "BookPart",
		Description: "Book part",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "tag id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of tag",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of tag",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of tag",
			},
			"Code": &graphql.Field{
				Type:        graphql.String,
				Description: "code of tag",
			},
			"Label": &graphql.Field{
				Type:        graphql.String,
				Description: "label of tag",
			},
			"Enable": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "availability of tag",
			},
			"StartedAt": &graphql.Field{
				Type:        graphql.DateTime,
				Description: "start date of availability of tag",
			},
			"ClosedAt": &graphql.Field{
				Type:        graphql.DateTime,
				Description: "end date of availability of tag",
			},
			"Type": &graphql.Field{
				Type:        graphql.String,
				Description: "type of tag",
			},
			"Mature": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "tag maturity",
			},
		},
	})
}
