package forecast

import (
	"fmt"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
)

type service struct {
	repo ports.ForecastRepository
}

func NewService(repo ports.ForecastRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetDistinctCountryAndCityUsers() ([]entity.User, error) {
	var users []entity.User

	err := s.repo.GetDistinctCountryAndCityUsers(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *service) CreateForecast(createReq *entity.Forecast) error {
	// Crea una nueva entidad de pronóstico
	forecast := &entity.Forecast{
		Country:        createReq.Country,
		State:          createReq.State,
		AvgTemperature: createReq.AvgTemperature,
		MaxTemperature: createReq.MaxTemperature,
		MinTemperature: createReq.MinTemperature,
		Description:    createReq.Description,
		Humidity:       createReq.Humidity,
		Code:           createReq.Code,
		Wind:           createReq.Wind,
		UV:             createReq.UV,
		Date:           createReq.Date,
	}

	// Guarda el registro de pronóstico en la base de datos
	err := s.repo.Create(forecast)
	if err != nil {
		return fmt.Errorf("error creating forecast record: %s", err)
	}

	return nil
}
