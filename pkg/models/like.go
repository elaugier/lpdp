package models

import (
	"time"

	"github.com/google/uuid"
)

//Like ...
type Like struct {
	ID        uuid.UUID `sql:"type:uuid;primary key;default:gen_random_uuid()" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	PostRef    uuid.UUID `sql:"type:uuid"`
	CommentRef uuid.UUID `sql:"type:uuid"`
}

//TableName ...
func (Like) TableName() string {
	return "likes"
}
