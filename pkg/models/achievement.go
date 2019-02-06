package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

//Achievement ...
type Achievement struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Title     string     `gorm:"type:varchar(100);not null" json:"title"`
	Rank      int        `gorm:"type:integer" json:"rank"`
	Score     int        `gorm:"type:integer" json:"score"`
	OwnerRef  uuid.UUID  `sql:"type:uuid;not null" json:"owner_ref"`
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
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if achievement, ok := p.Source.(Achievement); ok {
					return achievement.ID, nil
				}
				return nil, nil
			},
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "Creation Date of Achievement",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if achievement, ok := p.Source.(Achievement); ok {
					return achievement.CreatedAt, nil
				}
				return nil, nil
			},
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "Update Date of Achievement",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if achievement, ok := p.Source.(Achievement); ok {
					return achievement.UpdatedAt, nil
				}
				return nil, nil
			},
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "Deletion Date of Achievement",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if achievement, ok := p.Source.(Achievement); ok {
					return achievement.DeletedAt, nil
				}
				return nil, nil
			},
		},
		"Title": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Achievement Title",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if achievement, ok := p.Source.(Achievement); ok {
					return achievement.Title, nil
				}
				return nil, nil
			},
		},
		"Rank": &graphql.Field{
			Type:        graphql.Int,
			Description: "Rank (optionnal)",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if achievement, ok := p.Source.(Achievement); ok {
					return achievement.Rank, nil
				}
				return nil, nil
			},
		},
		"Score": &graphql.Field{
			Type:        graphql.Int,
			Description: "Score (optionnal)",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if achievement, ok := p.Source.(Achievement); ok {
					return achievement.Score, nil
				}
				return nil, nil
			},
		},
		"OwnerRef": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Owner of Achievement",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if achievement, ok := p.Source.(Achievement); ok {
					return achievement.OwnerRef, nil
				}
				return nil, nil
			},
		},
	},
})
