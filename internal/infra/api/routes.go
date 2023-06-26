package api

import (
	healthcheck "github.com/RaMin0/gin-health-check" // Importing health check package for gin
	"github.com/emur-uy/backend/docs"
	"github.com/emur-uy/backend/internal/infra/api/answer"
	"github.com/emur-uy/backend/internal/infra/api/article"
	"github.com/emur-uy/backend/internal/infra/api/category"
	"github.com/emur-uy/backend/internal/infra/api/healthservice"
	"github.com/emur-uy/backend/internal/infra/api/maps"
	"github.com/emur-uy/backend/internal/infra/api/medical"
	"github.com/emur-uy/backend/internal/infra/api/medicalrecord"
	"github.com/emur-uy/backend/internal/infra/api/monitoring"
	"github.com/emur-uy/backend/internal/infra/api/question"
	"github.com/emur-uy/backend/internal/infra/api/recipe"
	"github.com/emur-uy/backend/internal/infra/api/reminder"
	"github.com/emur-uy/backend/internal/infra/api/symptom"
	"github.com/emur-uy/backend/internal/infra/api/treatment"
	"github.com/emur-uy/backend/internal/infra/api/user"
	"github.com/gin-gonic/gin" // Importing gin package for http web framework
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	symptom.RegisterRoutes(e)
	monitoring.RegisterRoutes(e)
	medicalrecord.RegisterRoutes(e)
	maps.RegisterRoutes(e)

	// use ginSwagger middleware to serve the API docs
	e.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	docs.SwaggerInfo.Title = "Emur API Backend"
	docs.SwaggerInfo.Description = "This is an API for the EMUR mobile application microservice."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

}
