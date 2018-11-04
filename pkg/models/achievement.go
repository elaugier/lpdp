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
	Rank      int       `gorm:"type:integer" json:"rank"`
	Score     int       `gorm:"type:integer" json:"score"`
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
			Description: "Achievement Identifier",
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "Creation Date of Achievement",
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "Update Date of Achievement",
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "Deletion Date of Achievement",
		},
		"Title": &graphql.Field{
			Type:        graphql.String,
			Description: "Achievement Title",
		},
		"Rank": &graphql.Field{
			Type:        graphql.Int,
			Description: "Rank (optionnal)",
		},
		"Score": &graphql.Field{
			Type:        graphql.Int,
			Description: "Score (optionnal)",
		},
		"OwnerRef": &graphql.Field{
			Type:        graphql.String,
			Description: "Owner of Achievement",
		},
	},
})
