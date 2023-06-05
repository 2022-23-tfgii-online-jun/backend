package healthservice

import (
	"fmt"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type service struct {
	repo ports.HealthServiceRepository
}

// NewService returns a new instance of the health service with the given health service repository.
func NewService(healthServiceRepo ports.HealthServiceRepository) ports.HealthServiceService {
	return &service{
		repo: healthServiceRepo,
	}
}

// CreateHealthService is the service for creating a health service and saving it in the database
func (s *service) CreateHealthService(c *gin.Context, createReq *entity.RequestCreateHealthService) (string, int, error) {
	// Create a new health service
	healthService := &entity.HealthService{
		Name: createReq.Name,
	}

	// Save the health service to the database
	err := s.repo.Create(healthService)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("error creating health service: %s", err)
	}

	// Return the name of the created health service and the HTTP OK status code if the create operation is successful
	return healthService.Name, http.StatusOK, nil
}

// GetAllHealthServices returns all health services stored in the database
func (s *service) GetAllHealthServices() ([]*entity.HealthService, error) {
	// Get all health services from the database
	var healthServices []*entity.HealthService
	if err := s.repo.Find(&healthServices); err != nil {
		return nil, err
	}

	return healthServices, nil
}
