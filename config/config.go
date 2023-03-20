package config

import (
	"errors"
	"log"
	"sync"

	"github.com/spf13/viper"
)

// Define a global variable to store the configuration
// and use 'sync.Once' to ensure it is loaded only once
var cfg Config
var doOnce sync.Once

// Config is a structure that contains
// all the configuration variables for the application
type Config struct {
	DatabaseHost     string `mapstructure:"DB_HOST"`
	DatabaseName     string `mapstructure:"DB_NAME"`
	DatabasePort     string `mapstructure:"DB_PORT"`
	DatabaseUser     string `mapstructure:"DB_USER"`
	DatabasePassword string `mapstructure:"DB_PASS"`
	DatabaseTLS      string `mapstructure:"DB_TLS"`
	SentryKey        string `mapstructure:"SENTRY_KEY"`
	GinMode          string `mapstructure:"GIN_MODE"`
	SecretKey        string `mapstructure:"SECRET_KEY"`
	JWTTokenKey      string `mapstructure:"JWT_TOKEN_KEY"`
	JWTTokenExpired  int    `mapstructure:"JWT_TOKEN_EXPIRED"`
}

// Get returns the configuration instance.
// It uses 'sync.Once' to ensure the configuration is loaded only once.
func Get() (Config, error) {
	// Configure Viper to read environment variables and configuration files
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	// Set the appropriate config file based on the environment
	env := viper.GetString("APP_ENV")
	if env == "" {
		env = "dev"
	}
	viper.SetConfigName(env)

	// Set Viper to read environment variables automatically
	viper.AutomaticEnv()

	// Try to read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		// If there is an error, log a message
		log.Printf("[ReadENV]: cannot read env file (%s): %v\n", env, err)
	}

	// Use 'doOnce' to ensure the configuration is loaded only once
	doOnce.Do(func() {
		// Deserialize the environment variables into the 'Config' structure
		err = viper.Unmarshal(&cfg)
		if err != nil {
			// If there is an error, log the error message
			log.Printf("Error unmarshalling config: %v\n", err)
		}
	})

	// Validate the required configuration variables
	if err := validateConfig(cfg); err != nil {
		return Config{}, err
	}

	// Return the configuration instance
	return cfg, nil
}

// validateConfig checks if all the required configuration variables are present
func validateConfig(c Config) error {
	if c.DatabaseHost == "" || c.DatabaseName == "" || c.DatabasePort == "" || c.DatabaseUser == "" || c.DatabasePassword == "" {
		return errors.New("missing required database configuration")
	}

	if c.SentryKey == "" {
		return errors.New("missing Sentry key")
	}

	if c.SecretKey == "" || c.JWTTokenKey == "" || c.JWTTokenExpired == 0 {
		return errors.New("missing required JWT configuration")
	}

	return nil
}