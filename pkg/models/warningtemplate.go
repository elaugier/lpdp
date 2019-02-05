package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//WarningTemplate ...
type WarningTemplate struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `gorm:"type:varchar(200);unique_index" json:"name"`
	Content   string     `sql:"type:text" json:"content"`
	Warnings  []Warning  `gorm:"foreignkey:WarningTemplateRef"`
}

//TableName ...
func (WarningTemplate) TableName() string {
	return "warningtemplates"
}

//WarningTemplateT ...
var WarningTemplateT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "WarningTemplate",
	Description: "Warning Template",
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
			Type:        graphql.DateTime,
			Description: "deletion date",
		},
		"Name": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Name",
		},
		"Content": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Content",
		},
	},
})
