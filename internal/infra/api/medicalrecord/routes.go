package medicalrecord

import (
	"github.com/emur-uy/backend/internal/infra/api/middlewares"
	"github.com/emur-uy/backend/internal/infra/api/middlewares/constants"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/medicalrecord"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the medical record-related routes on the given gin.Engine instance.
// It initializes the necessary components, such as the repository, service, and handler,
// to handle medical record-related operations in a hexagonal architecture.
func RegisterRoutes(e *gin.Engine) {
	// Initialize the repository by creating a new PostgreSQL client.
	repo := postgresql.NewClient()

	// Create a new MedicalRecordService instance by injecting the repository.
	service := medicalrecord.NewService(repo)

	// Create a new medicalRecordHandler instance by injecting the MedicalRecordService.
	handler := newHandler(service)

	// Group the medical record routes together.
	medicalRecordRoutes := e.Group("/api/v1/medicalrecords")

	// Register routes requiring authentication and authorization for user and admin roles.
	medicalRecordRoutes.Use(middlewares.Authenticate(), middlewares.Authorize(constants.RoleUser, constants.RoleAdmin))

	// POST and PUT endpoints accessible only for user role.
	userRoutes := medicalRecordRoutes.Group("", middlewares.Authorize(constants.RoleUser))
	userRoutes.GET("/", handler.GetMedicalRecord)
	userRoutes.PUT("/:uuid", handler.UpdateMedicalRecord)
	userRoutes.POST("/", handler.CreateMedicalRecord)
}
