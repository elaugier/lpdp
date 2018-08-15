package db

import (
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/jinzhu/gorm"
)

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

//DatabaseInitialization ...
func DatabaseInitialization(conn *gorm.DB) {

	if (!conn.HasTable(&models.Achievement{})) {
		conn.CreateTable(&models.Achievement{})
	}

	if (!conn.HasTable(&models.Activity{})) {
		conn.CreateTable(&models.Activity{})
	}

	return
}
