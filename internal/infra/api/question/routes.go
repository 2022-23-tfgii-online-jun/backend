package question

import (
	"github.com/emur-uy/backend/internal/infra/api/middlewares"
	"github.com/emur-uy/backend/internal/infra/api/middlewares/constants"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/question"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the question-related routes on the given gin.Engine instance.
// It initializes the necessary components, such as the repository, service, and handler,
// to handle question-related operations in a hexagonal architecture.
func RegisterRoutes(e *gin.Engine) {
	// Initialize the repository by creating a new PostgreSQL client.
	repo := postgresql.NewClient()

	// Create a new QuestionService instance by injecting the repository.
	service := question.NewService(repo)

	// Create a new questionHandler instance by injecting the QuestionService.
	handler := newHandler(service)

	// Group the question routes together.
	questionRoutes := e.Group("/api/v1/questions")

	// Register route for getting all questions accessible to both admin and user roles.
	allowedRoles := []string{constants.RoleAdmin, constants.RoleUser}
	questionRoutes.GET("", middlewares.Authenticate(), middlewares.Authorize(allowedRoles...), handler.GetAllQuestions)
	questionRoutes.POST("", middlewares.Authenticate(), middlewares.Authorize(allowedRoles...), handler.CreateQuestion)

}
