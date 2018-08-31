package models

import (
	"time"

	"github.com/google/uuid"
)

//Comment ...
type Comment struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Content   string    `sql:"type:text"`
	AuthorRef uuid.UUID `sql:"type:uuid"`
	Likes     []Like    `gorm:"foreignkey:CommentRef"`
}

//TableName ...
func (Comment) TableName() string {
	return "comments"
}
