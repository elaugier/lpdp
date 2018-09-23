package db

import (
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/google/uuid"
)

//Achievements ...
type Achievements struct {
	AchievementFn       func(id uuid.UUID) (*models.Achievement, error)
	CreateAchievementFn func(achievement models.Achievement) (bool, error)
}

//Achievement ...
func (s *Achievements) Achievement(id uuid.UUID) (*models.Achievement, error) {
	achievement := &models.Achievement{}
	Inst.c.Find(&achievement, id)
	return s.AchievementFn(id)
}

//CreateAchievement ...
func (s *Achievements) CreateAchievement(achievement models.Achievement) (bool, error) {
	return s.CreateAchievementFn(achievement)
}
