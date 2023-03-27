package user

import (
	"fmt"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"log"
	"net/http"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

// service is a private struct implementing the ports.UserService interface, which
// encapsulates the business logic related to user operations.
type service struct {
	repo ports.UserRepository // repo is an instance of the UserRepository interface for data persistence.
}

// NewService is a factory function that returns a new service instance, initialized with
// the provided UserRepository for data persistence.
func NewService(repo ports.UserRepository) *service {
	return &service{
		repo: repo,
	}
}

// CreateUser is a method that creates a new user in the database, performing data validation
// and transformations before persisting the new user record.
func (s *service) CreateUser(user *entity.User) (int, error) {
	// Step 1: Validate email address format.
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		log.Printf("error parsing email: %s", err.Error())
		return http.StatusInternalServerError, err
	}

	// Step 2: Encrypt the user's password using bcrypt.
	encryptedPass, err := encryptPassword(user.Password)
	if err != nil {
		log.Printf("error while encrypting the password: %s", err.Error())
		return http.StatusInternalServerError, err
	}

	// Step 3: Modify data before saving to the database.
	user.Password = encryptedPass
	user.IsActive = true

	// Step 4: Save the user record to the database.
	err = s.repo.CreateWithOmit("uuid", user)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

// encryptPassword is a helper function that takes a plain-text password and returns
// its bcrypt hash. This function is used to securely store user passwords.
func encryptPassword(password string) (string, error) {
	cost := 8 // Use bcrypt's default cost of 8 for hashing the password.
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// UpdateUser is a method that updates an existing user in the database,
// performing data validation and transformations before persisting the updated record.
func (s *service) UpdateUser(updateData *entity.UpdateUser) (int, error) {

	if updateData == nil || updateData.City == nil || updateData.Country == nil {
		return http.StatusInternalServerError, fmt.Errorf("invalid data to update")
	}

	// Step 2: Modify data before saving to the database.
	user := &entity.User{
		ID:          1,
		FirstName:   *updateData.FirstName,
		LastName:    *updateData.LastName,
		DateOfBirth: *updateData.DateOfBirth,
		Sex:         *updateData.Sex,
		UserType:    *updateData.UserType,
		City:        *updateData.City,
		Country:     *updateData.Country,
	}

	// Step 3: Save the user record to the database.
	err := s.repo.Update(user)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// GetUser is the service for retrieving information about a user.
func (s *service) GetUser(userUUID string) (*entity.User, error) {
	// Initialize an empty User entity.
	user := &entity.User{}

	// Search for a user with the given UUID in the repository.
	// If a matching user is found, its data will be stored in the `user` variable.
	if err := s.repo.First(user, "uuid= ?", userUUID); err != nil {
		// If there's an error during the search, return a nil User pointer and the error.
		return nil, err
	}

	// If the user is found successfully, return the User pointer and a nil error.
	return user, nil
}

// UpdateActiveStatus updates the active status of a user.
func (s *service) UpdateActiveStatus(userUUID string, isActive bool) (int, error) {
	// Find user by UUID
	user, err := s.repo.FindByUUID(userUUID)
	if err != nil {
		// Return error if the user is not found
		return http.StatusInternalServerError, err
	}

	// Update the "is_active" column of the user in the database
	if err := s.repo.UpdateColumns(user, "is_active", isActive); err != nil {
		// Return error if the update fails
		return http.StatusInternalServerError, err
	}

	// Return the HTTP OK status code if the update is successful
	return http.StatusOK, nil
}

// UpdateBannedStatus updates the banned status of a user.
func (s *service) UpdateBannedStatus(userUUID string, isBanned bool) (int, error) {
	// Find user by UUID
	user, err := s.repo.FindByUUID(userUUID)
	if err != nil {
		// Return error if the user is not found
		return http.StatusInternalServerError, err
	}

	// Update the "is_banned" column of the user in the database
	if err := s.repo.UpdateColumns(user, "is_banned", isBanned); err != nil {
		// Return error if the update fails
		return http.StatusInternalServerError, err
	}

	// Return the HTTP OK status code if the update is successful
	return http.StatusOK, nil
}
