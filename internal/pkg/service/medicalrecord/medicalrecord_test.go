package medicalrecord

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
var testMedicalRecordUuid = uuid.MustParse("bfb23f5c-a664-432b-b6cc-b7cd17bacf5b")
var testMedicalRecordUuidUnAuthorizedToUpdate = uuid.MustParse("1b06f20e-55f3-48ea-9754-d09a0c58a3fd")

type mockMedicalRecordRepository struct{}

func (m mockMedicalRecordRepository) FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error) {
	if uuid == testUserUuid {
		usr := &entity.User{
			ID:   1,
			UUID: testUserUuid,
		}
		return usr, nil
	}
	if uuid == testMedicalRecordUuid {
		rec := &entity.MedicalRecord{
			ID:     2,
			UUID:   testMedicalRecordUuid,
			UserID: 1,
		}
		return rec, nil
	}
	if uuid == testMedicalRecordUuidUnAuthorizedToUpdate {
		rec := &entity.MedicalRecord{
			ID:     2,
			UUID:   testMedicalRecordUuid,
			UserID: 2,
		}
		return rec, nil
	}
	return nil, errors.New("not found")
}

func (m mockMedicalRecordRepository) CreateWithOmit(omit string, value interface{}) error {
	return nil
}

func (m mockMedicalRecordRepository) Update(value interface{}) error {
	return nil
}

func (m mockMedicalRecordRepository) First(out interface{}, conditions ...interface{}) error {
	if conditions[0] == "user_id = ?" && conditions[1] == 1 {
		return nil
	}
	return errors.New("not found")
}

var medicalRecord = &entity.MedicalRecord{
	ID:                      0,
	UUID:                    uuid.UUID{},
	UserID:                  0,
	HealthCareProvider:      "",
	EmergencyMedicalService: "",
	MultipleSclerosisType:   "",
	LaboralCondition:        "",
	Conmorbidity:            false,
	TreatingNeurologist:     "",
	SupportNetwork:          false,
	IsDisabled:              false,
	EducationalLevel:        "",
}

func TestCreateMedicalRecord(t *testing.T) {

	// Create a mock repository
	repo := &mockMedicalRecordRepository{}

	svc := NewService(repo)

	// Create a test context and request
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	testCases := []struct {
		name        string
		request     *entity.MedicalRecord
		expectError bool
		userUUID    uuid.UUID
	}{
		{"medicalrecord creation successful", medicalRecord, false, testUserUuid},
		{"medicalrecord creation failed, user doesn't exist", nil, true, uuid.New()},
		{"medicalrecord creation failed, invalid request", nil, true, testUserUuid},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, statusCode, err := svc.CreateMedicalRecord(c, tc.userUUID, tc.request)
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

func TestGetMedicalRecord(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockMedicalRecordRepository{}
	s := NewService(mockRepo)

	// Create a test context and request
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	testCases := []struct {
		name        string
		expectError bool
		userUuid    uuid.UUID
	}{
		{"fetch medicalrecord successful", false, testUserUuid},
		{"fetch medicalrecord failed, user doesn't exist", true, uuid.New()},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, statusCode, err := s.GetMedicalRecord(c, tc.userUuid)
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

func TestUpdateMedicalRecord(t *testing.T) {

	// Create a mock repository
	repo := &mockMedicalRecordRepository{}

	svc := NewService(repo)

	// Create a test context and request
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	testCases := []struct {
		name              string
		request           *entity.MedicalRecord
		expectError       bool
		userUuid          uuid.UUID
		medicalRecordUuid uuid.UUID
	}{
		{"medical record update successful", medicalRecord, false, testUserUuid, testMedicalRecordUuid},
		{"medical record update failed, user doesn't exist", medicalRecord, true, uuid.New(), testMedicalRecordUuid},
		{"medical record update failed, medical record doesn't exist", medicalRecord, true, testUserUuid, uuid.New()},
		{"medical record update failed, user unauthorized to update", medicalRecord, true, testUserUuid, testMedicalRecordUuidUnAuthorizedToUpdate},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, statusCode, err := svc.UpdateMedicalRecord(c, tc.userUuid, tc.medicalRecordUuid, tc.request)
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
