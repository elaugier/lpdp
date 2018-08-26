package models

import (
	"time"

	"github.com/google/uuid"
)

//Request ...
type Request struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (Request) TableName() string {
	return "requests"
}
