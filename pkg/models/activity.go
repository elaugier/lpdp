package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

//Activity ...
type Activity struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Message   string    `gorm:"type:varchar(255);not null" json:"message"`
	Type      string    `gorm:"type:varchar(100);not null" json:"type"`
	OwnerRef  uuid.UUID `sql:"type:uuid;not null" json:"owner_ref"`
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
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date of activity",
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date of activity",
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "deletion date of activity",
		},
		"Message": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "message of activity",
		},
		"Type": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "type of activity",
		},
		"OwnerRef": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "owner of activity",
		},
	},
})
