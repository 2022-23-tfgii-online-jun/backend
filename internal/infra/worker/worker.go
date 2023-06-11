package worker

import (
	"time"

	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/weather"
	"github.com/go-co-op/gocron"
)

func Start() {
	postgresql.Connect()

	repo := postgresql.NewClient()
	weatherService := weather.NewService(repo)
	weatherWorker := weather.NewWorker(weatherService)

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Hour().Do(weatherWorker.CheckWeather)

	s.StartBlocking()
}
