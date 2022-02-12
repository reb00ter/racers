package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reb00ter/racers/internal/context"
	"github.com/reb00ter/racers/internal/core/errors"
	"github.com/reb00ter/racers/internal/models"
)

type (
	Racer          struct{}
	RacerViewModel struct {
		Name string
		ID   string
	}
)

func (ctrl Racer) GetRacer(c echo.Context) error {
	cc := c.(*context.AppContext)
	racerID := c.Param("id")

	racer := models.Racer{ID: racerID}

	err := cc.RacerStore.First(&racer)

	if err != nil {
		b := errors.NewBoom(errors.RacerNotFound, errors.ErrorText(errors.RacerNotFound), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	vm := RacerViewModel{
		Name: racer.Name,
		ID:   racer.ID,
	}

	return c.Render(http.StatusOK, "racer.html", vm)

}

func (ctrl Racer) GetRacerJSON(c echo.Context) error {
	cc := c.(*context.AppContext)
	racerID := c.Param("id")

	racer := models.Racer{ID: racerID}

	err := cc.RacerStore.First(&racer)

	if err != nil {
		b := errors.NewBoom(errors.RacerNotFound, errors.ErrorText(errors.RacerNotFound), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	vm := RacerViewModel{
		Name: racer.Name,
		ID:   racer.ID,
	}

	return c.JSON(http.StatusOK, vm)
}
