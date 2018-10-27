package models

import (
	"time"

	"github.com/google/uuid"
)

//Alert ...
type Alert struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Type      string    `gorm:"type:varchar(100)" json:"type"`
	Details   string    `gorm:"type:varchar(2000)" json:"details"`
	UserRef   string    `sql:"type:uuid" json:"user_ref"`
}

//TableName ...
func (Alert) TableName() string {
	return "alerts"
}
