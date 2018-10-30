package models

import (
	"time"

	"github.com/google/uuid"
)

//BookPart ...
type BookPart struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Title     string    `json:"title"`
	BookRef   uuid.UUID `sql:"type:uuid"`
}

//TableName ...
func (BookPart) TableName() string {
	return "bookparts"
}
