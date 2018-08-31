package models

import (
	"time"

	"github.com/google/uuid"
)

//CorrectionRequest ...
type CorrectionRequest struct {
	ID           uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	RequesterRef uuid.UUID `sql:"type:uuid"`
}

//TableName ...
func (CorrectionRequest) TableName() string {
	return "correctionrequests"
}
