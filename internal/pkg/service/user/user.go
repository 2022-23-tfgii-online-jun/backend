package user

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"time"

	"github.com/emur-uy/backend/config"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

// JWTCost is the cost factor for hashing passwords using bcrypt.
const JWTCost = 8

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

// Login authenticates a user against the database.
func (s *service) Login(credentials *entity.DefaultCredentials) (string, error) {
	user, err := s.findUserByEmail(credentials.Email)
	if err != nil {
		return "", err
	}

	if err := s.verifyPassword(user.Password, credentials.Password); err != nil {
		return "", err
	}

	return s.generateJWTToken(user)
}

// findUserByEmail retrieves a user from the database by email.
func (s *service) findUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	if err := s.repo.First(user, "email = ?", email); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return user, nil
}

// verifyPassword checks if the provided password matches the stored hash.
func (s *service) verifyPassword(userPassword, credentialsPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(credentialsPassword)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return fmt.Errorf("incorrect/mismatch password")
		}
		return err
	}
	return nil
}

// generateJWTToken generates a JWT token with custom claims for the authenticated user.
func (s *service) generateJWTToken(user *entity.User) (string, error) {
	type jwtCustomClaims struct {
		Email    string    `json:"email"`
		UserUIID uuid.UUID `json:"user_uuid"`
		Role     string    `json:"role"`
		jwt.StandardClaims
	}

	jwtKey := []byte(config.Get().JWTTokenKey)
	expirationTime := time.Now().Add(time.Duration(config.Get().JWTTokenExpired) * time.Hour)

	userRoleData, _ := s.GetUserRole(user.ID)
	var roleData *entity.Role
	if userRoleData != nil {
		roleData, _ = s.GetRole(userRoleData.RoleID)
	}

	var role string
	if roleData != nil {
		role = roleData.Role
	}

	claims := &jwtCustomClaims{
		Email:    user.Email,
		UserUIID: user.UUID,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := jwtToken.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return "Bearer " + strToken, nil
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

	log.Printf("Creating user with values: %+v", user)
	return http.StatusCreated, nil
}

// encryptPassword is a helper function that takes a plain-text password and returns
// its bcrypt hash. This function is used to securely store user passwords.
func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), JWTCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// UpdateUser is a method that updates an existing user in the database,
// performing data validation and transformations before persisting the updated record.
func (s *service) UpdateUser(userUUID uuid.UUID, updateData *entity.UpdateUser) (int, error) {
	user := &entity.User{}

	// Find user by UUID
	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		// Return error if the user is not found
		return http.StatusInternalServerError, err
	}

	// Perform type assertion to convert foundUser to *entity.User
	user, ok := foundUser.(*entity.User)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("type assertion failed")
	}

	if updateData == nil || updateData.City == nil || updateData.Country == nil {
		return http.StatusInternalServerError, fmt.Errorf("invalid data to update")
	}

	layout := "02-01-2006"
	dateOfBirth, err := time.Parse(layout, updateData.DateOfBirth)
	if err != nil {
		// Manejar el error de análisis de fecha
	}

	// Modify data before saving to the database.
	user.FirstName = *updateData.FirstName
	user.LastName = *updateData.LastName
	user.DateOfBirth = dateOfBirth
	user.Sex = *updateData.Sex
	user.UserType = *updateData.UserType
	user.City = *updateData.City
	user.Country = *updateData.Country

	// Save the user record to the database.
	err = s.repo.Update(user)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// GetUser is the service for retrieving information about a user.
func (s *service) GetUser(userUUID uuid.UUID) (*entity.User, error) {
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
func (s *service) UpdateActiveStatus(userUUID uuid.UUID, isActive bool) (int, error) {
	user := &entity.User{}

	// Find user by UUID
	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		// Return error if the user is not found
		return http.StatusInternalServerError, err
	}

	// Perform type assertion to convert foundUser to *entity.User
	user, ok := foundUser.(*entity.User)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("type assertion failed")
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
func (s *service) UpdateBannedStatus(userUUID uuid.UUID, isBanned bool) (int, error) {
	user := &entity.User{}

	// Find user by UUID
	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		// Return error if the user is not found
		return http.StatusInternalServerError, err
	}

	// Perform type assertion to convert foundUser to *entity.User
	user, ok := foundUser.(*entity.User)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("type assertion failed")
	}

	// Update the "is_banned" column of the user in the database
	if err := s.repo.UpdateColumns(user, "is_banned", isBanned); err != nil {
		// Return error if the update fails
		return http.StatusInternalServerError, err
	}

	// Return the HTTP OK status code if the update is successful
	return http.StatusOK, nil
}

// GetUserRole is the service for retrieving information about a user role.
func (s *service) GetUserRole(userID int) (*entity.UserRole, error) {
	// Initialize an empty UserRole entity.
	userRole := &entity.UserRole{}

	// Search for a user role with the given ID in the repository.
	// If a matching user role is found, its data will be stored in the `userRole` variable.
	if err := s.repo.First(userRole, "user_id = ?", userID); err != nil {
		// If there's an error during the search, return a nil UserRole pointer and the error.
		return nil, err
	}

	// If the user role is found successfully, return the UserRole pointer and a nil error.
	return userRole, nil
}

// GetRole is the service for retrieving information about a role.
func (s *service) GetRole(roleID int) (*entity.Role, error) {
	// Initialize an empty Role entity.
	role := &entity.Role{}

	// Search for a role with the given ID in the repository.
	// If a matching role is found, its data will be stored in the `role` variable.
	if err := s.repo.First(role, "id = ?", roleID); err != nil {
		// If there's an error during the search, return a nil role pointer and the error.
		return nil, err
	}

	// If the role is found successfully, return the role pointer and a nil error.
	return role, nil
}
