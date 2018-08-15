package models

import (
	"time"

	"github.com/google/uuid"
)

//Achievement ...
type Achievement struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (Achievement) TableName() string {
	return "achievements"
}

//Activity ...
type Activity struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Message   string
	Type      string
}

//TableName ...
func (Activity) TableName() string {
	return "activities"
}

//Alert ...
type Alert struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Type      string `gorm:"type:varchar(100)"`
	Details   string `gorm:"type:varchar(2000)"`
	User      string `gorm:"type:uuid()"`
}

//TableName ...
func (Alert) TableName() string {
	return "alerts"
}

//AlertAction ...
type AlertAction struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (AlertAction) TableName() string {
	return "alertactions"
}

//Badge ...
type Badge struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (Badge) TableName() string {
	return "badges"
}

//BadgeType ...
type BadgeType struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (BadgeType) TableName() string {
	return "badgetypes"
}

//BadIPAddress ...
type BadIPAddress struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (BadIPAddress) TableName() string {
	return "badipaddresses"
}

//BookPart ...
type BookPart struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (BookPart) TableName() string {
	return "bookparts"
}

//Book ...
type Book struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Title     string    `gorm:"type:varchar(100)"`
	Summary   string    `sql:"type:text;"`
	Author    uuid.UUID `sql:"type:uuid"`
}

//TableName ...
func (Book) TableName() string {
	return "books"
}

//CoAuthor ...
type CoAuthor struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (CoAuthor) TableName() string {
	return "coauthors"
}

//Comment ...
type Comment struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (Comment) TableName() string {
	return "comments"
}

//Contact ...
type Contact struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (Contact) TableName() string {
	return "contacts"
}

//ContestRound ...
type ContestRound struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (ContestRound) TableName() string {
	return "contestrounds"
}

//Contest ...
type Contest struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (Contest) TableName() string {
	return "contests"
}

//CorrectionRequest ...
type CorrectionRequest struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Requester uuid.UUID `sql:"type:uuid"`
}

//TableName ...
func (CorrectionRequest) TableName() string {
	return "correctionrequests"
}

//CorrectionRequestAction ...
type CorrectionRequestAction struct {
	ID         uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	RequestRef uuid.UUID `sql:"type:uuid"`
}

//TableName ...
func (CorrectionRequestAction) TableName() string {
	return "correctionrequestactions"
}

//IPHistory ...
type IPHistory struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	User      uuid.UUID `sql:"type:uuid"`
	IP        string
}

//TableName ...
func (IPHistory) TableName() string {
	return "iphistories"
}

//Message ...
type Message struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Sender    User
	ReplyTo   uuid.UUID `sql:"type:uuid"`
	Recipient User
	Content   string `sql:"type:text"`
}

//TableName ...
func (Message) TableName() string {
	return "messages"
}

//Post ...
type Post struct {
	ID                  uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           time.Time
	PublishedAt         time.Time
	ApprovedAt          time.Time
	ApprouvedBy         uuid.UUID `sql:"type:uuid"`
	ApprouvedByUserName string
	Title               string
	Summary             string    `sql:"type:text"`
	Content             string    `sql:"type:text"`
	Author              uuid.UUID `sql:"type:uuid"`
}

//TableName ...
func (Post) TableName() string {
	return "posts"
}

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

//Warning ...
type Warning struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//TableName ...
func (Warning) TableName() string {
	return "warnings"
}
