package models

import (
	"time"

	"github.com/google/uuid"
)

//IPHistory ...
type IPHistory struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	User      uuid.UUID `sql:"type:uuid"`
	IP        string
}

//TableName ...
func (IPHistory) TableName() string {
	return "iphistories"
}
