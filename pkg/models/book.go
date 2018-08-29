package models

import (
	"time"

	"github.com/google/uuid"
)

//Book ...
type Book struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Title     string    `gorm:"type:varchar(100)"`
	Summary   string    `sql:"type:text;"`
	Author    uuid.UUID `sql:"type:uuid"`
}

//TableName ...
func (Book) TableName() string {
	return "books"
}
