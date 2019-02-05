package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//ExitReason ...
type ExitReason struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Reason    string     `gorm:"not null" json:"title"`
}

//TableName ...
func (ExitReason) TableName() string {
	return "exitreasons"
}

//ExitReasonT ...
var ExitReasonT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ExitReason",
	Description: "Exit Reason of removed accounts",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Identifier",
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "Creation Date",
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "Update Date",
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "Deletion Date",
		},
		"Reason": &graphql.Field{
			Type:        graphql.String,
			Description: "Reason of removed",
		},
	},
})
