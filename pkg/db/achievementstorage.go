package db

import (
	"errors"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//GetAchievementStorage ...
func GetAchievementStorage(db *gorm.DB) *AchievementStorage {
	return &AchievementStorage{db}
}

//AchievementStorage ...
type AchievementStorage struct {
	db *gorm.DB
}

//Get ...
func (s *AchievementStorage) Get(id uuid.UUID) (*models.Achievement, error) {
	achievement := &models.Achievement{}
	if Inst.c.Find(&achievement, id).RecordNotFound() {
		return nil, errors.New("achievement not found")
	}
	return achievement, nil
}

//GetAll ...
func (s *AchievementStorage) GetAll() (map[int64]*models.Achievement, error) {
	var achievements []models.Achievement
	s.db.Find(&achievements)
	if s.db.Error != nil {
		return nil, s.db.Error
	}

	var count int64
	achievementMap := make(map[int64]*models.Achievement)
	for i := range achievements {
		u := &achievements[i]
		achievementMap[count] = u
		count++
	}
	return achievementMap, nil
}

//Insert ...
func (s *AchievementStorage) Insert(u models.Achievement) error {
	return nil
}

//Update ...
func (s *AchievementStorage) Update(u models.Achievement) error {
	return nil
}

//Remove ...
func (s *AchievementStorage) Remove(id uuid.UUID) error {
	return nil
}
