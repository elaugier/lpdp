package models

import (
	"time"

	"github.com/google/uuid"
)

//CorrectionRequestAction ...
type CorrectionRequestAction struct {
	ID         uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt  time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	RequestRef uuid.UUID  `gorm:"not null" sql:"type:uuid"`
}

//TableName ...
func (CorrectionRequestAction) TableName() string {
	return "correctionrequestactions"
}
