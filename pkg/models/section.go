package models

import (
	"time"

	"github.com/google/uuid"
)

//Section ...
type Section struct {
	ID          uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
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
