package db

import (
	"errors"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//GetAlertActionStorage ...
func GetAlertActionStorage(db *gorm.DB) *AlertActionStorage {
	return &AlertActionStorage{db}
}

//AlertActionStorage ...
type AlertActionStorage struct {
	db *gorm.DB
}

//Get ...
func (s *AlertActionStorage) Get(id uuid.UUID) (*models.AlertAction, error) {
	alertaction := &models.AlertAction{}
	if Inst.c.Find(&alertaction, id).RecordNotFound() {
		return nil, errors.New("Alert Action not found")
	}
	return alertaction, nil
}

//GetAll ...
func (s *AlertActionStorage) GetAll() (map[int64]*models.AlertAction, error) {
	var alertactions []models.AlertAction
	s.db.Find(&alertactions)
	if s.db.Error != nil {
		return nil, s.db.Error
	}

	var count int64
	AlertActionMap := make(map[int64]*models.AlertAction)
	for i := range alertactions {
		u := &alertactions[i]
		AlertActionMap[count] = u
		count++
	}
	return AlertActionMap, nil
}

//Insert ...
func (s *AlertActionStorage) Insert(u models.AlertAction) error {
	return nil
}

//Update ...
func (s *AlertActionStorage) Update(u models.AlertAction) error {
	return nil
}

//Remove ...
func (s *AlertActionStorage) Remove(id uuid.UUID) error {
	return nil
}
