package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//Post ...
type Post struct {
	ID                  uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt           time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt           time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt           *time.Time `json:"deleted_at"`
	PublishedAt         *time.Time `json:"published_at"`
	ApprovedAt          *time.Time `json:"approved_at"`
	ApprouvedBy         uuid.UUID  `sql:"type:uuid" json:"approved_by"`
	ApprouvedByUserName string     `gorm:"type:varchar(100)" json:"approved_by_username"`
	Title               string     `gorm:"type:varchar(255);not null" json:"title"`
	Summary             string     `sql:"type:text" json:"summary"`
	Content             string     `sql:"type:text" json:"content"`
	ViewsCount          int        `json:"views_count"`
	AuthorRef           uuid.UUID  `sql:"type:uuid" json:"author_ref"`
	Likes               []Like     `gorm:"foreignkey:PostRef"`
}

//TableName ...
func (Post) TableName() string {
	return "posts"
}

//PostT ...
var PostT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Post",
	Description: "Post",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Post Identifier",
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "Creation Date of Post",
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "Update Date of Post",
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "Deletion Date of Post",
		},
		"Title": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Post Title",
		},
	},
})
