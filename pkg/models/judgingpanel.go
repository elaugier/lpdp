package models

import (
	"time"

	"github.com/google/uuid"
)

//JudgingPanel ...
type JudgingPanel struct {
	ID                  uuid.UUID            `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt           time.Time            `gorm:"not null" json:"created_at"`
	UpdatedAt           time.Time            `gorm:"not null" json:"updated_at"`
	DeletedAt           *time.Time           `json:"deleted_at"`
	JudgingPanelMembers []JudgingPanelMember `gorm:"foreignkey:JudgingPanelRef"`
}

//TableName ...
func (JudgingPanel) TableName() string {
	return "judgingpanels"
}
