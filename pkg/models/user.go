package models

import (
	"time"

	"github.com/google/uuid"
)

//User ...
type User struct {
	ID             uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	Email          string `gorm:"type:varchar(200);unique_index"`
	UserName       string `gorm:"type:varchar(100);unique_index"`
	Password       string `gorm:"type:varchar(100)"`
	FirstName      string `gorm:"type:varchar(100)"`
	LastName       string `gorm:"type:varchar(100)"`
	SurName        string `gorm:"type:varchar(100)"`
	Pseudo         string `gorm:"type:varchar(100);unique_index"`
	BirthDate      time.Time
	Gender         uint
	PostalAddress  string `gorm:"type:varchar(100)"`
	PostalAddress2 string `gorm:"type:varchar(100)"`
	Job            string `gorm:"type:varchar(100)"`
	Hobbies        string `gorm:"type:varchar(100)"`
	AccountStatus  string `gorm:"type:varchar(100)"`
	AccountType    string `gorm:"type:varchar(100)"`
	Role           string `gorm:"type:varchar(100)"`
}

//TableName ...
func (User) TableName() string {
	return "users"
}
