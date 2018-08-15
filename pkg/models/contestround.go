package models

import (
	"time"

	"github.com/google/uuid"
)

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
