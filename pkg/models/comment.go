package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

//Comment ...
type Comment struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Content   string     `sql:"type:text"`
	AuthorRef uuid.UUID  `sql:"type:uuid" json:"author_ref"`
	Likes     []Like     `gorm:"foreignkey:CommentRef"`
}

//TableName ...
func (Comment) TableName() string {
	return "comments"
}

//CommentT ...
var CommentT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Comment",
	Description: "Comment",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "id",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(Comment); ok {
					return comment.ID, nil
				}
				return nil, nil
			},
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(Comment); ok {
					return comment.CreatedAt, nil
				}
				return nil, nil
			},
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(Comment); ok {
					return comment.UpdatedAt, nil
				}
				return nil, nil
			},
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "deletion date",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(Comment); ok {
					return comment.DeletedAt, nil
				}
				return nil, nil
			},
		},
		"Content": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Content of the comment",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(Comment); ok {
					return comment.Content, nil
				}
				return nil, nil
			},
		},
		"AuthorRef": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Reference to the author",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(Comment); ok {
					return comment.AuthorRef, nil
				}
				return nil, nil
			},
		},
	},
})
