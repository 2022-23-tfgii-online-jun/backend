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
	Message interface{} `json:"message,omitempty"`
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
	// Step 1: Get the user UUID from the claims.
	// The user UUID is stored in the claims of the JWT token.
	// We use the `MustGet` function to safely retrieve the value from the `Context`.
	//userUUID := c.MustGet("userUUID").(string)
	userUUID := "3a793ec2-0685-4708-a861-2f47cc2dd0ff"
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

func handleError(c *gin.Context, statusCode int, message string, err error) {
	log.Printf("[UserHandler]: %s, %v", message, err)
	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": message,
		"data":    nil,
	})
}

func handleSignUpError(c *gin.Context, statusCode int, message string, err error) {
	log.Printf("[SignUp]: %s, %v", message, err)
	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": message,
		"data":    nil,
	})
}

func handleUserError(c *gin.Context, statusCode int, message string, err error) {
	log.Printf("[UpdateUser]: %s, %v", message, err)
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
