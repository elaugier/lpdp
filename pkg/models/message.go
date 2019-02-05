package models

import (
	"time"

	"github.com/google/uuid"
)

//Message ...
type Message struct {
	ID           uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt    time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	SenderRef    User
	ReplyTo      uuid.UUID `sql:"type:uuid"`
	RecipientRef User
	Content      string `sql:"type:text"`
}

//TableName ...
func (Message) TableName() string {
	return "messages"
}
