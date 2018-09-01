package models

import (
	"time"

	"github.com/google/uuid"
)

//Payment ...
type Payment struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	PayerRef  uuid.UUID `sql:"type:uuid" json:"payer_ref"`
}

//TableName ...
func (Payment) TableName() string {
	return "payments"
}
