package models

import (
	"time"

	"github.com/google/uuid"
)

//WarningTemplate ...
type WarningTemplate struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (WarningTemplate) TableName() string {
	return "warningtemplates"
}
