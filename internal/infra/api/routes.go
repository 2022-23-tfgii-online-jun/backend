package api

import (
	healthcheck "github.com/RaMin0/gin-health-check" // Importing health check package for gin
	"github.com/gin-gonic/gin"                       // Importing gin package for http web framework
)

// RegisterRoutes registers all the application routes.
func RegisterRoutes(e *gin.Engine) {
	// Registering a route for checking the health of the microservice
	// It is recommended to use a dedicated health check package for this purpose
	e.GET("/health", healthcheck.Default())
}
