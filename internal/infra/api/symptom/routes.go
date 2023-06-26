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
	symptomRoutes := e.Group("/api/v1/symptoms")

	// Register routes requiring authentication and authorization for admin role.
	adminRoutes := symptomRoutes.Group("", middlewares.Authenticate(), middlewares.Authorize(constants.RoleAdmin))
	adminRoutes.POST("", handler.CreateSymptom)
	adminRoutes.GET("", handler.GetAllSymptoms)

	// Register routes requiring authentication and authorization for user role.
	userRoutes := symptomRoutes.Group("", middlewares.Authenticate(), middlewares.Authorize(constants.RoleUser))
	userRoutes.POST("/add", handler.AddUserToSymptom)
	userRoutes.POST("/remove", handler.RemoveUserFromSymptom)
	userRoutes.GET("", handler.GetSymptomsByUser)
}
