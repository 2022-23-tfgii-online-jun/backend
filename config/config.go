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
	AwsBucketName    string `mapstructure:"AWS_BUCKET_NAME"`
	AwsFolderName    string `mapstructure:"AWS_FOLDER_NAME"`
	AwsRegionName    string `mapstructure:"AWS_REGION_NAME"`
	AwsAccessKey     string `mapstructure:"AWS_ACCESS_KEY"`
	AwsSecretKey     string `mapstructure:"AWS_SECRET_KEY"`
	AwsEndpoint      string `mapstructure:"AWS_ENDPOINT"`
	ForecastKey      string `mapstructure:"FORECAST_KEY"`
	ForecastAPI      string `mapstructure:"FORECAST_API"`
	EncryptionKey    string `mapstructure:"ENCRYPTION_KEY"`
}

func Get() Config {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("dev")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("[ReadENV]: cannot read env file: %v", err)
	}

	if viper.ConfigFileUsed() == "" {
		log.Println("dev.env not found, using environment variables")
		err := viper.Unmarshal(&cfg)
		if err != nil {
			log.Fatalln("error unmarshalling config", err)
		}
	}

	doOnce.Do(func() {
		err := viper.Unmarshal(&cfg)
		if err != nil {
			log.Fatalln("error unmarshalling config", err)
		}
	})

	return cfg
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
