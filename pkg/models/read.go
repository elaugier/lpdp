package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//Read ...
type Read struct {
	ID         uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	PostRef    uuid.UUID `sql:"type:uuid" json:"post_ref"`
	CommentRef uuid.UUID `sql:"type:uuid" json:"comment_ref"`
	UserRef    uuid.UUID `sql:"type:uuid" json:"user_ref"`
}

//TableName ...
func (Read) TableName() string {
	return "likes"
}

//ReadT ...
var ReadT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Read",
	Description: "Read",
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
		"PostRef": &graphql.Field{
			Type:        graphql.String,
			Description: "reference of post",
		},
		"CommentRef": &graphql.Field{
			Type:        graphql.String,
			Description: "reference of comment",
		},
		"UserRef": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "reference of user",
		},
	},
})
