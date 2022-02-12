package core

import (
	"github.com/jinzhu/gorm"
	"github.com/reb00ter/racers/internal/models"
)

// RacerStore implements the RacerStore interface
type RacerStore struct {
	DB *gorm.DB
}

func (s *RacerStore) First(m *models.Racer) error {
	return s.DB.First(m).Error
}

func (s *RacerStore) Create(m *models.Racer) error {
	return s.DB.Create(m).Error
}

func (s *RacerStore) Find(m *[]models.Racer) error {
	return s.DB.Find(m).Error
}

func (s *RacerStore) Ping() error {
	return s.DB.DB().Ping()
}
