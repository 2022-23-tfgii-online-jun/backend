package user

import (
	"log"
	"net/http"
	"net/mail"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"

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
