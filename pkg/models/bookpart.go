package models

import (
	"time"

	"github.com/google/uuid"
)

//BookPart ...
type BookPart struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	BookRef   uuid.UUID `sql:"type:uuid"`
}

//TableName ...
func (BookPart) TableName() string {
	return "bookparts"
}
