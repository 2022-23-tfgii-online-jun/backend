package monitoring

import (
	"github.com/emur-uy/backend/internal/infra/api/middlewares"
	"github.com/emur-uy/backend/internal/infra/api/middlewares/constants"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/monitoring"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the monitoring-related routes on the given gin.Engine instance.
// It initializes the necessary components, such as the repository, service, and handler,
// to handle monitoring-related operations in a hexagonal architecture.
func RegisterRoutes(e *gin.Engine) {
	// Initialize the repository by creating a new PostgreSQL client.
	repo := postgresql.NewClient()

	// Create a new MonitoringService instance by injecting the repository.
	service := monitoring.NewService(repo)

	// Create a new monitoringHandler instance by injecting the MonitoringService.
	handler := newHandler(service)

	// Group the monitoring routes together.
	monitoringRoutes := e.Group("/api/v1/monitorings")

	// Register user routes requiring authentication and authorization for user role.
	userRoutes := monitoringRoutes.Group("", middlewares.Authenticate(), middlewares.Authorize(constants.RoleUser))
	userRoutes.POST("/", handler.CreateMonitoring)
	userRoutes.GET("/", handler.GetAllMonitorings)

}
