package db

//https://github.com/hnakamur/api2go-gorm-gin-crud-example/blob/master/storage/storage_user.go

import (
	"errors"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//GetUserStorage ...
func GetUserStorage(db *gorm.DB) *UserStorage {
	return &UserStorage{db}
}

//UserStorage ...
type UserStorage struct {
	db *gorm.DB
}

//Get ...
func (s *UserStorage) Get(id uuid.UUID) (*models.User, error) {
	user := &models.User{}
	if Inst.c.Find(&user, id).RecordNotFound() {
		return nil, errors.New("user not found")
	}
	return user, nil
}

//GetAll ...
func (s *UserStorage) GetAll() (map[int64]*models.User, error) {
	var users []models.User
	s.db.Find(&users)
	if s.db.Error != nil {
		return nil, s.db.Error
	}

	var count int64
	userMap := make(map[int64]*models.User)
	for i := range users {
		u := &users[i]
		userMap[count] = u
		count++
	}
	return userMap, nil
}
