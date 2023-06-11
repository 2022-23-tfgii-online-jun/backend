package medical

import (
	"github.com/emur-uy/backend/internal/infra/api/middlewares"
	"github.com/emur-uy/backend/internal/infra/api/middlewares/constants"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/medical"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the medical-related routes on the given gin.Engine instance.
// It initializes the necessary components, such as the repository, service, and handler,
// to handle medical-related operations in a hexagonal architecture.
func RegisterRoutes(e *gin.Engine) {
	// Initialize the repository by creating a new PostgreSQL client.
	repo := postgresql.NewClient()

	// Create a new MedicalService instance by injecting the repository.
	service := medical.NewService(repo)

	// Create a new medicalHandler instance by injecting the MedicalService.
	handler := newHandler(service)

	// Group the medical routes together.
	medicalRoutes := e.Group("/api/v1/medical")

	// Register route for getting all medical records accessible to both admin and user roles.
	allowedRoles := []string{constants.RoleAdmin, constants.RoleUser}
	medicalRoutes.GET("", middlewares.Authenticate(), middlewares.Authorize(allowedRoles...), handler.GetAllMedicalRecords)

	// Register route for uploading a CSV file accessible only to admin role.
	adminRoutes := medicalRoutes.Group("", middlewares.Authenticate(), middlewares.Authorize(constants.RoleAdmin))
	adminRoutes.POST("", handler.UploadCSV)

	// Register route for adding a rating to a medical record accessible only to user role.
	userRoutes := medicalRoutes.Group("", middlewares.Authenticate(), middlewares.Authorize(constants.RoleUser))
	userRoutes.POST("/rating", handler.AddRatingToMedical)
}
