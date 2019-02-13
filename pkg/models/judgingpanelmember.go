package models

import (
	"time"

	"github.com/google/uuid"
)

//JudgingPanelMember ...
type JudgingPanelMember struct {
	ID              uuid.UUID  `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt       time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	JudgingPanelRef uuid.UUID  `sql:"type:uuid" json:"judging_panel_ref"`
	UserRef         uuid.UUID  `sql:"type:uuid" json:"user_ref"`
	Title           string     `json:"title"`
}

//TableName ...
func (JudgingPanelMember) TableName() string {
	return "judgingpanelsmembers"
}
