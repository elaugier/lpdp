package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//Tag ...
type Tag struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Code      string     `gorm:"not null;type:varchar(3);unique"`
	Label     string     `gorm:"not null;type:varchar(100);unique"`
	Enabled   bool       `gorm:"default:false"`
	StartedAt time.Time  `gorm:"column:start_at" json:"start_at"`
	ClosedAt  time.Time  `gorm:"column:closed_at" json:"closed_at"`
	Type      string     `gorm:"not null;unique"`
	Mature    bool       `gorm:"default:false"`
}

//TableName ...
func (Tag) TableName() string {
	return "tags"
}

//TagT ...
var TagT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Tag",
	Description: "Tag",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "tag id",
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date of tag",
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date of tag",
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "deletion date of tag",
		},
		"Code": &graphql.Field{
			Type:        graphql.String,
			Description: "code of tag",
		},
		"Label": &graphql.Field{
			Type:        graphql.String,
			Description: "label of tag",
		},
		"Enable": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "availability of tag",
		},
		"StartedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "start date of availability of tag",
		},
		"ClosedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "end date of availability of tag",
		},
		"Type": &graphql.Field{
			Type:        graphql.String,
			Description: "type of tag",
		},
		"Mature": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "tag maturity",
		},
	},
})
