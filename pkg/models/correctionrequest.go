package models

import (
	"time"

	"github.com/google/uuid"
)

//CorrectionRequest ...
type CorrectionRequest struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Requester uuid.UUID `sql:"type:uuid"`
}

//TableName ...
func (CorrectionRequest) TableName() string {
	return "correctionrequests"
}
