package ports

import (
	"errors"

	"github.com/emur-uy/backend/internal/pkg/entity"
)

var ErrUserNotFound = errors.New("user not found")

// UserRepository is an interface that represents the contract that any data access
// implementation must satisfy in order to interact with user data.
type UserRepository interface {
	FindByUUID(uuid string) (*entity.User, error)

	// CreateWithOmit creates a new user record while omitting specific fields.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// UpdateUser updates an existing user record with the provided user data.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first record that matches the given conditions from the database
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error

	// UpdateColumns updates specified columns of an existing record in the database using the given value.
	// Returns an error if the operation fails.
	UpdateColumns(value interface{}, column string, updateValue interface{}) error
}

// UserService is an interface that represents the contract for the business logic implementation
// related to user operations. This is the primary port in the hexagonal architecture.
type UserService interface {

	// Login authenticates a user and returns a JWT token if successful, or an error if not.
	Login(credentials *entity.DefaultCredentials) (string, error)

	// CreateUser creates a new user with the provided user data.
	// Returns an HTTP status code and an error (if any).
	CreateUser(user *entity.User) (int, error)

	// UpdateUser updates an existing user record with the provided user data.
	// Returns an HTTP status code and an error (if any).
	UpdateUser(updateData *entity.UpdateUser) (int, error)

	// GetUser retrieves user information for the user with the provided UUID.
	// If the user is not found in the database, the error returned should be `ports.ErrUserNotFound`.
	GetUser(userUUID string) (*entity.User, error)

	// UpdateActiveStatus updates the is_active status of the user with the provided UUID.
	// Returns an HTTP status code and an error (if any).
	UpdateActiveStatus(userUUID string, isActive bool) (int, error)

	// UpdateBannedStatus updates the is_banned status of the user with the provided UUID.
	// Returns an HTTP status code and an error (if any).
	UpdateBannedStatus(userUUID string, isBanned bool) (int, error)
}
