package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

//ContestRound ...
type ContestRound struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Order     int        `gorm:"not null" json:"order"`
}

//TableName ...
func (ContestRound) TableName() string {
	return "contestrounds"
}

//ContestRoundT ...
var ContestRoundT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ContestRound",
	Description: "ContestRound",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "id",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contestround, ok := p.Source.(ContestRound); ok {
					return contestround.ID, nil
				}
				return nil, nil
			},
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contestround, ok := p.Source.(ContestRound); ok {
					return contestround.CreatedAt, nil
				}
				return nil, nil
			},
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contestround, ok := p.Source.(ContestRound); ok {
					return contestround.UpdatedAt, nil
				}
				return nil, nil
			},
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "deletion date",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contestround, ok := p.Source.(ContestRound); ok {
					return contestround.DeletedAt, nil
				}
				return nil, nil
			},
		},
		"Order": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "ContestRound order",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contestround, ok := p.Source.(ContestRound); ok {
					return contestround.Order, nil
				}
				return nil, nil
			},
		},
	},
})
