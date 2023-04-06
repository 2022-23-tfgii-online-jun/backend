package user

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService ports.UserService
}

type ResponseUserData struct {
	Message interface{} `json:"message,omitempty"`
}

// newHandler returns a new instance of userHandler with the given userService.
func newHandler(userService ports.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

// Login handles user authentication and token generation.
func (u *userHandler) Login(c *gin.Context) {
	credentials := &entity.DefaultCredentials{}

	// 1. Bind the JSON payload to a DefaultCredentials struct.
	if err := c.ShouldBindJSON(credentials); err != nil {
		handleErrorLogin(c, http.StatusBadRequest, "invalid input", err)
		return
	}

	// 2. Authenticate the user and generate a JWT token.
	token, err := u.userService.Login(credentials)
	if err != nil {
		handleError(c, http.StatusBadRequest, "failed to generate token", err)
		return
	}

	// 3. Return the generated token in the response.
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"token": token,
	})
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
		"data":    nil,
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
		"data":    nil,
	})
}

// GetUser handles the HTTP request for getting user information.
func (u *userHandler) GetUser(c *gin.Context) {
	//  1. Get user uuid from context
	userUUID, _ := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))

	// Step 2: Get the user from the database.
	// We call the `GetUser` function from the `UserService` to retrieve the user information from the database.
	user, err := u.userService.GetUser(userUUID)
	if err != nil {
		// If an error occurs while retrieving the user, we return an error response to the client.
		if errors.Is(err, ports.ErrUserNotFound) {
			handleUserError(c, http.StatusNotFound, "User not found", err)
		} else {
			handleUserError(c, http.StatusInternalServerError, "An error occurred while getting the user", err)
		}
		return
	}

	// Step 3: Return the user information.
	// We set the `data` field in the response directly to the user information.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "User information retrieved successfully",
		"data":    user,
	})
}

// SetActiveStatus handles the HTTP request for updating the user's is_active status
func (u *userHandler) SetActiveStatus(c *gin.Context) {
	//  Get user uuid from context
	userUUID, _ := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))

	// Define a struct to hold the request body data.
	type RequestBody struct {
		IsActive bool `json:"is_active"`
	}

	// Parse the request body JSON.
	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		handleUserError(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	// Update the user's is_active status in the database.
	statusCode, err := u.userService.UpdateActiveStatus(userUUID, requestBody.IsActive)
	if err != nil {
		handleUserError(c, statusCode, "An error occurred while updating the user's active status", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "User active status updated successfully",
		"data":    nil,
	})
}

// SetBannedStatus handles the HTTP request for updating the user's is_banned status
func (u *userHandler) SetBannedStatus(c *gin.Context) {
	// Get user uuid from context
	userUUID, _ := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))

	// Define a struct to hold the request body data.
	type RequestBody struct {
		IsBanned bool `json:"is_banned"`
	}

	// Parse the request body JSON.
	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		handleUserError(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	// Update the user's is_banned status in the database.
	statusCode, err := u.userService.UpdateBannedStatus(userUUID, requestBody.IsBanned)
	if err != nil {
		handleUserError(c, statusCode, "An error occurred while updating the user's banned status", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "User banned status updated successfully",
		"data":    nil,
	})
}

// handleError is a generic error handler that logs the error and responds
func handleError(c *gin.Context, statusCode int, message string, err error) {
	// Log the error message and the error itself
	log.Printf("[UserHandler]: %s, %v", message, err)

	// Send the JSON response with the status code and error message
	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": message,
		"data":    nil,
	})
}

// handleSignUpError is an error handler specific to the sign-up process
func handleSignUpError(c *gin.Context, statusCode int, message string, err error) {
	// Log the error message and the error itself
	log.Printf("[SignUp]: %s, %v", message, err)

	// Send the JSON response with the status code and error message
	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": message,
		"data":    nil,
	})
}

// handleError is a utility function to handle errors, log them, and return an appropriate HTTP response.
func handleErrorLogin(c *gin.Context, httpCode int, errMsg string, err error) {
	// Check if the error message should be replaced with the actual error.
	if strings.Contains(err.Error(), "user not found") || strings.Contains(err.Error(), "incorrect/mismatch password") {
		errMsg = err.Error()
	}

	// Log the error message using Sentry.
	sentry.CaptureMessage(fmt.Sprintf("[Login]: %s, %v", errMsg, err))

	// Return an HTTP response with the error message.
	c.JSON(httpCode, gin.H{
		"code":  httpCode,
		"error": errMsg,
	})
}

// handleUserError is an error handler specific to updating user information
func handleUserError(c *gin.Context, statusCode int, message string, err error) {
	// Log the error message and the error itself
	log.Printf("[UpdateUser]: %s, %v", message, err)

	// Send the JSON response with the status code and error message
	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": message,
		"data":    nil,
	})
}

// validateUpdateInput checks whether the updateData is valid or not.
func validateUpdateInput(updateData *entity.UpdateUser) error {
	if updateData == nil {
		return errors.New("user data info is either empty or null")
	}
	return nil
}

// validateUserInput checks whether the signUpData is valid or not.
func validateUserInput(signUpData *entity.SignUp) error {
	if signUpData == nil || signUpData.Email == "" || signUpData.Password == "" {
		return errors.New("username/password is either empty or null")
	}
	return nil
}
