package symptom

import (
	"fmt"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type service struct {
	repo ports.SymptomRepository
}

// NewService returns a new instance of the symptom service with the given symptom repository.
func NewService(symptomRepo ports.SymptomRepository) ports.SymptomService {
	return &service{
		repo: symptomRepo,
	}
}

// CreateSymptom is the service for creating a symptom and saving it in the database
func (s *service) CreateSymptom(c *gin.Context, createReq *entity.RequestCreateSymptom) (*entity.Symptom, int, error) {
	// Create a new symptom
	symptom := &entity.Symptom{
		Name:     createReq.Name,
		IsActive: createReq.IsActive,
		Scale:    createReq.Scale,
	}

	// Save the symptom to the database
	err := s.repo.CreateWithOmit("uuid", symptom)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error creating symptom: %s", err)
	}

	// Return the symptom and the HTTP OK status code if the create operation is successful
	return symptom, http.StatusOK, nil
}

// GetAllSymptoms returns all symptoms stored in the database
func (s *service) GetAllSymptoms() ([]*entity.Symptom, error) {
	// Get all symptoms from the database
	var symptoms []*entity.Symptom
	if err := s.repo.Find(&symptoms); err != nil {
		return nil, err
	}

	return symptoms, nil
}
