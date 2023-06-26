package symptom

import (
	"github.com/emur-uy/backend/internal/infra/api/middlewares"
	"github.com/emur-uy/backend/internal/infra/api/middlewares/constants"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/symptom"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the symptom-related routes on the given gin.Engine instance.
// It initializes the necessary components, such as the repository, service, and handler,
// to handle symptom-related operations in a hexagonal architecture.
func RegisterRoutes(e *gin.Engine) {
	// Initialize the repository by creating a new PostgreSQL client.
	repo := postgresql.NewClient()

	// Create a new SymptomService instance by injecting the repository.
	service := symptom.NewService(repo)

	// Create a new symptomHandler instance by injecting the SymptomService.
	handler := newHandler(service)

	// Group the symptom routes together.
	symptomRoutes := e.Group("/api/v1/symptoms", middlewares.Authenticate())

	// Register routes for admin role.
	symptomRoutes.POST("", middlewares.Authorize(constants.RoleAdmin), handler.CreateSymptom)
	symptomRoutes.GET("", middlewares.Authorize(constants.RoleAdmin, constants.RoleUser), handler.GetAllSymptoms)

	// Register routes for user role.
	symptomRoutes.POST("/add", middlewares.Authorize(constants.RoleUser), handler.AddUserToSymptom)
	symptomRoutes.POST("/remove", middlewares.Authorize(constants.RoleUser), handler.RemoveUserFromSymptom)
	symptomRoutes.GET("", middlewares.Authorize(constants.RoleUser), handler.GetSymptomsByUser)
}
