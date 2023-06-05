package healthservice

import (
	"github.com/emur-uy/backend/internal/infra/api/middlewares"
	"github.com/emur-uy/backend/internal/infra/api/middlewares/constants"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/healthservice"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the healthservice-related routes on the given gin.Engine instance.
// It initializes the necessary components, such as the repository, service, and handler,
// to handle healthservice-related operations in a hexagonal architecture.
func RegisterRoutes(e *gin.Engine) {
	// Initialize the repository by creating a new PostgreSQL client.
	repo := postgresql.NewClient()

	// Create a new HealthService instance by injecting the repository.
	service := healthservice.NewService(repo)

	// Create a new healthServiceHandler instance by injecting the HealthService.
	handler := newHandler(service)

	// Group the healthservice routes together.
	healthServiceRoutes := e.Group("/api/v1/healthservices")

	// Register admin routes requiring authentication and authorization for admin role.
	//adminRoutes := healthServiceRoutes.Group("", middlewares.Authenticate(), middlewares.Authorize(constants.RoleAdmin))
	//adminRoutes.POST("", handler.CreateHealthService)

	// Register route for getting all health services accessible to both admin and user roles.
	allowedRoles := []string{constants.RoleAdmin, constants.RoleUser}
	healthServiceRoutes.GET("", middlewares.Authenticate(), middlewares.Authorize(allowedRoles...), handler.GetAllHealthServices)
	healthServiceRoutes.POST("", middlewares.Authenticate(), middlewares.Authorize(allowedRoles...), handler.CreateHealthService)
}
