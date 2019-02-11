package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

//IPHistory ...
type IPHistory struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	UserRef   uuid.UUID  `sql:"type:uuid" json:"user_ref"`
	IPAddress string     `gorm:"not null;type:varchar(46)" json:"ip_address"`
	ExpiresAt *time.Time `json:"expires_at"`
}

//TableName ...
func (IPHistory) TableName() string {
	return "iphistories"
}

//IPHistoryT ...
var IPHistoryT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "IPHistory",
	Description: "IP history og user",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Identifier",
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "Creation Date",
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "Update Date",
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "Deletion Date",
		},
	},
})
