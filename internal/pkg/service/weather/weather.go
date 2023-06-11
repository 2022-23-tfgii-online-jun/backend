package weather

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
)

type service struct {
	repo ports.WeatherRepository
}

func NewService(repo ports.WeatherRepository) *service {
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
