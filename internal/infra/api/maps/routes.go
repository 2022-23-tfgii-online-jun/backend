package maps

import (
	"github.com/emur-uy/backend/internal/infra/api/middlewares"
	"github.com/emur-uy/backend/internal/infra/api/middlewares/constants"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/maps"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the map-related routes on the given gin.Engine instance.
// It initializes the necessary components, such as the repository, service, and handler,
// to handle map-related operations in a hexagonal architecture.
func RegisterRoutes(e *gin.Engine) {
	// Initialize the repository by creating a new PostgreSQL client.
	repo := postgresql.NewClient()

	// Create a new MapService instance by injecting the repository.
	service := maps.NewService(repo)

	// Create a new mapHandler instance by injecting the MapService.
	handler := newHandler(service)

	// Group the map routes together.
	mapRoutes := e.Group("/api/v1/maps")

	// Register admin routes requiring authentication and authorization for admin role.
	adminRoutes := mapRoutes.Group("", middlewares.Authenticate(), middlewares.Authorize(constants.RoleAdmin))
	adminRoutes.POST("", handler.CreateMap)
	adminRoutes.PUT("/:uuid", handler.UpdateMap)
	adminRoutes.DELETE("/:uuid", handler.DeleteMap)

	// Register route for getting all maps accessible to both admin and user roles.
	allowedRoles := []string{constants.RoleAdmin, constants.RoleUser}
	mapRoutes.GET("", middlewares.Authenticate(), middlewares.Authorize(allowedRoles...), handler.GetAllMaps)
}
