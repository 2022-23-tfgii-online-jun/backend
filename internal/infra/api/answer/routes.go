
// RegisterRoutes sets up the answer-related routes on the given gin.Engine instance.
func RegisterRoutes(e *gin.Engine) {
	// Initialize the repository by creating a new PostgreSQL client.
	repo := postgresql.NewClient()

	// Create a new AnswerService instance by injecting the repository.
	service := answer.NewService(repo)

	// Create a new answerHandler instance by injecting the AnswerService.
	handler := newHandler(service)

	// Group the answer routes together.
	answerRoutes := e.Group("/api/v1/answers")

	// Register routes for creating answers and listing all answers.
	answerRoutes.POST("", middlewares.Authenticate(), handler.CreateAnswer)
}
