package models

import (
	"time"

	"github.com/google/uuid"
)

//Activity ...
type Activity struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Message   string
	Type      string
}

//TableName ...
func (Activity) TableName() string {
	return "activities"
}
