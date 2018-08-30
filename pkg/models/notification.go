package models

import (
	"time"

	"github.com/google/uuid"
)

//Notification ...
type Notification struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (Notification) TableName() string {
	return "notifications"
}
