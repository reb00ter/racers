package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/reb00ter/racers/config"
	"github.com/reb00ter/racers/internal/controller"
	"github.com/reb00ter/racers/internal/core"
	"github.com/reb00ter/racers/internal/models"
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
	healthCtrl := &controller.Healthcheck{}

	// api endpoints
	g := server.Echo.Group("/api")
	g.GET("/racers/:id", racerCtrl.GetRacerJSON)

	// pages
	u := server.Echo.Group("/racers")
	u.GET("", racerListCtrl.GetRacers)
	u.GET("/:id", racerCtrl.GetRacer)

	// metric / health endpoint according to RFC 5785
	server.Echo.GET("/.well-known/health-check", healthCtrl.GetHealthcheck)
	server.Echo.GET("/.well-known/metrics", echo.WrapHandler(promhttp.Handler()))

	// migration for dev
	user := models.Racer{Name: "Peter"}
	mr := server.GetModelRegistry()
	err = mr.Register(user)

	if err != nil {
		server.Echo.Logger.Fatal(err)
	}

	mr.AutoMigrateAll()
	mr.Create(&user)
	// Start server
	go func() {
		if err := server.Start(appConfig.Address); err != nil {
			server.Echo.Logger.Info("shutting down the server")
		}
	}()

	server.GracefulShutdown()
}
