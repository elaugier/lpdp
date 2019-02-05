package models

import (
	"time"

	"github.com/google/uuid"
)

//Contest ...
type Contest struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//TableName ...
func (Contest) TableName() string {
	return "contests"
}
