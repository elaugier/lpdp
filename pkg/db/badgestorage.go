package db

import (
	"errors"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//GetBadgeStorage ...
func GetBadgeStorage(db *gorm.DB) *BadgeStorage {
	return &BadgeStorage{db}
}

//BadgeStorage ...
type BadgeStorage struct {
	db *gorm.DB
}

//Get ...
func (s *BadgeStorage) Get(id uuid.UUID) (*models.Badge, error) {
	badge := &models.Badge{}
	if Inst.c.Find(&badge, id).RecordNotFound() {
		return nil, errors.New("Badge not found")
	}
	return badge, nil
}

//GetAll ...
func (s *BadgeStorage) GetAll() (map[int64]*models.Badge, error) {
	var badges []models.Badge
	s.db.Find(&badges)
	if s.db.Error != nil {
		return nil, s.db.Error
	}

	var count int64
	BadgeMap := make(map[int64]*models.Badge)
	for i := range badges {
		u := &badges[i]
		BadgeMap[count] = u
		count++
	}
	return BadgeMap, nil
}

//Insert ...
func (s *BadgeStorage) Insert(u models.Badge) error {
	return nil
}

//Update ...
func (s *BadgeStorage) Update(u models.Badge) error {
	return nil
}

//Remove ...
func (s *BadgeStorage) Remove(id uuid.UUID) error {
	return nil
}
