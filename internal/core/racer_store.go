package core

import (
	"errors"
	"github.com/google/uuid"
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
	return s.DB.Order("rating").Find(m).Error
}

func (s *RacerStore) Challenge(m *[]models.Racer) error {
	query := s.DB.Debug().Raw("SELECT * FROM racers ORDER BY RANDOM() LIMIT ?", 2).Scan(&m)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected != 2 {
		return errors.New("not enough racers to challenge")
	}
	return nil
}

func (s *RacerStore) VoteUp(racerId uuid.UUID) error {
	query := s.DB.Debug().Exec("UPDATE racers SET rating = rating + 1 WHERE id = ?", racerId)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected != 1 {
		return errors.New("racer not found")
	}
	return nil
}

func (s *RacerStore) VoteDown(racerId uuid.UUID) error {
	query := s.DB.Debug().Exec("UPDATE racers SET rating = rating - 1 WHERE id = ?", racerId)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected != 1 {
		return errors.New("racer not found")
	}
	return nil
}

func (s *RacerStore) Ping() error {
	return s.DB.DB().Ping()
}
