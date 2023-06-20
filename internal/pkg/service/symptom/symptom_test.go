package symptom

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
var testSymptomUuid = uuid.MustParse("bfb23f5c-a664-432b-b6cc-b7cd17bacf5b")
var testUserUuidWithoutSymptom = uuid.MustParse("1b06f20e-55f3-48ea-9754-d09a0c58a3fd")

type mockSymptomRepository struct{}

func (m mockSymptomRepository) Create(value interface{}) error {
	return nil
}

func (m mockSymptomRepository) CreateWithOmit(omit string, value interface{}) error {
	return nil
}

func (m mockSymptomRepository) Update(value interface{}) error {
	return nil
}

func (m mockSymptomRepository) First(out interface{}, conditions ...interface{}) error {
	return nil
}

func (m mockSymptomRepository) Find(out interface{}, conditions ...interface{}) error {
	if len(conditions) == 0 {
		_, ok := out.(*[]*entity.Symptom)
		if ok {
			return nil
		}
	}
	if len(conditions) > 1 && conditions[0] == "user_id = ?" && conditions[1] == 1 {
		outSlice, _ := out.(*[]*entity.SymptomUser)
		if outSlice == nil {
			outSlice = &[]*entity.SymptomUser{} // Initialize outSlice if it is nil
		}
		newSymptomUser := &entity.SymptomUser{
			SymptomID: 99,
		}
		*outSlice = append(*outSlice, newSymptomUser)
		return nil
	}
	if len(conditions) > 1 && conditions[0] == "user_id = ?" && conditions[1] == 2 {
		outSlice, _ := out.(*[]*entity.SymptomUser)
		if outSlice == nil {
			outSlice = &[]*entity.SymptomUser{} // Initialize outSlice if it is nil
		}
		newSymptomUser := &entity.SymptomUser{
			SymptomID: 9,
		}
		*outSlice = append(*outSlice, newSymptomUser)
		return nil
	}
	if len(conditions) > 1 && conditions[0] == "id = ?" && conditions[1] == 99 {
		outSlice, _ := out.(*[]*entity.Symptom)
		if outSlice == nil {
			outSlice = &[]*entity.Symptom{} // Initialize outSlice if it is nil
		}
		newSymptom := &entity.Symptom{
			ID:   1,
			UUID: testSymptomUuid,
		}
		*outSlice = append(*outSlice, newSymptom)
		return nil
	}
	if len(conditions) > 2 && conditions[0] == "user_id = ? AND symptom_id = ?" && conditions[1] == 1 && conditions[2] == 1 {
		outSlice, _ := out.(*[]*entity.SymptomUser)
		if outSlice == nil {
			outSlice = &[]*entity.SymptomUser{} // Initialize outSlice if it is nil
		}
		newSymptomUser := &entity.SymptomUser{
			SymptomID: 99,
		}
		*outSlice = append(*outSlice, newSymptomUser)
		return nil
	}
	return errors.New("not found")
}

func (m mockSymptomRepository) Delete(out interface{}) error {
	return nil
}

func (m mockSymptomRepository) FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error) {
	if uuid == testUserUuid {
		usr := &entity.User{
			ID:   1,
			UUID: testUserUuid,
		}
		return usr, nil
	}
	if uuid == testSymptomUuid {
		res := &entity.Symptom{
			ID:   1,
			UUID: testSymptomUuid,
			Name: "Test Name",
		}
		return res, nil
	}
	if uuid == testUserUuidWithoutSymptom {
		usr := &entity.User{
			ID:   2,
			UUID: testUserUuidWithoutSymptom,
		}
		return usr, nil
	}
	return nil, errors.New("not found")
}

func TestCreateSymptom(t *testing.T) {

	// Create a mock repository
	repo := &mockSymptomRepository{}

	svc := NewService(repo)

	// Create a test context and request
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	testCases := []struct {
		name        string
		request     *entity.RequestCreateSymptom
		expectError bool
	}{
		{"symptom creation successful", &entity.RequestCreateSymptom{Name: "Test treatment"}, false},
		{"symptom creation failed, invalid request", nil, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, statusCode, err := svc.CreateSymptom(c, tc.request)
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

func TestGetAllSymptoms(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockSymptomRepository{}
	s := NewService(mockRepo)

	// Test case 1: symptoms fetched successfully
	_, err := s.GetAllSymptoms()
	assert.Nil(t, err)
}

func TestAddUserToSymptom(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockSymptomRepository{}
	s := NewService(mockRepo)

	testCases := []struct {
		name        string
		request     *entity.RequestCreateSymptomUser
		expectError bool
		userUuid    uuid.UUID
	}{
		{"add user to symptom successful", &entity.RequestCreateSymptomUser{SymptomUUID: testSymptomUuid}, false, testUserUuid},
		{"add user to symptom failed, user doesn't exist", nil, true, uuid.New()},
		{"add user to symptom failed, symptom doesn't exist", &entity.RequestCreateSymptomUser{SymptomUUID: uuid.New()}, true, testUserUuid},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			statusCode, err := s.AddUserToSymptom(tc.userUuid, tc.request)
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

func TestGetSymptomsByUser(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockSymptomRepository{}
	s := NewService(mockRepo)

	testCases := []struct {
		name        string
		expectError bool
		userUuid    uuid.UUID
	}{
		{"get symptom by user successful", false, testUserUuid},
		{"get symptom by user failed, user doesn't exist", true, uuid.New()},
		{"get symptom by user failed, symptom user doesn't exist", true, testUserUuidWithoutSymptom},
		{"get symptom by user failed, symptom doesn't exist", true, testUserUuidWithoutSymptom},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := s.GetSymptomsByUser(tc.userUuid)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned
				require.Error(t, err)
			} else {
				// If no error is expected, ensure there is no error returned
				require.NoError(t, err)
			}
		})
	}
}

func TestRemoveUserFromSymptom(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockSymptomRepository{}
	s := NewService(mockRepo)

	testCases := []struct {
		name        string
		request     *entity.RequestCreateSymptomUser
		expectError bool
		userUuid    uuid.UUID
	}{
		{"remove user from symptom successful", &entity.RequestCreateSymptomUser{SymptomUUID: testSymptomUuid}, false, testUserUuid},
		{"remove user from symptom failed, user doesn't exist", nil, true, uuid.New()},
		{"remove user from symptom failed, symptom doesn't exist", &entity.RequestCreateSymptomUser{SymptomUUID: uuid.New()}, true, testUserUuid},
		{"remove user from symptom failed, symptom user doesn't exist", &entity.RequestCreateSymptomUser{SymptomUUID: testSymptomUuid}, true, testUserUuidWithoutSymptom},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			statusCode, err := s.RemoveUserFromSymptom(tc.userUuid, tc.request)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned
				require.Error(t, err)
				assert.NotEqual(t, http.StatusOK, statusCode)
			} else {
				// If no error is expected, ensure there is no error returned
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, statusCode)
			}
		})
	}
}
