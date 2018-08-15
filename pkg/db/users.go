package db

import "github.com/elaugier/lpdp/pkg/models"

//Users ...
type Users struct {
	UserFn       func(id int) (*models.User, error)
	CreateUserFn func(user models.User) (bool, error)
}

//User ...
func (s *Users) User(id int) (*models.User, error) {
	return s.UserFn(id)
}

//CreateUser ...
func (s *Users) CreateUser(user models.User) (bool, error) {
	return s.CreateUserFn(user)
}
