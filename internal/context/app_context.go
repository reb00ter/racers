package context

import (
	"github.com/labstack/echo/v4"
	"github.com/reb00ter/racers/config"
	"github.com/reb00ter/racers/internal/i18n"
	"github.com/reb00ter/racers/internal/store"
)

// AppContext is the new context in the request / response cycle
// We can use the db store, cache and central configuration
type AppContext struct {
	echo.Context
	RacerStore store.Racer
	Cache      store.Cache
	Config     *config.Configuration
	Loc        i18n.I18ner
}
