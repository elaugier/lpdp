package models

import (
	"time"

	"github.com/google/uuid"
)

//BadIPAddress ...
type BadIPAddress struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	IPAddress string    `json:"ip_address"`
	Reason    string    `json:"reason"`
	ExpiresAt time.Time `json:"expires_at"`
}

//TableName ...
func (BadIPAddress) TableName() string {
	return "badipaddresses"
}
