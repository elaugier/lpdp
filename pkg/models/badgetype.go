package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//BadgeType ...
type BadgeType struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Title     string     `gorm:"type:varchar(200);unique_index"`
	Badges    []Badge    `gorm:"foreignkey:TypeRef"`
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
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badgetype, ok := p.Source.(BadgeType); ok {
					return badgetype.ID, nil
				}
				return nil, nil
			},
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date of badge type",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badgetype, ok := p.Source.(BadgeType); ok {
					return badgetype.CreatedAt, nil
				}
				return nil, nil
			},
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date of badge type",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badgetype, ok := p.Source.(BadgeType); ok {
					return badgetype.UpdatedAt, nil
				}
				return nil, nil
			},
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "deletion date of badge type",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badgetype, ok := p.Source.(BadgeType); ok {
					return badgetype.DeletedAt, nil
				}
				return nil, nil
			},
		},
		"Title": &graphql.Field{
			Type:        graphql.String,
			Description: "title of badge type",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badgetype, ok := p.Source.(BadgeType); ok {
					return badgetype.Title, nil
				}
				return nil, nil
			},
		},
	},
})
