package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

//AlertAction ...
type AlertAction struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Title     string     `gorm:"type:varchar(100)" json:"title"`
	Details   string     `gorm:"type:varchar(2000)" json:"details"`
	AlertRef  string     `sql:"type:uuid" json:"alert_ref"`
	UserRef   string     `sql:"type:uuid" json:"user_ref"`
}

//TableName ...
func (AlertAction) TableName() string {
	return "alertactions"
}

//AlertActionT ...
var AlertActionT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "AlertAction",
	Description: "Alert action for LPDP Alert",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "alert action id",
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date of alert action",
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date of alert action",
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "deletion date of alert action",
		},
		"Title": &graphql.Field{
			Type:        graphql.String,
			Description: "title of alert action",
		},
		"Details": &graphql.Field{
			Type:        graphql.String,
			Description: "Details of alert action",
		},
		"AlertRef": &graphql.Field{
			Type:        graphql.String,
			Description: "alert targeted by the alert action",
		},
		"UserRef": &graphql.Field{
			Type:        graphql.String,
			Description: "owner of alert action",
		},
	},
})
