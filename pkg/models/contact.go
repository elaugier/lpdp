package models

import (
	"time"

	"github.com/google/uuid"
)

//Contact ...
type Contact struct {
	ID             uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	LastName       string
	FirstName      string
	BirthDate      time.Time
	OwnerRef       uuid.UUID `sql:"type:uuid"`
	UserContactRef uuid.UUID `sql:"type:uuid"`
}

//TableName ...
func (Contact) TableName() string {
	return "contacts"
}
