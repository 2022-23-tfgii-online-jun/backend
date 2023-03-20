package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
)

// UserRepository is an interface that represents the contract that any data access
// implementation must satisfy in order to interact with the User data.
type UserRepository interface {

	// CreateWithOmit creates a new user record while omitting specific fields.
	CreateWithOmit(omit string, value interface{}) error
}

// UserService is an interface that represents the contract for the business logic implementation
// related to user operations. This is the primary port in the hexagonal architecture.
type UserService interface {
	// CreateUser creates a new user with the provided user data.
	CreateUser(user *entity.User) (int, error)
}
