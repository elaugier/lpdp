package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

//Contact ...
type Contact struct {
	ID             uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt      time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
	LastName       string     `gorm:"type:varchar(200)" json:"last_name" binding:"required"`
	FirstName      string     `gorm:"not null;type:varchar(200)" json:"first_name" binding:"required"`
	BirthDate      time.Time  `json:"birth_date"`
	OwnerRef       uuid.UUID  `sql:"type:uuid" json:"owner_ref" binding:"required"`
	UserContactRef uuid.UUID  `sql:"type:uuid" json:"user_contact_ref" binding:"required"`
}

//TableName ...
func (Contact) TableName() string {
	return "contacts"
}

//ContactT ...
var ContactT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Contact",
	Description: "Contact",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "id",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contact, ok := p.Source.(Contact); ok {
					return contact.ID, nil
				}
				return nil, nil
			},
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contact, ok := p.Source.(Contact); ok {
					return contact.CreatedAt, nil
				}
				return nil, nil
			},
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contact, ok := p.Source.(Contact); ok {
					return contact.UpdatedAt, nil
				}
				return nil, nil
			},
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "deletion date",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contact, ok := p.Source.(Contact); ok {
					return contact.DeletedAt, nil
				}
				return nil, nil
			},
		},
	},
})
