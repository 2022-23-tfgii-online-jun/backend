package treatment

import (
	"github.com/emur-uy/backend/internal/infra/api/middlewares"
	"github.com/emur-uy/backend/internal/infra/api/middlewares/constants"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/treatment"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the treatment-related routes on the given gin.Engine instance.
// It initializes the necessary components, such as the repository, service, and handler,
// to handle treatment-related operations in a hexagonal architecture.
func RegisterRoutes(e *gin.Engine) {
	// Initialize the repository by creating a new PostgreSQL client.
	repo := postgresql.NewClient()

	// Create a new TreatmentService instance by injecting the repository.
	service := treatment.NewService(repo)

	// Create a new treatmentHandler instance by injecting the TreatmentService.
	handler := newHandler(service)

	// Group the treatment routes together.
	treatmentRoutes := e.Group("/api/v1/treatments")

	// Register routes requiring authentication and authorization for user role.
	userRoutes := treatmentRoutes.Group("", middlewares.Authenticate(), middlewares.Authorize(constants.RoleUser))
	userRoutes.POST("", handler.CreateTreatment)
	userRoutes.DELETE("/:uuid", handler.DeleteTreatment)
	userRoutes.PUT("/:uuid", handler.UpdateTreatment)
	userRoutes.GET("", handler.GetAllTreatments)
}
