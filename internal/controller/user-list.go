package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reb00ter/racers/internal/context"
	"github.com/reb00ter/racers/internal/core/errors"
	"github.com/reb00ter/racers/internal/models"
)

type (
	RacerList          struct{}
	RacerListViewModel struct {
		Racers []RacerViewModel
	}
)

func (ctrl RacerList) GetRacers(c echo.Context) error {
	cc := c.(*context.AppContext)

	var racers []models.Racer

	err := cc.RacerStore.Find(&racers)

	if err != nil {
		b := errors.NewBoom(errors.RacerNotFound, errors.ErrorText(errors.RacerNotFound), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	viewModel := RacerListViewModel{
		Racers: make([]RacerViewModel, len(racers)),
	}

	for index, racer := range racers {
		viewModel.Racers[index] = RacerViewModel{
			Name: racer.Name,
			ID:   racer.ID,
		}
	}

	return c.Render(http.StatusOK, "racer-list.html", viewModel)

}
