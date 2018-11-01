package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

//WarningAction ...
type WarningAction struct {
	ID         uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	WarningRef uuid.UUID `json:"warning_ref"`
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
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "deletion date",
		},
		"WarningRef": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Warning Reference",
		},
	},
})
