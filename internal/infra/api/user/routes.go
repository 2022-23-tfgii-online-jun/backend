package user

import (
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/user"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes is a function that sets up the user-related routes on the given gin.Engine instance.
// It initializes the necessary components, such as the repository, service, and handler, to handle
// user-related operations in a hexagonal architecture.
func RegisterRoutes(e *gin.Engine) {
	// Step 1: Initialize the repository by creating a new PostgreSQL client.
	repo := postgresql.NewClient()

	// Step 2: Create a new UserService instance by injecting the repository.
	service := user.NewService(repo)

	// Step 3: Create a new userHandler instance by injecting the UserService.
	handler := newHandler(service)

	// Step 4: Register the SignUp route with the handler.
	e.POST("/api/v1/users", handler.SignUp)
	e.PATCH("/api/v1/users", handler.UpdateUser)
	e.GET("/api/v1/users", handler.GetUser)
	e.PUT("/api/v1/users/active", handler.SetActiveStatus)
	e.PUT("/api/v1/users/banned", handler.SetBannedStatus)
}
