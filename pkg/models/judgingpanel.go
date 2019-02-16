package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

//JudgingPanel ...
type JudgingPanel struct {
	ID                  uuid.UUID            `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt           time.Time            `gorm:"not null" json:"created_at"`
	UpdatedAt           time.Time            `gorm:"not null" json:"updated_at"`
	DeletedAt           *time.Time           `json:"deleted_at"`
	Title               string               `gorm:"type:varchar(100);not null" json:"title"`
	JudgingPanelMembers []JudgingPanelMember `gorm:"foreignkey:JudgingPanelRef"`
}

//TableName ...
func (JudgingPanel) TableName() string {
	return "judgingpanels"
}

//JudgingPanelT ...
var JudgingPanelT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "JudgingPanel",
	Description: "Judging Panel",
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
	},
})
