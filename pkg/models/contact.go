package models

import (
	"time"

	"github.com/google/uuid"
)

//Contact ...
type Contact struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	LastName       string `gorm:"type:varchar(200)"`
	FirstName      string `gorm:"not null;type:varchar(200)"`
	BirthDate      time.Time
	OwnerRef       uuid.UUID `sql:"type:uuid"`
	UserContactRef uuid.UUID `sql:"type:uuid"`
}

//TableName ...
func (Contact) TableName() string {
	return "contacts"
}
