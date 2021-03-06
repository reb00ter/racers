package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/reb00ter/racers/internal/context"
)

func AppContext(cc *context.AppContext) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc.Context = c
			return h(cc)
		}
	}
}
