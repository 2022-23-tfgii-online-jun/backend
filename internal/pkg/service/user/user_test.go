package user_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/service/user"
	"github.com/stretchr/testify/assert"
)

// Define a test UUID for testing purposes.
var testUuid = uuid.MustParse("24df3f36-ca63-11ed-afa1-0242ac120002")

// MockUserRepository is a mock implementation of the UserRepository interface for testing.
type MockUserRepository struct{}

// CreateWithOmit is a mock implementation of the CreateWithOmit method.
func (m *MockUserRepository) CreateWithOmit(omit string, value interface{}) error {
	return nil
}

// Update is a mock implementation of the Update method.
func (m *MockUserRepository) Update(value interface{}) error {
	return nil
}

// First is a mock implementation of the First method.
func (m *MockUserRepository) First(out interface{}, conditions ...interface{}) error {
	if conditions[0] == "uuid= ?" && conditions[1] == testUuid.String() {
		out.(*entity.User).UUID = testUuid
		return nil
	}
	return errors.New("user not found")
}

// UpdateColumns is a mock implementation of the UpdateColumns method.
func (m *MockUserRepository) UpdateColumns(value interface{}, column string, updateValue interface{}) error {
	return nil
}

// FindByUUID is a mock implementation of the FindByUUID method.
func (m *MockUserRepository) FindByUUID(userUUID string) (*entity.User, error) {
	if userUUID == testUuid.String() {
		return &entity.User{
			ID:        1,
			UUID:      testUuid,
			Email:     "test@example.com",
			Password:  "hashed_password",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}
	return nil, errors.New("user not found")
}

// TestCreateUser tests the CreateUser method of the UserService.
func TestCreateUser(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &MockUserRepository{}
	s := user.NewService(mockRepo)

	// Case 1: Valid user, should return HTTP status 201.
	u := &entity.User{
		Email:    "test@example.com",
		Password: "password",
	}
	status, err := s.CreateUser(u)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, status)

	// Case 2: Invalid email, should return HTTP status 500.
	u = &entity.User{
		Email:    "invalid_email",
		Password: "password",
	}
	status, err = s.CreateUser(u)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
}

// TestUpdateUser tests the UpdateUser method of the UserService.
func TestUpdateUser(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &MockUserRepository{}
	s := user.NewService(mockRepo)

	dateOfBirth := "12-12-2022"

	// Case 1: Valid update data, should return HTTP status 200.
	updateData := &entity.UpdateUser{
		FirstName:   stringPtr("John"),
		LastName:    stringPtr("Doe"),
		DateOfBirth: &dateOfBirth,
		Sex:         stringPtr("M"),
		UserType:    stringPtr("admin"),
		City:        stringPtr("Montevideo"),
		Country:     stringPtr("UY"),
	}

	status, err := s.UpdateUser(updateData)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)

	// Case 2: Invalid update data, should return HTTP status 500.
	updateData = &entity.UpdateUser{
		FirstName:   stringPtr("John"),
		LastName:    stringPtr("Doe"),
		DateOfBirth: &dateOfBirth,
		Sex:         stringPtr("M"),
		UserType:    stringPtr("admin"),
		City:        nil,
		Country:     nil,
	}
	status, err = s.UpdateUser(updateData)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestGetUser(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &MockUserRepository{}
	s := user.NewService(mockRepo)
	// Test case 1: user found
	mockUser, err := s.GetUser(testUuid.String())
	assert.Nil(t, err)
	assert.NotNil(t, mockUser)
	assert.Equal(t, testUuid, mockUser.UUID)

	// Test case 2: user not found
	mockUser, err = s.GetUser("not-found-uuid")
	assert.NotNil(t, err)
	assert.Nil(t, mockUser)
}

func TestUpdateActiveStatus(t *testing.T) {
	mockRepo := &MockUserRepository{}
	s := user.NewService(mockRepo)
	// Test case 1: user found and updated successfully
	status, err := s.UpdateActiveStatus(testUuid.String(), true)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)

	// Test case 2: user not found
	status, err = s.UpdateActiveStatus("not-found-uuid", true)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestUpdateBannedStatus(t *testing.T) {
	mockRepo := &MockUserRepository{}
	s := user.NewService(mockRepo)
	// Test case 1: user found and updated successfully
	status, err := s.UpdateBannedStatus(testUuid.String(), true)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)

	// Test case 2: user not found
	status, err = s.UpdateBannedStatus("not-found-uuid", true)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
}
func stringPtr(s string) *string {
	return &s
}
