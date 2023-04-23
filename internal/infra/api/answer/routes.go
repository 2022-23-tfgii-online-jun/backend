package answer

import (
	"github.com/emur-uy/backend/internal/infra/api/middlewares"
	"github.com/emur-uy/backend/internal/infra/api/middlewares/constants"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/answer"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the answer-related routes on the given gin.Engine instance.
func RegisterRoutes(e *gin.Engine) {
	// Initialize the repository by creating a new PostgreSQL client.
	repo := postgresql.NewClient()

	// Create a new AnswerService instance by injecting the repository.
	service := answer.NewService(repo)

	// Create a new answerHandler instance by injecting the AnswerService.
	handler := newHandler(service)

	// Group the answer routes together.
	answerRoutes := e.Group("/api/v1/questions/:question_uuid/answer")

	// Register route for getting all questions accessible to both admin and user roles.
	allowedRoles := []string{constants.RoleAdmin, constants.RoleUser}

	// Register route for answering a question using the answerHandler
	answerRoutes.POST("", middlewares.Authenticate(), middlewares.Authorize(allowedRoles...), handler.CreateAnswer)
}
