package models

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/google/uuid"
)

//User ...
type User struct {
	ID                 uuid.UUID           `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt          time.Time           `gorm:"not null" json:"created_at"`
	UpdatedAt          time.Time           `gorm:"not null" json:"updated_at"`
	DeletedAt          *time.Time          `json:"deleted_at"`
	Email              string              `gorm:"type:varchar(500);unique_index" json:"email"`
	IsEmailVerified    bool                `gorm:"type:boolean;unique_index" json:"is_email_verified"`
	UserName           string              `gorm:"type:varchar(100);unique_index" json:"username"`
	Password           string              `gorm:"type:varchar(100)" json:"password"`
	FirstName          string              `gorm:"type:varchar(100)" json:"firstname"`
	LastName           string              `gorm:"type:varchar(100)" json:"lastname"`
	SurName            string              `gorm:"type:varchar(100)" json:"surname"`
	Pseudo             string              `gorm:"type:varchar(100);unique_index" json:"pseudo"`
	DateOfBirth        time.Time           `json:"date_of_birth"`
	Gender             uint                `json:"gender"`
	TimeZone           string              `gorm:"type:varchar(100)" json:"time_zone"`
	PostalAddress      string              `gorm:"type:varchar(100)" json:"postal_address"`
	PostalAddress2     string              `gorm:"type:varchar(100)" json:"postal_address_2"`
	Job                string              `gorm:"type:varchar(100)" json:"job"`
	Hobbies            string              `gorm:"type:varchar(100)" json:"hobbies"`
	AccountStatus      string              `gorm:"type:varchar(100)" json:"account_status"`
	AccountType        uint                `json:"account_type"`
	Role               string              `gorm:"type:varchar(100)" json:"role"`
	Achievements       []Achievement       `gorm:"foreignkey:OwnerRef"`
	Badges             []Badge             `gorm:"foreignkey:OwnerRef"`
	Books              []Book              `gorm:"foreignkey:AuthorRef"`
	Posts              []Post              `gorm:"foreignkey:AuthorRef"`
	Comments           []Comment           `gorm:"foreignkey:AuthorRef"`
	Contacts           []Contact           `gorm:"foreignkey:OwnerRef"`
	InBoxMessages      []Message           `gorm:"foreignkey:SenderRef"`
	OutBoxMessages     []Message           `gorm:"foreignkey:RecipientRef"`
	CorrectionRequests []CorrectionRequest `gorm:"foreignkey:RequesterRef"`
	Payments           []Payment           `gorm:"foreignkey:PayerRef"`
	Votes              []Vote              `gorm:"foreignkey:OwnerRef"`
	Warnings           []Warning           `gorm:"foreignkey:UserRef"`
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
			Type:        graphql.DateTime,
			Description: "deletion date of user",
		},
		"Email": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "email address of user",
		},
		"IsEmailVerified": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Email check",
		},
		"UserName": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "User Name",
		},
		"Password": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "User Password",
		},
		"FirstName": &graphql.Field{
			Type:        graphql.String,
			Description: "User Firstname",
		},
		"LastName": &graphql.Field{
			Type:        graphql.String,
			Description: "User Lastname",
		},
		"SurName": &graphql.Field{
			Type:        graphql.String,
			Description: "User surname",
		},
		"Pseudo": &graphql.Field{
			Type:        graphql.String,
			Description: "User Pseudo",
		},
		"DateOfBirth": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.DateTime),
			Description: "birthdate of user",
		},
		"Gender": &graphql.Field{
			Type:        graphql.Int,
			Description: "User gender",
		},
		"TimeZone": &graphql.Field{
			Type:        graphql.String,
			Description: "User time zone",
		},
		"PostalAddress": &graphql.Field{
			Type:        graphql.String,
			Description: "User Postal address",
		},
		"PostalAddress2": &graphql.Field{
			Type:        graphql.String,
			Description: "User Postal address bis",
		},
		"Job": &graphql.Field{
			Type:        graphql.String,
			Description: "User Job",
		},
		"Hobbies": &graphql.Field{
			Type:        graphql.String,
			Description: "User hobbies",
		},
		"AccountStatus": &graphql.Field{
			Type:        graphql.String,
			Description: "Account status",
		},
		"AccountType": &graphql.Field{
			Type:        graphql.Int,
			Description: "Account type",
		},
		"Role": &graphql.Field{
			Type:        graphql.String,
			Description: "User Role\nAuthorized Values are:\n\t\"0\": Guest\n\t\"1\": Member\n\t\"2\": Moderator\n\t\"3\": Administrator\n\t\"4\": SuperUser",
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
