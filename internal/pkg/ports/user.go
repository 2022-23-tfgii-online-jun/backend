package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
)

// UserRepository is an interface that represents the contract that any data access
// implementation must satisfy in order to interact with user data.
type UserRepository interface {

	// CreateWithOmit creates a new user record while omitting specific fields.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// UpdateUser updates an existing user record with the provided user data.
	// Returns an error if the operation fails.
	Update(value interface{}) error
}

// UserService is an interface that represents the contract for the business logic implementation
// related to user operations. This is the primary port in the hexagonal architecture.
type UserService interface {

	// CreateUser creates a new user with the provided user data.
	// Returns an HTTP status code and an error (if any).
	CreateUser(user *entity.User) (int, error)

	// UpdateUser updates an existing user record with the provided user data.
	// Returns an HTTP status code and an error (if any).
	UpdateUser(updateData *entity.UpdateUser) (int, error)
}
