package user

import (
	"github.com/emur-uy/backend/internal/infra/api/middlewares"
	"github.com/emur-uy/backend/internal/infra/api/middlewares/constants"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/user"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the user-related routes on the given gin.Engine instance.
// It initializes the necessary components, such as the repository, service, and handler,
// to handle user-related operations in a hexagonal architecture.
func RegisterRoutes(e *gin.Engine) {
	// Initialize the repository by creating a new PostgreSQL client.
	repo := postgresql.NewClient()

	// Create a new UserService instance by injecting the repository.
	service := user.NewService(repo)

	// Create a new userHandler instance by injecting the UserService.
	handler := newHandler(service)

	// Register the SignUp and Login routes with the handler.
	e.POST("/api/v1/users/login", handler.Login)
	e.POST("/api/v1/users/signup", handler.SignUp)

	// Group the user routes together.
	userRoutes := e.Group("/api/v1/users")
	userRoutes.Use(middlewares.Authenticate())

	// Register the GET route for both admin and user.
	userRoutes.GET("", handler.GetUser)
	userRoutes.PUT("", handler.UpdateUser)

	// Register admin routes requiring authorization for admin role.
	adminRoutes := userRoutes.Group("")
	adminRoutes.Use(middlewares.Authorize(constants.RoleAdmin))
	adminRoutes.PUT("/active/:uuid", handler.SetActiveStatus)
	adminRoutes.PUT("/banned/:uuid", handler.SetBannedStatus)
}
