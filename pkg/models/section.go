package models

import (
	"time"

	"github.com/google/uuid"
)

//Section ...
type Section struct {
	ID          uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	Name        string    `gorm:"type:varchar(200);unique_index" json:"name"`
	ShortName   string    `gorm:"type:varchar(30);unique_index" json:"shortname"`
	Description string    `sql:"type:text" json:"description"`
	Order       uint      `json:"order"`
	SecShow     uint      `json:"secshow"`
	SecAdd      uint      `json:"secadd"`
	SecModify   uint      `json:"secmodify"`
	SecRemove   uint      `json:"secremove"`
}

//TableName ...
func (Section) TableName() string {
	return "sections"
}
