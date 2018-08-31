package models

import (
	"time"

	"github.com/google/uuid"
)

//CorrectionRequest ...
type CorrectionRequest struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	RequesterRef uuid.UUID `sql:"type:uuid"`
}

//TableName ...
func (CorrectionRequest) TableName() string {
	return "correctionrequests"
}
