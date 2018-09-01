package models

import (
	"time"

	"github.com/google/uuid"
)

//Tag ...
type Tag struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Code      string    `gorm:"not null;type:varchar(3);unique"`
	Label     string    `gorm:"not null;type:varchar(100);unique"`
	Enabled   bool      `gorm:"default:false"`
	StartedAt time.Time `gorm:"column:start_at" json:"start_at"`
	ClosedAt  time.Time `gorm:"column:closed_at" json:"closed_at"`
	Type      string    `gorm:"not null;unique"`
	Mature    bool      `gorm:"default:false"`
}

//TableName ...
func (Tag) TableName() string {
	return "tags"
}
