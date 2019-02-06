package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//Alert ...
type Alert struct {
	ID           uuid.UUID     `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt    time.Time     `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time     `gorm:"not null" json:"updated_at"`
	DeletedAt    *time.Time    `json:"deleted_at"`
	Type         string        `gorm:"type:varchar(100)" json:"type"`
	Details      string        `gorm:"type:varchar(2000)" json:"details"`
	PostRef      string        `sql:"type:uuid" json:"post_ref"`
	CommentRef   string        `sql:"type:uuid" json:"comment_ref"`
	UserRef      string        `sql:"type:uuid" json:"user_ref"`
	AlertActions []AlertAction `gorm:"foreignkey:AlertRef"`
}

//TableName ...
func (Alert) TableName() string {
	return "alerts"
}

//AlertT ...
var AlertT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Alert",
	Description: "Alert for LPDP Post or Comment",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "alert id",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if alert, ok := p.Source.(Alert); ok {
					return alert.ID, nil
				}
				return nil, nil
			},
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date of alert",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if alert, ok := p.Source.(Alert); ok {
					return alert.CreatedAt, nil
				}
				return nil, nil
			},
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date of alert",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if alert, ok := p.Source.(Alert); ok {
					return alert.UpdatedAt, nil
				}
				return nil, nil
			},
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "deletion date of alert",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if alert, ok := p.Source.(Alert); ok {
					return alert.DeletedAt, nil
				}
				return nil, nil
			},
		},
		"Type": &graphql.Field{
			Type:        graphql.String,
			Description: "type of alert",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if alert, ok := p.Source.(Alert); ok {
					return alert.Type, nil
				}
				return nil, nil
			},
		},
		"Details": &graphql.Field{
			Type:        graphql.String,
			Description: "Details of alert",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if alert, ok := p.Source.(Alert); ok {
					return alert.Details, nil
				}
				return nil, nil
			},
		},
		"PostRef": &graphql.Field{
			Type:        graphql.String,
			Description: "post targeted by the alert",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if alert, ok := p.Source.(Alert); ok {
					return alert.PostRef, nil
				}
				return nil, nil
			},
		},
		"CommentRef": &graphql.Field{
			Type:        graphql.String,
			Description: "comment targeted of alert",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if alert, ok := p.Source.(Alert); ok {
					return alert.CommentRef, nil
				}
				return nil, nil
			},
		},
		"UserRef": &graphql.Field{
			Type:        graphql.String,
			Description: "owner of alert",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if alert, ok := p.Source.(Alert); ok {
					return alert.UserRef, nil
				}
				return nil, nil
			},
		},
	},
})
