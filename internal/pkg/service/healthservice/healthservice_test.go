package healthservice

import (
	"errors"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"reflect"
	"testing"
)

type mockHealthServiceRepository struct{}

func (m mockHealthServiceRepository) FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error) {
	return nil, nil
}

func (m mockHealthServiceRepository) Create(value interface{}) error {
	if value == nil || reflect.ValueOf(value).IsNil() {
		return errors.New("nil value")
	}
	return nil
}

func (m mockHealthServiceRepository) Update(value interface{}) error {
	return nil
}

func (m mockHealthServiceRepository) First(out interface{}, conditions ...interface{}) error {
	return nil
}

func (m mockHealthServiceRepository) CreateWithOmit(omit string, value interface{}) error {
	return nil
}

func (m mockHealthServiceRepository) Delete(value interface{}) error {
	return nil
}

func (m mockHealthServiceRepository) Find(dest interface{}, conditions ...interface{}) error {
	return nil
}

func TestCreateHealthService(t *testing.T) {

	// Create a mock repository
	repo := &mockHealthServiceRepository{}

	svc := NewService(repo)

	testCases := []struct {
		name        string
		request     *entity.RequestCreateHealthService
		expectError bool
	}{
		{"health service creation successful", &entity.RequestCreateHealthService{Name: "test"}, false},
		{"health service creation failed, bad request", nil, true},
	}

	// Create a test context and request
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, statusCode, err := svc.CreateHealthService(c, tc.request)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned and the statusCode is not OK.
				require.Error(t, err)
				assert.NotEqual(t, http.StatusOK, statusCode)
			} else {
				// If no error is expected, ensure there is no error returned and the statusCode is OK.
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, statusCode)
			}
		})
	}
}

func TestGetAllHealthServices(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockHealthServiceRepository{}
	s := NewService(mockRepo)

	testCases := []struct {
		name        string
		expectError bool
	}{
		{"fetch health services successful", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := s.GetAllHealthServices()
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned and the healthServices is nil.
				require.Error(t, err)
			} else {
				// If no error is expected, ensure there is no error returned and the healthServices is not nil.
				require.NoError(t, err)
			}
		})
	}
}

func TestAddRatingToHealthService(t *testing.T) {

	// Create a mock repository
	repo := &mockHealthServiceRepository{}

	svc := NewService(repo)

	testCases := []struct {
		name        string
		request     *entity.HealthServiceRating
		expectError bool
	}{
		{"health service rating successful", &entity.HealthServiceRating{HealthServiceID: 1, ReminderID: 1}, false},
		{"health service rating failed, bad request", &entity.HealthServiceRating{HealthServiceID: 0, ReminderID: 0}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			statusCode, err := svc.AddRatingToHealthService(tc.request)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned and the statusCode is not OK.
				require.Error(t, err)
				assert.NotEqual(t, http.StatusOK, statusCode)
			} else {
				// If no error is expected, ensure there is no error returned and the statusCode is OK.
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, statusCode)
			}
		})
	}
}
