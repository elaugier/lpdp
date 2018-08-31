package models

import (
	"time"

	"github.com/google/uuid"
)

//CorrectionRequestAction ...
type CorrectionRequestAction struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	RequestRef uuid.UUID `sql:"type:uuid"`
}

//TableName ...
func (CorrectionRequestAction) TableName() string {
	return "correctionrequestactions"
}
