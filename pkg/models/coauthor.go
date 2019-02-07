package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//CoAuthor ...
type CoAuthor struct {
	ID          uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt   time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	CoAuthorRef uuid.UUID  `sql:"type:uuid" json:"coauthor_ref"`
	PostRef     uuid.UUID  `sql:"type:uuid" json:"post_ref"`
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
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if coauthor, ok := p.Source.(CoAuthor); ok {
					return coauthor.ID, nil
				}
				return nil, nil
			},
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if coauthor, ok := p.Source.(CoAuthor); ok {
					return coauthor.CreatedAt, nil
				}
				return nil, nil
			},
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if coauthor, ok := p.Source.(CoAuthor); ok {
					return coauthor.UpdatedAt, nil
				}
				return nil, nil
			},
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "deletion date",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if coauthor, ok := p.Source.(CoAuthor); ok {
					return coauthor.DeletedAt, nil
				}
				return nil, nil
			},
		},
		"CoAuthorRef": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "CoAuthor Reference",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if coauthor, ok := p.Source.(CoAuthor); ok {
					return coauthor.CoAuthorRef, nil
				}
				return nil, nil
			},
		},
		"PostRef": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Post Reference",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if coauthor, ok := p.Source.(CoAuthor); ok {
					return coauthor.PostRef, nil
				}
				return nil, nil
			},
		},
	},
})
