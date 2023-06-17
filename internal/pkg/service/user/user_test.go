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
	"github.com/stretchr/testify/require"
)

// Define a test UUID for testing purposes.
var testUuid = uuid.MustParse("24df3f36-ca63-11ed-afa1-0242ac120002")

// Define a test email for testing purposes.
const testEmail = "test@example.com"

// Define a test password for testing purposes.
const testPassword = "$2a$08$dawrvtxhIXVQZ0pz7o809uSsNkSaZp2gW2vwNQERHn37bwFOWIku." // bcrypt hash for "password"

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
	if conditions[0] == "uuid= ?" && conditions[1] == testUuid {
		out.(*entity.User).UUID = testUuid
		return nil
	}
	if conditions[0] == "email = ?" && conditions[1] == testEmail {
		out.(*entity.User).ID = 1
		out.(*entity.User).UUID = testUuid
		out.(*entity.User).Email = testEmail
		out.(*entity.User).Password = testPassword
		out.(*entity.User).IsActive = true
		out.(*entity.User).CreatedAt = time.Now()
		out.(*entity.User).UpdatedAt = time.Now()
		return nil
	}
	if conditions[0] == "user_id = ?" && conditions[1] == 1 {
		out.(*entity.UserRole).UserID = 1
		return nil
	}
	if conditions[0] == "id = ?" && conditions[1] == 1 {
		out.(*entity.Role).ID = 1
		return nil
	}
	return errors.New("not found")
}

// UpdateColumns is a mock implementation of the UpdateColumns method.
func (m *MockUserRepository) UpdateColumns(value interface{}, column string, updateValue interface{}) error {
	return nil
}

// FindByUUID is a mock implementation of the FindByUUID method.
func (m *MockUserRepository) FindByUUID(userUUID uuid.UUID, out interface{}) (interface{}, error) {
	if userUUID == testUuid {
		usr := &entity.User{
			ID:        1,
			UUID:      testUuid,
			Email:     "test@example.com",
			Password:  "hashed_password",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		out.(*entity.User).ID = usr.ID
		out.(*entity.User).UUID = usr.UUID
		out.(*entity.User).Email = usr.Email
		out.(*entity.User).Password = usr.Password
		out.(*entity.User).IsActive = usr.IsActive
		out.(*entity.User).CreatedAt = usr.CreatedAt
		out.(*entity.User).UpdatedAt = usr.UpdatedAt
		return usr, nil
	}
	return nil, errors.New("user not found")
}

// TestLogin checks the Login method of the UserService.
// This test function verifies that the UserService's Login method works as expected.
func TestLogin(t *testing.T) {
	// Set up the mock repository and service.
	mockRepo := &MockUserRepository{}
	s := user.NewService(mockRepo)

	// Define test cases.
	testCases := []struct {
		name        string
		credentials *entity.DefaultCredentials
		expectError bool
	}{
		{"successful login", &entity.DefaultCredentials{Email: "test@example.com", Password: "password"}, false},
		{"failed login - incorrect email", &entity.DefaultCredentials{Email: "nonexistent@example.com", Password: "password"}, true},
		{"failed login - incorrect password", &entity.DefaultCredentials{Email: "test@example.com", Password: "wrong_password"}, true},
	}

	// Execute test cases.
	// Iterate through each test case and run the corresponding test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Login method with the test case's credentials.
			token, err := s.Login(tc.credentials)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned and the token is empty.
				require.Error(t, err)
				assert.Empty(t, token)
			} else {
				// If no error is expected, ensure there is no error returned and the token is not empty.
				require.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
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

	dateOfBirth := time.Now()

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
	mockUser, err := s.GetUser(testUuid)
	assert.Nil(t, err)
	assert.NotNil(t, mockUser)
	assert.Equal(t, testUuid, mockUser.UUID)

	// Test case 2: user not found
	mockUser, err = s.GetUser(uuid.New()) //some random non-existing UUID
	assert.NotNil(t, err)
	assert.Nil(t, mockUser)
}

func TestUpdateActiveStatus(t *testing.T) {
	mockRepo := &MockUserRepository{}
	s := user.NewService(mockRepo)
	// Test case 1: user found and updated successfully
	status, err := s.UpdateActiveStatus(testUuid, true)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)

	// Test case 2: user not found
	status, err = s.UpdateActiveStatus(uuid.New(), true)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestUpdateBannedStatus(t *testing.T) {
	mockRepo := &MockUserRepository{}
	s := user.NewService(mockRepo)
	// Test case 1: user found and updated successfully
	status, err := s.UpdateBannedStatus(testUuid, true) //some random non-existing UUID
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)

	// Test case 2: user not found
	status, err = s.UpdateBannedStatus(uuid.New(), true) //some random non-existing UUID
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestGetUserRole(t *testing.T) {
	mockRepo := &MockUserRepository{}
	s := user.NewService(mockRepo)
	// Test case 1: user role found
	mockUserRole, err := s.GetUserRole(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, mockUserRole.UserID)

	// Test case 2: user role not found
	mockUserRole, err = s.GetUserRole(91)
	assert.NotNil(t, err)
	assert.Nil(t, mockUserRole)
}

func TestGetRole(t *testing.T) {
	mockRepo := &MockUserRepository{}
	s := user.NewService(mockRepo)
	// Test case 1: role found
	mockRole, err := s.GetRole(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, mockRole.ID)

	// Test case 2: role not found
	mockRole, err = s.GetRole(-1)
	assert.NotNil(t, err)
	assert.Nil(t, mockRole)
}

func stringPtr(s string) *string {
	return &s
}
