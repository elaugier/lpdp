package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

//Section ...
type Section struct {
	ID          uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	Name        string    `gorm:"type:varchar(200);unique_index" json:"name"`
	ShortName   string    `gorm:"type:varchar(30);unique_index" json:"shortname"`
	Description string    `sql:"type:text" json:"description"`
	Order       uint      `json:"order"`
	SecShow     uint      `json:"secshow"`
	SecAdd      uint      `json:"secadd"`
	SecModify   uint      `json:"secmodify"`
	SecRemove   uint      `json:"secremove"`
}

//TableName ...
func (Section) TableName() string {
	return "sections"
}

//SectionT ...
var SectionT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Section",
	Description: "Section",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "id",
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date",
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date",
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "deletion date",
		},
		"Name": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Section name",
		},
		"ShortName": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Section name",
		},
		"Description": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Description of the section",
		},
		"Order": &graphql.Field{
			Type:        graphql.Int,
			Description: "Section order",
		},
		"SecShow": &graphql.Field{
			Type:        graphql.Int,
			Description: "access control entries for show content",
		},
		"SecAdd": &graphql.Field{
			Type:        graphql.Int,
			Description: "access control entries for add content",
		},
		"SecModify": &graphql.Field{
			Type:        graphql.Int,
			Description: "access control entries for modify content",
		},
		"SecRemove": &graphql.Field{
			Type:        graphql.Int,
			Description: "access control entries for remove content",
		},
	},
})
