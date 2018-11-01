package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//BadgeType ...
type BadgeType struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Title     string    `gorm:"type:varchar(200);unique_index"`
	Badges    []Badge   `gorm:"foreignkey:TypeRef"`
}

//TableName ...
func (BadgeType) TableName() string {
	return "badgetypes"
}

//BadgeTypeT ...
var BadgeTypeT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "BadgeType",
	Description: "Badge Type",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "badge type id",
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date of badge type",
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date of badge type",
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "deletion date of badge type",
		},
		"Title": &graphql.Field{
			Type:        graphql.String,
			Description: "title of badge type",
		},
	},
})
