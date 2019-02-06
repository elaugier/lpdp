package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

//Activity ...
type Activity struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Message   string     `gorm:"type:varchar(255);not null" json:"message"`
	Type      string     `gorm:"type:varchar(100);not null" json:"type"`
	OwnerRef  uuid.UUID  `sql:"type:uuid;not null" json:"owner_ref"`
}

//TableName ...
func (Activity) TableName() string {
	return "activities"
}

//ActivityT ...
var ActivityT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Activity",
	Description: "Activity for LPDP Member",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "activity id",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if activity, ok := p.Source.(Activity); ok {
					return activity.ID, nil
				}
				return nil, nil
			},
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date of activity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if activity, ok := p.Source.(Activity); ok {
					return activity.CreatedAt, nil
				}
				return nil, nil
			},
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date of activity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if activity, ok := p.Source.(Activity); ok {
					return activity.UpdatedAt, nil
				}
				return nil, nil
			},
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "deletion date of activity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if activity, ok := p.Source.(Activity); ok {
					return activity.DeletedAt, nil
				}
				return nil, nil
			},
		},
		"Message": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "message of activity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if activity, ok := p.Source.(Activity); ok {
					return activity.Message, nil
				}
				return nil, nil
			},
		},
		"Type": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "type of activity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if activity, ok := p.Source.(Activity); ok {
					return activity.Type, nil
				}
				return nil, nil
			},
		},
		"OwnerRef": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "owner of activity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if activity, ok := p.Source.(Activity); ok {
					return activity.OwnerRef, nil
				}
				return nil, nil
			},
		},
	},
})
