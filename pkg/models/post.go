package models

import (
	"time"

	"github.com/google/uuid"
)

//Post ...
type Post struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	PublishedAt         time.Time
	ApprovedAt          time.Time
	ApprouvedBy         uuid.UUID `sql:"type:uuid"`
	ApprouvedByUserName string
	Title               string
	Summary             string    `sql:"type:text"`
	Content             string    `sql:"type:text"`
	AuthorRef           uuid.UUID `sql:"type:uuid"`
	Likes               []Like    `gorm:"foreignkey:PostRef"`
}

//TableName ...
func (Post) TableName() string {
	return "posts"
}
