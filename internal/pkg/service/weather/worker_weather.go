package weather

import (
	"fmt"

	"github.com/emur-uy/backend/internal/infra/weather"
)

type Worker struct {
	service *service
}

func NewWorker(service *service) *Worker {
	return &Worker{
		service: service,
	}
}

func (w *Worker) CheckWeather() {
	users, err := w.service.GetDistinctCountryAndCityUsers()
	if err != nil {
		fmt.Println("Error getting distinct users:", err)
		return
	}

	for _, user := range users {
		weather := weather.GetWeatherApi("es", user.Country, user.City)
		fmt.Println(weather)
	}
}
