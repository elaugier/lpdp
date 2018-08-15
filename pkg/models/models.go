package models

import (
	"time"
)

//Achievement ...
type Achievement struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	RemovedAt time.Time
}

func (Achievement) tableName() string {
	return "achievements"
}

//Activity ...
type Activity struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	RemovedAt time.Time
}

//TableName ...
func (Activity) TableName() string {
	return "activities"
}

//User ...
type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	RemovedAt time.Time
	Email     string `gorm:"type:varchar(100);unique_index"`
	UserName  string
	Password  string
	FirstName string
	LastName  string
	BirthDate time.Time
	Gender    uint
}

//TableName ...
func (User) TableName() string {
	return "users"
}

//Warning ...
type Warning struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	RemovedAt time.Time
}

//TableName ...
func (Warning) TableName() string {
	return "warnings"
}
