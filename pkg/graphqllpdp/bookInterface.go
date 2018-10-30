package graphqllpdp

import (
	"github.com/graphql-go/graphql"
)

func bookInterface() *graphql.Interface {
	return graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "Book",
		Description: "Book",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "book id",
			},
			"CreatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "creation date of book",
			},
			"UpdatedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "update date of book",
			},
			"DeletedAt": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "deletion date of book",
			},
			"Title": &graphql.Field{
				Type:        graphql.String,
				Description: "title of book",
			},
			"Summary": &graphql.Field{
				Type:        graphql.String,
				Description: "title of book",
			},
			"Author": &graphql.Field{
				Type:        graphql.String,
				Description: "author's uuid of book",
			},
		},
	})
}
