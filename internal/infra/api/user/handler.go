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

func newHandler(userService ports.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (u *userHandler) SignUp(c *gin.Context) {
	signUpData := &entity.SignUp{}

	// Step 1: Bind incoming JSON payload to the signUpData struct.
	if err := c.ShouldBindJSON(signUpData); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Step 2: Validate the user's username and password.
	if err := validateUserInput(signUpData); err != nil {
		handleError(c, http.StatusBadRequest, err.Error(), err)
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
		handleError(c, resCode, errMsg, err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "User registered successfully",
	})
}

func handleError(c *gin.Context, statusCode int, message string, err error) {
	log.Printf("[SignUp]: %s, %v", message, err)
	c.JSON(statusCode, gin.H{
		"code":  statusCode,
		"error": message,
	})
}

func validateUserInput(signUpData *entity.SignUp) error {
	if signUpData == nil || signUpData.Email == "" || signUpData.Password == "" {
		return errors.New("username/password is either empty or null")
	}
	return nil
}
