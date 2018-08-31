package models

import (
	"time"

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
}

//TableName ...
func (User) TableName() string {
	return "users"
}
