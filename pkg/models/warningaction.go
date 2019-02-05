package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

//WarningAction ...
type WarningAction struct {
	ID             uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt      time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
	Message        string     `json:"message"`
	ActionType     string     `json:"action_type"`
	WarningRef     uuid.UUID  `gorm:"not null" sql:"type:uuid" json:"warning_ref"`
	OwnerActionRef uuid.UUID  `gorm:"not null" sql:"type:uuid" json:"owner_action_ref"`
}

//TableName ...
func (WarningAction) TableName() string {
	return "warningactions"
}

//WarningActionT ...
var WarningActionT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "WarningAction",
	Description: "Warning Action",
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
		"Message": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Message of the action",
		},
		"ActionType": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Type of action",
		},
		"WarningRef": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Warning Reference",
		},
		"OwnerActionRef": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Action Owner Reference",
		},
	},
})
