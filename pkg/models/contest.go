package models

import (
	"time"

	"github.com/google/uuid"
)

//Contest ...
type Contest struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (Contest) TableName() string {
	return "contests"
}
