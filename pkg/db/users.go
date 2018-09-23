package db

import (
	"errors"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/google/uuid"
)

//Users ...
type Users struct {
}

//GetUser ...
func (s *Users) GetUser(id uuid.UUID) (*models.User, error) {
	user := &models.User{}
	if Inst.c.Find(&user, id).RecordNotFound() {
		return nil, errors.New("user not found")
	}
	return user, nil
}
