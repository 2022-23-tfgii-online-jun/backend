package category

import (
	"github.com/emur-uy/backend/internal/infra/api/middlewares"
	"github.com/emur-uy/backend/internal/infra/api/middlewares/constants"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/category"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the category-related routes on the given gin.Engine instance.
// It initializes the necessary components, such as the repository, service, and handler,
// to handle category-related operations in a hexagonal architecture.
func RegisterRoutes(e *gin.Engine) {
	// Initialize the repository by creating a new PostgreSQL client.
	repo := postgresql.NewClient()

	// Create a new CategoryService instance by injecting the repository.
	service := category.NewService(repo)

	// Create a new categoryHandler instance by injecting the CategoryService.
	handler := newHandler(service)

	// Group the category routes together.
	categoryRoutes := e.Group("/api/v1/categories")

	// Register admin routes requiring authentication and authorization for admin role.
	adminRoutes := categoryRoutes.Group("", middlewares.Authenticate(), middlewares.Authorize(constants.RoleAdmin))
	adminRoutes.POST("", handler.CreateCategory)
	adminRoutes.DELETE("/:uuid", handler.DeleteCategory)
	adminRoutes.PUT("/:uuid", handler.UpdateCategory)
}
