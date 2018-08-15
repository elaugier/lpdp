package models

import (
	"time"

	"github.com/google/uuid"
)

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
