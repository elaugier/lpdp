package models

import (
	"time"

	"github.com/google/uuid"
)

//Section ...
type Section struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Name        string `gorm:"type:varchar(200)"`
	ShortName   string `gorm:"type:varchar(30)"`
	Description string `sql:"type:text"`
	Order       uint
	SecShow     uint
	SecAdd      uint
	SecModify   uint
	SecRemove   uint
}

//TableName ...
func (Section) TableName() string {
	return "sections"
}
