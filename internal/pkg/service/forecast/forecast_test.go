package forecast

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type mockForecastRepository struct{}

func (m mockForecastRepository) Create(value interface{}) error {
	return nil
}

func (m mockForecastRepository) GetDistinctCountryAndCityUsers(users *[]entity.User) error {
	return nil
}

func TestGetDistinctCountryAndCityUsers(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockForecastRepository{}
	s := NewService(mockRepo)

	// Test case 1: data found
	_, err := s.GetDistinctCountryAndCityUsers()
	assert.Nil(t, err)
}

func TestCreateForecast(t *testing.T) {

	// Create a mock repository
	repo := &mockForecastRepository{}

	svc := NewService(repo)

	testCases := []struct {
		name        string
		request     *entity.Forecast
		expectError bool
	}{
		{"forecast creation successful", &entity.Forecast{Country: "test"}, false},
		{"forecast creation failed", nil, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := svc.CreateForecast(tc.request)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned and the statusCode is not OK.
				require.Error(t, err)
			} else {
				// If no error is expected, ensure there is no error returned and the statusCode is OK.
				require.NoError(t, err)
			}
		})
	}
}
