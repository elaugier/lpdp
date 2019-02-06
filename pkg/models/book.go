package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//Book ...
type Book struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Title     string     `gorm:"type:varchar(100)" json:"title"`
	Summary   string     `sql:"type:text;" json:"summary"`
	AuthorRef uuid.UUID  `sql:"type:uuid" json:"author_ref"`
	BookParts []BookPart `gorm:"foreignkey:BookRef"`
}

//TableName ...
func (Book) TableName() string {
	return "books"
}

//BookT ...
var BookT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Book",
	Description: "Book",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "book id",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if book, ok := p.Source.(Book); ok {
					return book.ID, nil
				}
				return nil, nil
			},
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date of book",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if book, ok := p.Source.(Book); ok {
					return book.CreatedAt, nil
				}
				return nil, nil
			},
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date of book",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if book, ok := p.Source.(Book); ok {
					return book.UpdatedAt, nil
				}
				return nil, nil
			},
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "deletion date of book",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if book, ok := p.Source.(Book); ok {
					return book.DeletedAt, nil
				}
				return nil, nil
			},
		},
		"Title": &graphql.Field{
			Type:        graphql.String,
			Description: "title of book",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if book, ok := p.Source.(Book); ok {
					return book.Title, nil
				}
				return nil, nil
			},
		},
		"Summary": &graphql.Field{
			Type:        graphql.String,
			Description: "title of book",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if book, ok := p.Source.(Book); ok {
					return book.Summary, nil
				}
				return nil, nil
			},
		},
		"AuthorRef": &graphql.Field{
			Type:        graphql.String,
			Description: "author's uuid of book",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if book, ok := p.Source.(Book); ok {
					return book.AuthorRef, nil
				}
				return nil, nil
			},
		},
	},
})
