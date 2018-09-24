package db

import (
	"errors"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//GetAlertStorage ...
func GetAlertStorage(db *gorm.DB) *AlertStorage {
	return &AlertStorage{db}
}

//AlertStorage ...
type AlertStorage struct {
	db *gorm.DB
}

//Get ...
func (s *AlertStorage) Get(id uuid.UUID) (*models.Alert, error) {
	alert := &models.Alert{}
	if Inst.c.Find(&alert, id).RecordNotFound() {
		return nil, errors.New("Alert not found")
	}
	return alert, nil
}

//GetAll ...
func (s *AlertStorage) GetAll() (map[int64]*models.Alert, error) {
	var alerts []models.Alert
	s.db.Find(&alerts)
	if s.db.Error != nil {
		return nil, s.db.Error
	}

	var count int64
	AlertMap := make(map[int64]*models.Alert)
	for i := range alerts {
		u := &alerts[i]
		AlertMap[count] = u
		count++
	}
	return AlertMap, nil
}

//Insert ...
func (s *AlertStorage) Insert(u models.Alert) error {
	return nil
}

//Update ...
func (s *AlertStorage) Update(u models.Alert) error {
	return nil
}

//Remove ...
func (s *AlertStorage) Remove(id uuid.UUID) error {
	return nil
}
