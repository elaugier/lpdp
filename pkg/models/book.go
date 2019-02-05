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
			Type:        graphql.DateTime,
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
