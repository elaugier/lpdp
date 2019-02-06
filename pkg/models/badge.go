package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//Badge ...
type Badge struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Message   string     `gorm:"type:varchar(255)" json:"message"`
	TypeRef   uuid.UUID  `sql:"type:uuid" json:"type_ref"`
	OwnerRef  uuid.UUID  `sql:"type:uuid" json:"owner_ref"`
}

//TableName ...
func (Badge) TableName() string {
	return "badges"
}

//BadgeT ...
var BadgeT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Badge",
	Description: "Badge for LPDP Member",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "badge id",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badge, ok := p.Source.(Badge); ok {
					return badge.ID, nil
				}
				return nil, nil
			},
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date of badge",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badge, ok := p.Source.(Badge); ok {
					return badge.CreatedAt, nil
				}
				return nil, nil
			},
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date of badge",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badge, ok := p.Source.(Badge); ok {
					return badge.UpdatedAt, nil
				}
				return nil, nil
			},
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "deletion date of badge",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badge, ok := p.Source.(Badge); ok {
					return badge.DeletedAt, nil
				}
				return nil, nil
			},
		},
		"Message": &graphql.Field{
			Type:        graphql.String,
			Description: "message of badge",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badge, ok := p.Source.(Badge); ok {
					return badge.Message, nil
				}
				return nil, nil
			},
		},
		"TypeRef": &graphql.Field{
			Type:        graphql.String,
			Description: "type of badge",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badge, ok := p.Source.(Badge); ok {
					return badge.TypeRef, nil
				}
				return nil, nil
			},
		},
		"OwnerRef": &graphql.Field{
			Type:        graphql.String,
			Description: "owner of badge",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badge, ok := p.Source.(Badge); ok {
					return badge.OwnerRef, nil
				}
				return nil, nil
			},
		},
	},
})
