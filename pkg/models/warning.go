package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

//Warning ...
type Warning struct {
	ID                 uuid.UUID       `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt          time.Time       `gorm:"not null" json:"created_at"`
	UpdatedAt          time.Time       `gorm:"not null" json:"updated_at"`
	DeletedAt          *time.Time      `json:"deleted_at"`
	Status             string          `gorm:"type:varchar(10);not null" json:"status"`
	WarningTemplateRef uuid.UUID       `sql:"type:uuid" json:"warning_template_ref"`
	WarningActions     []WarningAction `gorm:"foreignkey:WarningRef"`
}

//TableName ...
func (Warning) TableName() string {
	return "warnings"
}

//WarningT ...
var WarningT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Warning",
	Description: "Warning",
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
		"Status": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Status",
		},
		"WarningTemplateRef": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Warning Template Reference",
		},
	},
})
