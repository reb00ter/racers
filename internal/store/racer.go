package store

import "github.com/reb00ter/racers/internal/models"

type Racer interface {
	First(m *models.Racer) error
	Find(m *[]models.Racer) error
	Create(m *models.Racer) error
	Ping() error
}
