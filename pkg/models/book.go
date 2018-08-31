package models

import (
	"time"

	"github.com/google/uuid"
)

//Book ...
type Book struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Title     string     `gorm:"type:varchar(100)"`
	Summary   string     `sql:"type:text;"`
	Author    uuid.UUID  `sql:"type:uuid"`
	BookParts []BookPart `gorm:"foreignkey:BookRef"`
}

//TableName ...
func (Book) TableName() string {
	return "books"
}
