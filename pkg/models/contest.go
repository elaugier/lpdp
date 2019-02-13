package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

//Contest ...
type Contest struct {
	ID            uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt     time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	Title         string     `gorm:"type:varchar(200)" json:"last_name" binding:"required"`
	ContestRounds []ContestRound
}

//TableName ...
func (Contest) TableName() string {
	return "contests"
}

//ContestT ...
var ContestT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Contest",
	Description: "Contest",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "id",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contest, ok := p.Source.(Contest); ok {
					return contest.ID, nil
				}
				return nil, nil
			},
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contest, ok := p.Source.(Contest); ok {
					return contest.CreatedAt, nil
				}
				return nil, nil
			},
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contest, ok := p.Source.(Contest); ok {
					return contest.UpdatedAt, nil
				}
				return nil, nil
			},
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "deletion date",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contest, ok := p.Source.(Contest); ok {
					return contest.DeletedAt, nil
				}
				return nil, nil
			},
		},
		"Title": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Contest title",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contest, ok := p.Source.(Contest); ok {
					return contest.Title, nil
				}
				return nil, nil
			},
		},
	},
})
