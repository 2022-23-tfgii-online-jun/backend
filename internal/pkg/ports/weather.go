package ports

import "github.com/emur-uy/backend/internal/pkg/entity"

type WeatherService interface {
	GetDistinctCountryAndCityUsers() ([]entity.User, error)
}

type WeatherRepository interface {
	GetDistinctCountryAndCityUsers(users *[]entity.User) error
}
