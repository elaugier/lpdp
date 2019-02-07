package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//BookPart ...
type BookPart struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Title     string     `json:"title"`
	BookRef   uuid.UUID  `sql:"type:uuid" json:"book_ref"`
}

//TableName ...
func (BookPart) TableName() string {
	return "bookparts"
}

//BookPartT ...
var BookPartT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "BookPart",
	Description: "Book part",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "book part id",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if bookpart, ok := p.Source.(BookPart); ok {
					return bookpart.ID, nil
				}
				return nil, nil
			},
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date of book part",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if bookpart, ok := p.Source.(BookPart); ok {
					return bookpart.CreatedAt, nil
				}
				return nil, nil
			},
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date of book part",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if bookpart, ok := p.Source.(BookPart); ok {
					return bookpart.UpdatedAt, nil
				}
				return nil, nil
			},
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "deletion date of book part",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if bookpart, ok := p.Source.(BookPart); ok {
					return bookpart.DeletedAt, nil
				}
				return nil, nil
			},
		},
		"Title": &graphql.Field{
			Type:        graphql.String,
			Description: "title of book part",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if bookpart, ok := p.Source.(BookPart); ok {
					return bookpart.Title, nil
				}
				return nil, nil
			},
		},
		"BookRef": &graphql.Field{
			Type:        graphql.String,
			Description: "book's uuid of book part",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if bookpart, ok := p.Source.(BookPart); ok {
					return bookpart.BookRef, nil
				}
				return nil, nil
			},
		},
	},
})
