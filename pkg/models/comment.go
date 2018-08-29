package models

import (
	"time"

	"github.com/google/uuid"
)

//Comment ...
type Comment struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Content   string    `sql:"type:text"`
	Author    uuid.UUID `sql:"type:uuid"`
	Likes     []Like    `gorm:"foreignkey:CommentRef"`
}

//TableName ...
func (Comment) TableName() string {
	return "comments"
}
