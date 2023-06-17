package treatment

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
var testTreatmentUuid = uuid.MustParse("bfb23f5c-a664-432b-b6cc-b7cd17bacf5b")

type mockTreatmentRepository struct{}

func (m mockTreatmentRepository) Update(value interface{}) error {
	return nil
}

func (m mockTreatmentRepository) First(out interface{}, conditions ...interface{}) error {
	return nil
}

func (m mockTreatmentRepository) Find(out interface{}, conditions ...interface{}) error {
	return nil
}

func (m mockTreatmentRepository) FindItemByIDs(firstID, secondID int, tableName string, column1Name string, column2Name string, dest interface{}) error {
	return nil
}

func (m mockTreatmentRepository) Create(value interface{}) error {
	return nil
}

func (m mockTreatmentRepository) FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error) {
	if uuid == testUserUuid {
		usr := &entity.User{
			ID:   1,
			UUID: testUserUuid,
		}
		return usr, nil
	}
	if uuid == testTreatmentUuid {
		cat := &entity.Treatment{
			ID:   1,
			UUID: testTreatmentUuid,
			Name: "Test Name",
		}
		return cat, nil
	}
	return nil, errors.New("not found")
}

func (m mockTreatmentRepository) CreateWithOmit(omit string, value interface{}) error {
	if value == nil {
		return errors.New("input value cannot be nil")
	}
	return nil
}

func (m mockTreatmentRepository) Delete(out interface{}) error {
	return nil
}

func TestCreateTreatment(t *testing.T) {

	// Create a mock repository
	repo := &mockTreatmentRepository{}

	svc := NewService(repo)

	// Create a test context and request
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	testCases := []struct {
		name        string
		category    *entity.RequestCreateTreatment
		expectError bool
		userUUID    uuid.UUID
	}{
		{"treatment creation successful", &entity.RequestCreateTreatment{Name: "Test treatment"}, false, testUserUuid},
		{"treatment creation failed, user doesn't exist", nil, true, uuid.New()},
		{"treatment creation failed, invalid request", nil, true, testUserUuid},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			treatments, statusCode, err := svc.CreateTreatment(c, tc.userUUID, tc.category)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned and the statusCode is not OK.
				require.Error(t, err)
				assert.NotEqual(t, http.StatusOK, statusCode)
				assert.Nil(t, treatments)
			} else {
				// If no error is expected, ensure there is no error returned and the statusCode is OK.
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, statusCode)
				assert.NotNil(t, treatments)
			}
		})
	}
}

func TestUpdateTreatment(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockTreatmentRepository{}
	s := NewService(mockRepo)

	updateData := &entity.RequestUpdateTreatment{
		Name: "Test Name",
	}

	// Test case 1: treatment found and updated successfully
	status, err := s.UpdateTreatment(testTreatmentUuid, updateData)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)

	// Test case 2: treatment not found
	status, err = s.UpdateTreatment(uuid.New(), updateData)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, status)
}

func TestGetAllTreatment(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockTreatmentRepository{}
	s := NewService(mockRepo)

	// Test case 1: user found & treatment fetched successfully
	_, err := s.GetAllTreatments(testUserUuid)
	assert.Nil(t, err)

	// Test case 2: user not found & treatment fetch failed
	_, err = s.GetAllTreatments(uuid.New())
	assert.NotNil(t, err)
}

func TestDeleteTreatment(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockTreatmentRepository{}
	s := NewService(mockRepo)

	// Test case 1: treatment found and deleted successfully
	status, err := s.DeleteTreatment(testTreatmentUuid)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)

	// Test case 2: treatment not found
	status, err = s.DeleteTreatment(uuid.New())
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, status)
}
