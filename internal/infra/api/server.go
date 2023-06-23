package api

import (
	"log"

	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	aws "github.com/emur-uy/backend/internal/infra/repositories/spaces"

	"github.com/gin-gonic/gin"
)

// RunServer starts the Gin server with registered routes and middleware.
func RunServer(address string) {
	// Connect to the PostgreSQL database
	postgresql.Connect()

	// Create a Gin server instance with default middlewares
	server := gin.Default()

	// Set CORS configuration as default for all routes
	server.Use(CORS())

	//start aws services
	err := aws.Init()
	if err != nil {
		log.Fatalf("error initializing aws services: %s", err)
	}

	// Register all API routes
	RegisterRoutes(server)

	// Start running the server on the specified address
	err = server.Run(":" + address)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// CORS returns a Gin middleware function with CORS configurations to allow requests from frontend.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set response headers for CORS
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // Set response status code to 204 if request method is OPTIONS
			return
		}

		c.Next() // Move on to the next middleware in the chain
	}
}
