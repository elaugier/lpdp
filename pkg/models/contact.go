package models

import (
	"time"

	"github.com/google/uuid"
)

//Contact ...
type Contact struct {
	ID             uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
	LastName       string    `gorm:"type:varchar(200)" json:"last_name" binding:"required"`
	FirstName      string    `gorm:"not null;type:varchar(200)" json:"first_name" binding:"required"`
	BirthDate      time.Time `json:"birth_date"`
	OwnerRef       uuid.UUID `sql:"type:uuid" json:"owner_ref" binding:"required"`
	UserContactRef uuid.UUID `sql:"type:uuid" json:"user_contact_ref" binding:"required"`
}

//TableName ...
func (Contact) TableName() string {
	return "contacts"
}
