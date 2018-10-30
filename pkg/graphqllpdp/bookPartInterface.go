package graphqllpdp

import (
	"github.com/graphql-go/graphql"
)

func bookPartInterface() *graphql.Interface {
	return graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "BookPart",
		Description: "Book part",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "book part id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of book part",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of book part",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of book part",
			},
			"Title": &graphql.Field{
				Type:        graphql.String,
				Description: "title of book part",
			},
			"BookRef": &graphql.Field{
				Type:        graphql.String,
				Description: "book's uuid of book part",
			},
		},
	})
}
