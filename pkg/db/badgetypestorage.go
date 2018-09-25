package db

import (
	"errors"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//GetBadgeTypeStorage ...
func GetBadgeTypeStorage(db *gorm.DB) *BadgeTypeStorage {
	return &BadgeTypeStorage{db}
}

//BadgeTypeStorage ...
type BadgeTypeStorage struct {
	db *gorm.DB
}

//Get ...
func (s *BadgeTypeStorage) Get(id uuid.UUID) (*models.BadgeType, error) {
	badgetype := &models.BadgeType{}
	if Inst.c.Find(&badgetype, id).RecordNotFound() {
		return nil, errors.New("Badge Type not found")
	}
	return badgetype, nil
}

//GetAll ...
func (s *BadgeTypeStorage) GetAll() (map[int64]*models.BadgeType, error) {
	var badgetypes []models.BadgeType
	s.db.Find(&badgetypes)
	if s.db.Error != nil {
		return nil, s.db.Error
	}

	var count int64
	BadgeTypeMap := make(map[int64]*models.BadgeType)
	for i := range badgetypes {
		u := &badgetypes[i]
		BadgeTypeMap[count] = u
		count++
	}
	return BadgeTypeMap, nil
}

//Insert ...
func (s *BadgeTypeStorage) Insert(u models.BadgeType) error {
	return nil
}

//Update ...
func (s *BadgeTypeStorage) Update(u models.BadgeType) error {
	return nil
}

//Remove ...
func (s *BadgeTypeStorage) Remove(id uuid.UUID) error {
	return nil
}
