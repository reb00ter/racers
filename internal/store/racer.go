package store

import (
	"github.com/google/uuid"
	"github.com/reb00ter/racers/internal/models"
)

type Racer interface {
	First(m *models.Racer) error
	Find(m *[]models.Racer) error
	Create(m *models.Racer) error
	Challenge(m *[]models.Racer) error
	VoteUp(racerId uuid.UUID) error
	VoteDown(racerId uuid.UUID) error
	Ping() error
}
