package models

import (
	"time"

	"github.com/google/uuid"
)

//Badge ...
type Badge struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Message   string    `gorm:"type:varchar(255)" json:"message"`
	Type      string    `gorm:"type:varchar(100)" json:"type"`
	OwnerRef  uuid.UUID `sql:"type:uuid" json:"owner_ref"`
}

//TableName ...
func (Badge) TableName() string {
	return "badges"
}
