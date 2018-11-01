package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//BadIPAddress ...
type BadIPAddress struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	IPAddress string    `json:"ip_address"`
	Reason    string    `json:"reason"`
	ExpiresAt time.Time `json:"expires_at"`
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
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date of bad ip address",
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date of bad ip address",
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "deletion date of bad ip address",
		},
		"IPAddress": &graphql.Field{
			Type:        graphql.String,
			Description: "ip address of bad ip address",
		},
		"Reason": &graphql.Field{
			Type:        graphql.String,
			Description: "reason of bad ip address",
		},
		"Expires": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "date expiration of bad ip address",
		},
	},
})
