package maps

import (
	"errors"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

var testUserUuid = uuid.MustParse("24df3f36-ca63-11ed-afa1-0242ac120002")
var testMapUuid = uuid.MustParse("bfb23f5c-a664-432b-b6cc-b7cd17bacf5b")

type mockMapRepository struct{}

func (m mockMapRepository) Find(out interface{}, conditions ...interface{}) error {
	return nil
}

func (m mockMapRepository) Update(value interface{}) error {
	return nil
}

func (m mockMapRepository) FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error) {
	if uuid == testUserUuid {
		usr := &entity.User{
			ID:   1,
			UUID: testUserUuid,
		}
		return usr, nil
	}
	if uuid == testMapUuid {
		res := &entity.Map{
			ID:   1,
			UUID: testMapUuid,
			Name: "Test Name",
		}
		return res, nil
	}
	return nil, errors.New("not found")
}

func (m mockMapRepository) CreateWithOmit(omit string, value interface{}) error {
	if value == nil {
		return errors.New("input value cannot be nil")
	}
	return nil
}

func (m mockMapRepository) Delete(out interface{}) error {
	return nil
}

func TestCreateMap(t *testing.T) {

	// Create a mock repository
	repo := &mockMapRepository{}

	svc := NewService(repo)

	// Create a test context and request
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	testCases := []struct {
		name        string
		req         *entity.RequestCreateUpdateMap
		expectError bool
	}{
		{"map creation successful", &entity.RequestCreateUpdateMap{Name: "Test"}, false},
		{"map creation failed, invalid request", nil, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, statusCode, err := svc.CreateMap(c, tc.req)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned and the statusCode is not OK.
				require.Error(t, err)
				assert.NotEqual(t, http.StatusOK, statusCode)
				assert.Nil(t, res)
			} else {
				// If no error is expected, ensure there is no error returned and the statusCode is OK.
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, statusCode)
				assert.NotNil(t, res)
			}
		})
	}
}

func TestUpdateMap(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockMapRepository{}
	s := NewService(mockRepo)

	// Create a test context and request
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	testCases := []struct {
		name        string
		req         *entity.RequestCreateUpdateMap
		expectError bool
		mapUId      uuid.UUID
	}{
		{"map updation successful", &entity.RequestCreateUpdateMap{Name: "Test"}, false, testMapUuid},
		{"map updation failed, map doesn't exist", nil, true, uuid.New()},
		{"map updation failed, invalid request", nil, true, testMapUuid},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, statusCode, err := s.UpdateMap(c, tc.mapUId, tc.req)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned and the statusCode is not OK.
				require.Error(t, err)
				assert.NotEqual(t, http.StatusOK, statusCode)
				assert.Nil(t, res)
			} else {
				// If no error is expected, ensure there is no error returned and the statusCode is OK.
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, statusCode)
				assert.NotNil(t, res)
			}
		})
	}
}

func TestGetAllMaps(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockMapRepository{}
	s := NewService(mockRepo)

	// Test case 1: maps fetched successfully
	_, err := s.GetAllMaps()
	assert.Nil(t, err)
}

func TestDeleteMap(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockMapRepository{}
	s := NewService(mockRepo)

	// Create a test context and request
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	testCases := []struct {
		name        string
		expectError bool
		mapUId      uuid.UUID
	}{
		{"map deletion successful", false, testMapUuid},
		{"map deletion failed, map doesn't exist", true, uuid.New()},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			statusCode, err := s.DeleteMap(c, tc.mapUId)
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
