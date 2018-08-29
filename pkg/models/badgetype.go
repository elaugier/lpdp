package models

import (
	"time"

	"github.com/google/uuid"
)

//BadgeType ...
type BadgeType struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (BadgeType) TableName() string {
	return "badgetypes"
}
