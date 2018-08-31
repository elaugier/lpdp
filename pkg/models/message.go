package models

import (
	"time"

	"github.com/google/uuid"
)

//Message ...
type Message struct {
	ID           uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	SenderRef    User
	ReplyTo      uuid.UUID `sql:"type:uuid"`
	RecipientRef User
	Content      string `sql:"type:text"`
}

//TableName ...
func (Message) TableName() string {
	return "messages"
}
