package postgresql

import (
	"fmt" // Importing fmt package for formatted string outputs
	"log" // Importing log package for error logging

	"github.com/emur-uy/backend/config" // Importing config package for accessing database credentials

	"gorm.io/driver/postgres" // Importing postgresql driver package for gorm
	"gorm.io/gorm"            // Importing gorm package for database ORM
)

var Db *gorm.DB // Defining a global variable for gorm DB object

func Connect() {

	// Fetching database credentials from the config package
	conf, err := config.Get()
	if err != nil {
		log.Fatalf("Error getting config: %v\n", err)
	}
	dbHost := conf.DatabaseHost
	dbPort := conf.DatabasePort
	dbName := conf.DatabaseName
	dbUser := conf.DatabaseUser
	dbPass := conf.DatabasePassword
	dbSslMode := conf.DatabaseTLS

	// Validating the database configuration
	if dbHost == "" || dbPort == "" || dbName == "" || dbUser == "" || dbPass == "" || dbSslMode == "" {
		log.Fatalf("Incomplete database configuration found in the config package. Please provide all the required configuration values.") // Log error and exit program if any database configuration is missing
	}

	// Creating database connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPass, dbName, dbSslMode)

	// Establishing database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Unable to initialize gorm connection: %s", err) // Log error and exit program if database connection fails
	}

	Db = db // Assigning the gorm DB object to the global variable for further use
}
