package forecast

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/emur-uy/backend/internal/infra/forecast"
	"github.com/emur-uy/backend/internal/pkg/entity"
)

type Worker struct {
	service *service
}

func NewWorker(service *service) *Worker {
	return &Worker{
		service: service,
	}
}

func (w *Worker) CheckForecast() {
	users, err := w.service.GetDistinctCountryAndCityUsers()
	if err != nil {
		fmt.Println("Error getting distinct users:", err)
		return
	}

	for _, user := range users {
		forecastData, _ := forecast.GetForecast("es", user.Country, user.City)

		// Iterate over each day of the forecast
		for _, d := range forecastData.ForecastInfo.ForecastDay {
			// Parse the date to the custom format
			fileName := getNameIcon(d.Day.Condition.Icon)

			// Create the forecast object
			forecast := &entity.Forecast{
				Country:        user.Country,
				State:          user.City,
				AvgTemperature: int(d.Day.AvgTemperature),
				MaxTemperature: int(d.Day.MaxTempC),
				MinTemperature: int(d.Day.MinTempC),
				Humidity:       d.Day.AvgHumidity,
				Code:           fileName,
				Description:    convertDescription(d.Day.Condition.Text),
				Wind:           d.Day.Wind,
				Date:           d.Date,
				UV:             int(d.Day.UV),
			}

			// Call the service function to create the forecast record in the database
			err := w.service.CreateForecast(forecast)
			if err != nil {
				fmt.Printf("Error creating forecast record for user %s (%s, %s): %s\n", user.Country, user.City, err)
				continue
			}

			// Forecast record created successfully
			fmt.Println("Forecast record created successfully")
		}
	}
}

func getNameIcon(fileName string) int {
	name := strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
	nameToInt, _ := strconv.Atoi(name)
	return nameToInt
}

// ConvertDescription this function convert long description to short description
func convertDescription(text string) string {
	switch text {
	case "Parcialmente nublado", "Cielo cubierto":
		return "Nublado"
	case "Lluvia  moderada a intervalos", "Lluvias ligeras a intervalos", "Ligeras lluvias", "Periodos de lluvia moderada", "Lluvia moderada", "Periodos de fuertes lluvias", "Fuertes lluvias",
		"Ligeras lluvias heladas", "Lluvias heladas fuertes o moderadas", "Ligeras precipitaciones", "Lluvias fuertes o moderadas", "Lluvias torrenciales":
		return "Lluvia"
	case "Nieve moderada a intervalos en las aproximaciones", "Nevadas ligeras a intervalos", "Nevadas ligeras", "Nieve moderada a intervalos", "Nieve moderada", "Fuertes nevadas", "Nevadas intensas",
		"Ligeras precipitaciones de nieve", "Patchy light snow in area with thunder", "Nieve moderada con tormenta en la región", "Nieve moderada o fuertes nevadas con tormenta en la región":
		return "Nieve"
	case "Aguanieve moderada a intervalos en las aproximaciones", "Aguanieve fuerte o moderada":
		return "Aguanieve"
	case "Llovizna helada a intervalos en las aproximaciones", "Llovizna a intervalos", "Llovizna helada", "Fuerte llovizna helada":
		return "LLovizna"
	case "Cielos tormentosos en las aproximaciones", "Intervalos de lluvias ligeras con tomenta en la región", "Lluvias con tormenta fuertes o moderadas en la región":
		return "Tormenta"
	case "Chubascos de nieve", "Ligeros chubascos de aguanieve", "Chubascos de aguanieve fuertes o moderados":
		return "Chubascos"
	case "Niebla moderada":
		return "Nieblina"
	case "Ligeros chubascos acompañados de granizo", "Chubascos fuertes o moderados acompañados de granizo":
		return "Granizo"
	}
	return text
}
