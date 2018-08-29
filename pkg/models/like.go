package models

import (
	"time"

	"github.com/google/uuid"
)

//Like ...
type Like struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (Like) TableName() string {
	return "likes"
}
