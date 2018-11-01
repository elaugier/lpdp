package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//User ...
type User struct {
	ID                 uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	DeletedAt          time.Time `json:"deleted_at"`
	Email              string    `gorm:"type:varchar(200);unique_index"`
	UserName           string    `gorm:"type:varchar(100);unique_index"`
	Password           string    `gorm:"type:varchar(100)"`
	FirstName          string    `gorm:"type:varchar(100)"`
	LastName           string    `gorm:"type:varchar(100)"`
	SurName            string    `gorm:"type:varchar(100)"`
	Pseudo             string    `gorm:"type:varchar(100);unique_index"`
	BirthDate          time.Time
	Gender             uint
	PostalAddress      string              `gorm:"type:varchar(100)"`
	PostalAddress2     string              `gorm:"type:varchar(100)"`
	Job                string              `gorm:"type:varchar(100)"`
	Hobbies            string              `gorm:"type:varchar(100)"`
	AccountStatus      string              `gorm:"type:varchar(100)"`
	AccountType        string              `gorm:"type:varchar(100)"`
	Role               string              `gorm:"type:varchar(100)"`
	Achievements       []Achievement       `gorm:"foreignkey:OwnerRef"`
	Posts              []Post              `gorm:"foreignkey:AuthorRef"`
	Comments           []Comment           `gorm:"foreignkey:AuthorRef"`
	Contacts           []Contact           `gorm:"foreignkey:OwnerRef"`
	InBoxMessages      []Message           `gorm:"foreignkey:SenderRef"`
	OutBoxMessages     []Message           `gorm:"foreignkey:RecipientRef"`
	CorrectionRequests []CorrectionRequest `gorm:"foreignkey:RequesterRef"`
	Payments           []Payment           `gorm:"foreignkey:PayerRef"`
}

//TableName ...
func (User) TableName() string {
	return "users"
}

//UserT ...
var UserT = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "Member of LPDP",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "user id",
		},
		"CreatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "creation date of user",
		},
		"UpdatedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "update date of user",
		},
		"DeletedAt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "deletion date of user",
		},
		"Email": &graphql.Field{
			Type:        graphql.String,
			Description: "email address of user",
		},
	},
})

//UserRole ...
type UserRole int

const (
	//Guest ...
	Guest UserRole = 0
	//Member ...
	Member UserRole = 1
	//Moderator ...
	Moderator UserRole = 2
	//Administrator ...
	Administrator UserRole = 3
	//SuperUser ...
	SuperUser UserRole = 4
)

//MemberType ...
type MemberType int

const (
	//Regular ...
	Regular MemberType = 0
	//Privilege ...
	Privilege MemberType = 1
)
