package models

import (
	"time"

	"github.com/google/uuid"
)

//JudgingPanel ...
type JudgingPanel struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

//TableName ...
func (JudgingPanel) TableName() string {
	return "judgingpanels"
}
