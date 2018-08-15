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

//Alert ...
type Alert struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	RemovedAt time.Time
}

//TableName ...
func (Alert) TableName() string {
	return "alerts"
}

//AlertAction ...
type AlertAction struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	RemovedAt time.Time
}

//TableName ...
func (AlertAction) TableName() string {
	return "alertactions"
}

//Badge ...
type Badge struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	RemovedAt time.Time
}

//TableName ...
func (Badge) TableName() string {
	return "badges"
}

//BadgeType ...
type BadgeType struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	RemovedAt time.Time
}

//TableName ...
func (BadgeType) TableName() string {
	return "badgetypes"
}

//BadIPAddress ...
type BadIPAddress struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	RemovedAt time.Time
}

//TableName ...
func (BadIPAddress) TableName() string {
	return "badipaddresses"
}

//BookPart ...
type BookPart struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	RemovedAt time.Time
}

//TableName ...
func (BookPart) TableName() string {
	return "bookparts"
}

//Book ...
type Book struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	RemovedAt time.Time
}

//TableName ...
func (Book) TableName() string {
	return "books"
}

//CoAuthor ...
type CoAuthor struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	RemovedAt time.Time
}

//TableName ...
func (CoAuthor) TableName() string {
	return "coauthors"
}

//Comment ...
type Comment struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	RemovedAt time.Time
}

//TableName ...
func (Comment) TableName() string {
	return "comments"
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
