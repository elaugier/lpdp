package models

import (
	"time"

	"github.com/google/uuid"
)

//Post ...
type Post struct {
	ID                  uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	DeletedAt           time.Time `json:"deleted_at"`
	PublishedAt         time.Time `json:"published_at"`
	ApprovedAt          time.Time `json:"approved_at"`
	ApprouvedBy         uuid.UUID `sql:"type:uuid" json:"approved_by"`
	ApprouvedByUserName string    `gorm:"type:varchar(100)" json:"approved_by_username"`
	Title               string    `gorm:"type:varchar(255)" json:"title"`
	Summary             string    `sql:"type:text" json:"summary"`
	Content             string    `sql:"type:text" json:"content"`
	AuthorRef           uuid.UUID `sql:"type:uuid" json:"author_ref"`
	Likes               []Like    `gorm:"foreignkey:PostRef"`
}

//TableName ...
func (Post) TableName() string {
	return "posts"
}
