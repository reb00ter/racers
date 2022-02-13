package controller

import (
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/reb00ter/racers/config"
	"github.com/reb00ter/racers/internal/core"
	"github.com/reb00ter/racers/internal/models"
)

var e struct {
	config    *config.Configuration
	logger    *log.Logger
	server    *core.Server
	testRacer *models.Racer
}

func TestMain(m *testing.M) {
	e.config = &config.Configuration{
		ConnectionString: "host=localhost user=gorm dbname=gorm sslmode=disable password=mypassword",
		TemplateDir:      "../templates/*.html",
		LayoutDir:        "../templates/layouts/*.html",
		Dialect:          "postgres",
		RedisAddr:        ":6379",
	}

	e.server = core.NewServer(e.config)

	setup()
	code := m.Run()
	tearDown()

	os.Exit(code)
}

func setup() {
	racerCtrl := &Racer{}
	healthCtrl := &Healthcheck{}

	g := e.server.Echo.Group("/api")
	g.GET("/racers/:id", racerCtrl.GetRacerJSON)

	u := e.server.Echo.Group("/racers")
	u.GET("/:id", racerCtrl.GetRacer)

	e.server.Echo.GET("/.well-known/health-check", healthCtrl.GetHealthcheck)
	e.server.Echo.GET("/.well-known/metrics", echo.WrapHandler(promhttp.Handler()))

	// test data
	racer := models.Racer{Name: "Peter"}
	mr := e.server.GetModelRegistry()
	err := mr.Register(racer)

	if err != nil {
		e.server.Echo.Logger.Fatal(err)
	}

	mr.AutoMigrateAll()
	mr.Save(&racer)

	e.testRacer = &racer
}

func tearDown() {
	e.server.GetModelRegistry().AutoDropAll()
}
