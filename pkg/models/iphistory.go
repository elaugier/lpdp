package models

import (
	"time"

	"github.com/google/uuid"
)

//IPHistory ...
type IPHistory struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	UserRef   uuid.UUID  `sql:"type:uuid" json:"user_ref"`
	IPAddress string     `gorm:"not null;type:varchar(46)" json:"ip_address"`
	ExpiresAt *time.Time `json:"expires_at"`
}

//TableName ...
func (IPHistory) TableName() string {
	return "iphistories"
}
