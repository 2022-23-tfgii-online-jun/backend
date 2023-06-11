package ports

import "github.com/emur-uy/backend/internal/pkg/entity"

type ForecastService interface {
	GetDistinctCountryAndCityUsers() ([]entity.User, error)
}

type ForecastRepository interface {
	Create(value interface{}) error
	GetDistinctCountryAndCityUsers(users *[]entity.User) error
}
