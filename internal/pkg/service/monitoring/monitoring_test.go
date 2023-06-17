package monitoring

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
var testSymptomUuidExistMonitoring = uuid.MustParse("bfb23f5c-a664-432b-b6cc-b7cd17bacf5b")
var testSymptomUuidNewMonitoring = uuid.MustParse("603205fa-dde1-4cc2-8ac7-755d88609f8e")

type mockMonitoringRepository struct{}

func (m mockMonitoringRepository) Create(value interface{}) error {
	return nil
}

func (m mockMonitoringRepository) FindItemByIDs(firstID, secondID int, tableName string, column1Name string, column2Name string, dest interface{}) error {
	if firstID == 1 && secondID == 11 {
		return nil
	}
	if firstID == 1 && secondID == 12 {
		return errors.New("record not found")
	}
	return errors.New("some other error")
}

func (m mockMonitoringRepository) Find(out interface{}, conditions ...interface{}) error {
	return nil
}

func (m mockMonitoringRepository) FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error) {
	if uuid == testUserUuid {
		usr := &entity.User{
			ID:   1,
			UUID: testUserUuid,
		}
		return usr, nil
	}
	if uuid == testSymptomUuidExistMonitoring {
		mon := &entity.Symptom{
			ID:   11,
			UUID: testSymptomUuidExistMonitoring,
			Name: "test",
		}
		return mon, nil
	}
	if uuid == testSymptomUuidNewMonitoring {
		mon := &entity.Symptom{
			ID:    12,
			UUID:  testSymptomUuidNewMonitoring,
			Name:  "test",
			Scale: 1,
		}
		return mon, nil
	}
	return nil, errors.New("not found")
}

func (m mockMonitoringRepository) CreateWithOmit(omit string, value interface{}) error {
	if value == nil {
		return errors.New("input value cannot be nil")
	}
	return nil
}

func TestCreateMonitoring(t *testing.T) {

	// Create a mock repository
	repo := &mockMonitoringRepository{}

	svc := NewService(repo)

	// Create a test context and request
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	testCases := []struct {
		name        string
		monitoring  *entity.RequestCreateMonitoring
		expectError bool
		userUUID    uuid.UUID
	}{
		{"monitoring creation successful", &entity.RequestCreateMonitoring{SymptomUUID: testSymptomUuidNewMonitoring, Scale: 0}, false, testUserUuid},
		{"monitoring creation failed, user doesn't exist", &entity.RequestCreateMonitoring{SymptomUUID: testSymptomUuidNewMonitoring}, true, uuid.New()},
		{"monitoring creation failed, symptom doesn't exist", &entity.RequestCreateMonitoring{SymptomUUID: uuid.New()}, true, testUserUuid},
		{"monitoring creation failed, invalid request", nil, true, testUserUuid},
		{"monitoring creation failed, monitoring already exists", &entity.RequestCreateMonitoring{SymptomUUID: testSymptomUuidExistMonitoring}, true, testUserUuid},
		{"monitoring creation failed, scale comparison failed", &entity.RequestCreateMonitoring{SymptomUUID: testSymptomUuidNewMonitoring, Scale: 2}, true, testUserUuid},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			monitoring, statusCode, err := svc.CreateMonitoring(c, tc.userUUID, tc.monitoring)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned and the statusCode is not OK.
				require.Error(t, err)
				assert.NotEqual(t, http.StatusOK, statusCode)
				assert.Nil(t, monitoring)
			} else {
				// If no error is expected, ensure there is no error returned and the statusCode is OK.
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, statusCode)
				assert.NotNil(t, monitoring)
			}
		})
	}
}

func TestGetAllMonitoring(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockMonitoringRepository{}
	s := NewService(mockRepo)

	// Create a test context and request
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	// Test case 1: user found & monitoring fetched successfully
	_, statusCode, err := s.GetAllMonitorings(c, testUserUuid)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

	// Test case 2: user not found & monitoring fetch failed
	_, statusCode, err = s.GetAllMonitorings(c, uuid.New())
	assert.NotNil(t, err)
	assert.NotEqual(t, http.StatusOK, statusCode)
}
