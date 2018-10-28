package models

import (
	"time"

	"github.com/google/uuid"
)

//AlertAction ...
type AlertAction struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Title     string    `gorm:"type:varchar(100)" json:"title"`
	Details   string    `gorm:"type:varchar(2000)" json:"details"`
	AlertRef  string    `sql:"type:uuid" json:"alert_ref"`
	UserRef   string    `sql:"type:uuid" json:"user_ref"`
}

//TableName ...
func (AlertAction) TableName() string {
	return "alertactions"
}
