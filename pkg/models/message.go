package models

import (
	"time"

	"github.com/google/uuid"
)

//Message ...
type Message struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Sender    User
	ReplyTo   uuid.UUID `sql:"type:uuid"`
	Recipient User
	Content   string `sql:"type:text"`
}

//TableName ...
func (Message) TableName() string {
	return "messages"
}