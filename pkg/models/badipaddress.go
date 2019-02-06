package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//BadIPAddress ...
type BadIPAddress struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	IPAddress string     `gorm:"type:varchar(200)" json:"ip_address"`
	Reason    string     `gorm:"type:varchar(200)" json:"reason"`
	ExpiresAt time.Time  `json:"expires_at"`
}

//TableName ...
func (BadIPAddress) TableName() string {
	return "badipaddresses"
}

//BadIPAddressT ...
var BadIPAddressT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "BadIPAddress",
	Description: "Bad IP Address",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "bad ip address id",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badipaddress, ok := p.Source.(BadIPAddress); ok {
					return badipaddress.ID, nil
				}
				return nil, nil
			},
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date of bad ip address",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badipaddress, ok := p.Source.(BadIPAddress); ok {
					return badipaddress.CreatedAt, nil
				}
				return nil, nil
			},
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date of bad ip address",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badipaddress, ok := p.Source.(BadIPAddress); ok {
					return badipaddress.UpdatedAt, nil
				}
				return nil, nil
			},
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "deletion date of bad ip address",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badipaddress, ok := p.Source.(BadIPAddress); ok {
					return badipaddress.DeletedAt, nil
				}
				return nil, nil
			},
		},
		"IPAddress": &graphql.Field{
			Type:        graphql.String,
			Description: "ip address of bad ip address",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badipaddress, ok := p.Source.(BadIPAddress); ok {
					return badipaddress.IPAddress, nil
				}
				return nil, nil
			},
		},
		"Reason": &graphql.Field{
			Type:        graphql.String,
			Description: "reason of bad ip address",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badipaddress, ok := p.Source.(BadIPAddress); ok {
					return badipaddress.Reason, nil
				}
				return nil, nil
			},
		},
		"ExpiresAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "date expiration of bad ip address",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if badipaddress, ok := p.Source.(BadIPAddress); ok {
					return badipaddress.ExpiresAt, nil
				}
				return nil, nil
			},
		},
	},
})
