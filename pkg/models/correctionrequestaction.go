package models

import (
	"time"

	"github.com/google/uuid"
)

//CorrectionRequestAction ...
type CorrectionRequestAction struct {
	ID         uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	RequestRef uuid.UUID `sql:"type:uuid"`
}

//TableName ...
func (CorrectionRequestAction) TableName() string {
	return "correctionrequestactions"
}