package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

//Achievement ...
type Achievement struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Title     string    `gorm:"type:varchar(100)" json:"title"`
	OwnerRef  uuid.UUID `sql:"type:uuid" json:"owner_ref"`
}

//TableName ...
func (Achievement) TableName() string {
	return "achievements"
}

//AchievementT ...
var AchievementT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Achievement",
	Description: "Achievement for LPDP Member",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "achievement id",
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date of achievement",
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date of achievement",
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "deletion date of achievement",
		},
		"Title": &graphql.Field{
			Type:        graphql.String,
			Description: "title of achievement",
		},
		"OwnerRef": &graphql.Field{
			Type:        graphql.String,
			Description: "owner of achievement",
		},
	},
})
