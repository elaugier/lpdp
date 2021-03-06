package models

import (
	"time"

	"github.com/google/uuid"
)

//Vote ...
type Vote struct {
	ID        uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	OwnerRef  uuid.UUID  `gorm:"not null" json:"owner_ref"`
}

//TableName ...
func (Vote) TableName() string {
	return "votes"
}
