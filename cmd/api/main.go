package main

import (
	"fmt"
	"log"

	"github.com/emur-uy/backend/config"
	"github.com/emur-uy/backend/internal/infra/api"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

// defaultPort defines the default server port.
const defaultPort = "8080"

func main() {
	// Get the configuration instance and handle errors.
	cfg := config.Get()

	// Establish a database connection and handle errors.
	postgresql.Connect()

	// Defer a function to close the database connection when the program exits.
	defer func() {
		dbInstance, _ := postgresql.Db.DB()
		_ = dbInstance.Close()
	}()

	// Set the Gin mode based on the configuration.
	gin.SetMode(cfg.GinMode)

	// Initialize Sentry for logging and handle errors.
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.SentryKey,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	// Create the address string with the defaultPort.
	address := fmt.Sprintf(":%s", defaultPort)

	// Start the API server with the address.
	api.Start(address)
}
