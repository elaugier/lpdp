package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//CoAuthor ...
type CoAuthor struct {
	ID          uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	CoAuthorRef uuid.UUID `sql:"type:uuid" json:"coauthor_ref"`
	PostRef     uuid.UUID `sql:"type:uuid" json:"post_ref"`
}

//TableName ...
func (CoAuthor) TableName() string {
	return "coauthors"
}

//CoAuthorT ...
var CoAuthorT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "CoAuthor",
	Description: "CoAuthor",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "id",
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date",
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date",
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "deletion date",
		},
		"CoAuthorRef": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "CoAuthor Reference",
		},
		"PostRef": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Post Reference",
		},
	},
})
