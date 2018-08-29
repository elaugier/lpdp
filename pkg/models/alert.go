package models

import (
	"time"

	"github.com/google/uuid"
)

//Alert ...
type Alert struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Type      string `gorm:"type:varchar(100)"`
	Details   string `gorm:"type:varchar(2000)"`
	User      string `gorm:"type:uuid()"`
}

//TableName ...
func (Alert) TableName() string {
	return "alerts"
}
