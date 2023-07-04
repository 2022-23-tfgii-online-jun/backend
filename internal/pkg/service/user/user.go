package user

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
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
const (
	JWTCost = 8
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
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		log.Printf("error parsing email: %s", err.Error())
		return http.StatusInternalServerError, err
	}

	encryptedPass, err := encryptPassword(user.Password)
	if err != nil {
		log.Printf("error while encrypting the password: %s", err.Error())
		return http.StatusInternalServerError, err
	}

	encryptedFirstName, err := encryptString(user.FirstName)
	if err != nil {
		log.Printf("error while encrypting the first name: %s", err.Error())
		return http.StatusInternalServerError, err
	}

	encryptedLastName, err := encryptString(user.LastName)
	if err != nil {
		log.Printf("error while encrypting the last name: %s", err.Error())
		return http.StatusInternalServerError, err
	}

	encryptedProfileImage, err := encryptString(user.ProfileImage)
	if err != nil {
		log.Printf("error while encrypting the profile image: %s", err.Error())
		return http.StatusInternalServerError, err
	}

	user.Password = encryptedPass
	user.FirstName = encryptedFirstName
	user.LastName = encryptedLastName
	user.ProfileImage = encryptedProfileImage
	user.IsActive = true

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

	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		return http.StatusInternalServerError, err
	}

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
		// Handle date parsing error
	}

	// Encrypt the updated fields
	encryptedFirstName, err := encryptString(*updateData.FirstName)
	if err != nil {
		log.Printf("error while encrypting the first name: %s", err.Error())
		return http.StatusInternalServerError, err
	}
	encryptedLastName, err := encryptString(*updateData.LastName)
	if err != nil {
		log.Printf("error while encrypting the last name: %s", err.Error())
		return http.StatusInternalServerError, err
	}

	// Update the user fields with the encrypted values
	user.FirstName = encryptedFirstName
	user.LastName = encryptedLastName
	user.DateOfBirth = dateOfBirth
	user.Sex = *updateData.Sex
	user.UserType = *updateData.UserType
	user.City = *updateData.City
	user.Country = *updateData.Country

	err = s.repo.Update(user)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// GetUser is the service for retrieving information about a user.
func (s *service) GetUser(userUUID uuid.UUID) (*entity.User, error) {
	user := &entity.User{}

	if err := s.repo.First(user, "uuid= ?", userUUID); err != nil {
		return nil, err
	}

	decryptedFirstName, err := decryptString(user.FirstName)
	if err != nil {
		log.Printf("error while decrypting the first name: %s", err.Error())
		return nil, err
	}

	decryptedLastName, err := decryptString(user.LastName)
	if err != nil {
		log.Printf("error while decrypting the last name: %s", err.Error())
		return nil, err
	}

	decryptedProfileImage, err := decryptString(user.ProfileImage)
	if err != nil {
		log.Printf("error while decrypting the profile image: %s", err.Error())
		return nil, err
	}

	user.FirstName = decryptedFirstName
	user.LastName = decryptedLastName
	user.ProfileImage = decryptedProfileImage

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

func encryptString(text string) (string, error) {
	plaintext := []byte(text)
	block, err := aes.NewCipher([]byte(config.Get().EncryptionKey))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func decryptString(ciphertext string) (string, error) {
	data, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(config.Get().EncryptionKey))
	if err != nil {
		return "", err
	}

	if len(data) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := data[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data[aes.BlockSize:], data[aes.BlockSize:])

	return string(data[aes.BlockSize:]), nil
}
