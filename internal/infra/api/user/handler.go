package user

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService ports.UserService
}

type ResponseUserData struct {
	Message *string `json:"message"`
}

// newHandler returns a new instance of userHandler with the given userService.
func newHandler(userService ports.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

// SignUp handles the HTTP request for registering a new user.
func (u *userHandler) SignUp(c *gin.Context) {
	signUpData := &entity.SignUp{}

	// Step 1: Bind incoming JSON payload to the signUpData struct.
	if err := c.ShouldBindJSON(signUpData); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Step 2: Validate the user's username and password.
	if err := validateUserInput(signUpData); err != nil {
		handleSignUpError(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	// Map SignUp data to a User struct
	user := &entity.User{
		Email:    signUpData.Email,
		Password: signUpData.Password,
	}

	// Step 3: Register the user in the database.
	resCode, err := u.userService.CreateUser(user)
	if err != nil {
		errMsg := "An error occurred while creating a new user"
		if strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint") {
			errMsg = "User already exists"
			resCode = http.StatusConflict
		}
		handleSignUpError(c, resCode, errMsg, err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "User registered successfully",
	})
}

// UpdateUser handles the HTTP request for updating an existing user.
func (u *userHandler) UpdateUser(c *gin.Context) {
	updateData := &entity.UpdateUser{}

	// Step 1: Bind incoming JSON payload to the updateData struct.
	if err := c.ShouldBindJSON(updateData); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Step 2: Validate the user's data.
	if err := validateUpdateInput(updateData); err != nil {
		handleUserError(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	// Step 3: Update the user in the database.
	resCode, err := u.userService.UpdateUser(updateData)
	if err != nil {
		handleUserError(c, resCode, "An error occurred while updating the user", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "User updated successfully",
	})
}

// handleError logs the error and sends an error response to the client.
func handleError(c *gin.Context, statusCode int, message string, err error) {
	log.Printf("[UserHandler]: %s, %v", message, err)
	c.JSON(statusCode, gin.H{
		"code":  statusCode,
		"error": message,
	})
}

// handleSignUpError logs the error and sends an error response to the client for SignUp requests.
func handleSignUpError(c *gin.Context, statusCode int, message string, err error) {
	log.Printf("[SignUp]: %s, %v", message, err)
	c.JSON(statusCode, gin.H{
		"code":  statusCode,
		"error": message,
	})
}

// validateUpdateInput checks whether the updateData is valid or not.
func validateUpdateInput(updateData *entity.UpdateUser) error {
	if updateData == nil {
		return errors.New("user data info is either empty or null")
	}
	return nil
}

// handleUserError logs the error and sends an error response to the client for UpdateUser requests.
func handleUserError(c *gin.Context, statusCode int, message string, err error) {
	log.Printf("[UpdateUser]: %s, %v", message, err)
	c.JSON(statusCode, gin.H{
		"code":  statusCode,
		"error": message,
	})
}

// validateUserInput checks whether the signUpData is valid or not.
func validateUserInput(signUpData *entity.SignUp) error {
	if signUpData == nil || signUpData.Email == "" || signUpData.Password == "" {
		return errors.New("username/password is either empty or null")
	}
	return nil
}
