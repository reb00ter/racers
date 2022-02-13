package controller

import (
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/reb00ter/racers/internal/context"
	"github.com/reb00ter/racers/internal/core/middleware"
	"github.com/reb00ter/racers/internal/models"
	"github.com/stretchr/testify/assert"
)

type RacerFakeStore struct{}

func (s *RacerFakeStore) Challenge(m *[]models.Racer) error {
	return nil
}

func (s *RacerFakeStore) VoteUp(racerId uuid.UUID) error {
	return nil
}

func (s *RacerFakeStore) VoteDown(racerId uuid.UUID) error {
	return nil
}

func (s *RacerFakeStore) First(m *models.Racer) error {
	return nil
}
func (s *RacerFakeStore) Find(m *[]models.Racer) error {
	return nil
}
func (s *RacerFakeStore) Create(m *models.Racer) error {
	return nil
}
func (s *RacerFakeStore) Ping() error {
	return nil
}

func TestUserPage(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/users/"+e.testRacer.ID, nil)
	rec := httptest.NewRecorder()
	e.server.Echo.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUnitGetUserJson(t *testing.T) {
	s := echo.New()
	g := s.Group("/api")

	req := httptest.NewRequest(echo.GET, "/api/users/"+e.testRacer.ID, nil)
	rec := httptest.NewRecorder()

	userCtrl := &Racer{}

	cc := &context.AppContext{
		Config:     e.config,
		RacerStore: &RacerFakeStore{},
	}

	s.Use(middleware.AppContext(cc))

	g.GET("/users/:id", userCtrl.GetRacerJSON)
	s.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
