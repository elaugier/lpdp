package models

import (
	"time"

	"github.com/google/uuid"
)

//Tag ...
type Tag struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Code      string `gorm:"not null;varchar(3);unique"`
	Label     string `gorm:"not null;varchar(100);unique"`
	Enabled   bool   `gorm:"default:false"`
	StartedAt time.Time
	ClosedAt  time.Time
	Type      string `gorm:"not null;unique"`
	Mature    bool   `gorm:"default:false"`
}

//TableName ...
func (Tag) TableName() string {
	return "tags"
}
