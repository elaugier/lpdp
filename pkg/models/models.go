package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//User ...
type User struct {
	gorm.Model
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	RemovedAt time.Time
	Email     string
	UserName  string
	Password  string
	FirstName string
	LastName  string
}

//Warning ...
type Warning struct {
	gorm.Model
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	RemovedAt time.Time
}
