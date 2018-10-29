package models

import (
	"time"

	"github.com/google/uuid"
)

//BadgeType ...
type BadgeType struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Title     string    `gorm:"type:varchar(200);unique_index"`
	Badges    []Badge   `gorm:"foreignkey:TypeRef"`
}

//TableName ...
func (BadgeType) TableName() string {
	return "badgetypes"
}
