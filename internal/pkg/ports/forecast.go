package ports

import "github.com/emur-uy/backend/internal/pkg/entity"

// ForecastService is an interface defining a contract for business logic operators related to Forecasts.
// It works with the entity layer to manipulate Forecast data.
type ForecastService interface {
	// GetDistinctCountryAndCityUsers retrieves distinct users based on their country and city.
	// Returns a slice of User entities and an error if any occurred.
	GetDistinctCountryAndCityUsers() ([]entity.User, error)
}

// ForecastRepository is an interface that acts as a contract for the data access layer,
// requiring implementations to provide methods for querying and modifying forecast data.
type ForecastRepository interface {
	// Create takes a new Forecast value and adds it to the data store.
	// Returns an error if the operation fails.
	Create(value interface{}) error

	// GetDistinctCountryAndCityUsers retrieves distinct users based on their country and city,
	// modifying the provided slice of Users.
	// Returns an error if the operation fails.
	GetDistinctCountryAndCityUsers(users *[]entity.User) error
}
