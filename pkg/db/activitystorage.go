package db

import (
	"errors"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//GetActivityStorage ...
func GetActivityStorage(db *gorm.DB) *ActivityStorage {
	return &ActivityStorage{db}
}

//ActivityStorage ...
type ActivityStorage struct {
	db *gorm.DB
}

//Get ...
func (s *ActivityStorage) Get(id uuid.UUID) (*models.Activity, error) {
	activity := &models.Activity{}
	if Inst.c.Find(&activity, id).RecordNotFound() {
		return nil, errors.New("activity not found")
	}
	return activity, nil
}

//GetAll ...
func (s *ActivityStorage) GetAll() (map[int64]*models.Activity, error) {
	var activities []models.Activity
	s.db.Find(&activities)
	if s.db.Error != nil {
		return nil, s.db.Error
	}

	var count int64
	activityMap := make(map[int64]*models.Activity)
	for i := range activities {
		u := &activities[i]
		activityMap[count] = u
		count++
	}
	return activityMap, nil
}

//Insert ...
func (s *ActivityStorage) Insert(u models.Activity) error {
	return nil
}

//Update ...
func (s *ActivityStorage) Update(u models.Activity) error {
	return nil
}

//Remove ...
func (s *ActivityStorage) Remove(id uuid.UUID) error {
	return nil
}
