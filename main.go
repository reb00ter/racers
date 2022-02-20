package main

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/reb00ter/racers/config"
	"github.com/reb00ter/racers/internal/controller"
	"github.com/reb00ter/racers/internal/core"
	"log"
)

func main() {
	appConfig, err := config.NewConfig()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	// create server
	server := core.NewServer(appConfig)
	// serve files for dev
	server.ServeStaticFiles()

	racerCtrl := &controller.Racer{}
	racerListCtrl := &controller.RacerList{}
	racerChallengeCtrl := controller.NewRacerChallenge()
	healthCtrl := &controller.Healthcheck{}

	// api endpoints
	g := server.Echo.Group("/api")
	g.GET("/racers/:id", racerCtrl.GetRacerJSON)

	// pages
	r := server.Echo.Group("/racers")
	r.GET("", racerListCtrl.GetRacers)
	r.GET("/challenge", racerChallengeCtrl.GetChallenge)
	r.POST("/challenge", racerChallengeCtrl.Vote)
	r.GET("/:id", racerCtrl.GetRacer)

	// metric / health endpoint according to RFC 5785
	server.Echo.GET("/.well-known/health-check", healthCtrl.GetHealthcheck)
	server.Echo.GET("/.well-known/metrics", echo.WrapHandler(promhttp.Handler()))

	// migration for dev
	mr := server.GetModelRegistry()
	mr.AutoMigrateAll()
	// Start server
	go func() {
		if err := server.Start("0.0.0.0:" + appConfig.Port); err != nil {
			server.Echo.Logger.Info("shutting down the server")
		}
	}()

	server.GracefulShutdown()
}
