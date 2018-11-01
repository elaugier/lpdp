package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//Badge ...
type Badge struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Message   string    `gorm:"type:varchar(255)" json:"message"`
	TypeRef   uuid.UUID `sql:"type:uuid" json:"type_ref"`
	OwnerRef  uuid.UUID `sql:"type:uuid" json:"owner_ref"`
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
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date of badge",
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date of badge",
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "deletion date of badge",
		},
		"Message": &graphql.Field{
			Type:        graphql.String,
			Description: "message of badge",
		},
		"TypeRef": &graphql.Field{
			Type:        graphql.String,
			Description: "type of badge",
		},
		"OwnerRef": &graphql.Field{
			Type:        graphql.String,
			Description: "owner of badge",
		},
	},
})
