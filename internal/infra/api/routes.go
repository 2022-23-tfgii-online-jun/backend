package api

import (
	healthcheck "github.com/RaMin0/gin-health-check" // Importing health check package for gin
	"github.com/emur-uy/backend/internal/infra/api/answer"
	"github.com/emur-uy/backend/internal/infra/api/article"
	"github.com/emur-uy/backend/internal/infra/api/category"
	"github.com/emur-uy/backend/internal/infra/api/healthservice"
	"github.com/emur-uy/backend/internal/infra/api/medical"
	"github.com/emur-uy/backend/internal/infra/api/question"
	"github.com/emur-uy/backend/internal/infra/api/recipe"
	"github.com/emur-uy/backend/internal/infra/api/reminder"
	"github.com/emur-uy/backend/internal/infra/api/treatment"
	"github.com/emur-uy/backend/internal/infra/api/user"
	"github.com/gin-gonic/gin" // Importing gin package for http web framework
)

// RegisterRoutes registers all the application routes.
func RegisterRoutes(e *gin.Engine) {
	// Add the Sentry middleware
	e.Use(SentryMiddleware())

	// Registering a route for checking the health of the microservice
	// It is recommended to use a dedicated health check package for this purpose
	e.GET("/health", healthcheck.Default())

	// Register user routes
	user.RegisterRoutes(e)
	article.RegisterRoutes(e)
	category.RegisterRoutes(e)
	question.RegisterRoutes(e)
	answer.RegisterRoutes(e)
	recipe.RegisterRoutes(e)
	reminder.RegisterRoutes(e)
	medical.RegisterRoutes(e)
	healthservice.RegisterRoutes(e)
	treatment.RegisterRoutes(e)

}
