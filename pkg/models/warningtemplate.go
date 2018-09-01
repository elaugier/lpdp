package models

import (
	"time"

	"github.com/google/uuid"
)

//WarningTemplate ...
type WarningTemplate struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Name      string    `gorm:"type:varchar(200);unique_index" json:"name"`
	Content   string    `sql:"type:text" json:"content"`
}

//TableName ...
func (WarningTemplate) TableName() string {
	return "warningtemplates"
}
