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
	Type      string    `gorm:"type:varchar(100)"`
	Details   string    `gorm:"type:varchar(2000)"`
	User      string    `sql:"type:uuid"`
}

//TableName ...
func (Alert) TableName() string {
	return "alerts"
}
