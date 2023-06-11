package worker

import (
	"time"

	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/forecast"
	"github.com/go-co-op/gocron"
)

func Start() {
	postgresql.Connect()

	repo := postgresql.NewClient()
	forecastService := forecast.NewService(repo)
	forecastWorker := forecast.NewWorker(forecastService)

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Hour().Do(forecastWorker.CheckForecast)

	s.StartBlocking()
}
