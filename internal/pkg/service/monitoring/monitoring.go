package monitoring

import (
	"fmt"
	"net/http"
	"time"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type service struct {
	repo ports.MonitoringRepository
}

// NewService returns a new instance of the monitoring service with the given monitoring repository.
func NewService(repo ports.MonitoringRepository) ports.MonitoringService {
	return &service{
		repo: repo,
	}
}

// CreateMonitoring is the service for creating a monitoring record and saving it in the database.
func (s *service) CreateMonitoring(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateMonitoring) (*entity.Monitoring, int, error) {
	// Find user by UUID
	user, err := s.repo.FindByUUID(userUUID, &entity.User{})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error finding user: %s", err)
	}

	// Ensure the found entity is of type *entity.User
	userEntity, ok := user.(*entity.User)
	if !ok {
		return nil, http.StatusInternalServerError, fmt.Errorf("error asserting user entity type")
	}

	// Find symptom by UUID
	symptom, err := s.repo.FindByUUID(createReq.SymptomUUID, &entity.Symptom{})
	if err != nil {
		if err.Error() == "record not found" {
			return nil, http.StatusBadRequest, fmt.Errorf("symptom not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("error finding symptom: %s", err)
	}

	// Ensure the found entity is of type *entity.Symptom
	symptomEntity, ok := symptom.(*entity.Symptom)
	if !ok {
		return nil, http.StatusInternalServerError, fmt.Errorf("error asserting symptom entity type")
	}

	// Check if the user already has a monitoring record for the same symptom on the current day
	var existingMonitoring entity.Monitoring
	err = s.repo.FindItemByIDs(userEntity.ID, symptomEntity.ID, "monitorings", "user_id", "symptom_id", &existingMonitoring)
	if err == nil {
		return nil, http.StatusBadRequest, fmt.Errorf("a monitoring record for the same symptom already exists for the current day")
	} else if err != nil && err.Error() != "record not found" {
		return nil, http.StatusInternalServerError, fmt.Errorf("error checking monitoring record: %s", err)
	}

	// Check if the scale value in createReq is equal or smaller than the scale value in symptomEntity
	if createReq.Scale > symptomEntity.Scale {
		return nil, http.StatusBadRequest, fmt.Errorf("the scale value in the request is greater than the allowed scale for the symptom")
	}

	// Create a new monitoring record
	monitoring := &entity.Monitoring{
		UserID:    userEntity.ID,
		SymptomID: symptomEntity.ID,
		Scale:     createReq.Scale,
		Date:      time.Now(),
	}

	// Save the monitoring record to the database
	err = s.repo.Create(monitoring)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error creating monitoring: %s", err)
	}

	// Return the monitoring record and the HTTP OK status code if the create operation is successful
	return monitoring, http.StatusOK, nil
}

// GetAllMonitorings retrieves all monitoring records for a user from the database.
func (s *service) GetAllMonitorings(c *gin.Context, userUUID uuid.UUID) ([]*entity.Monitoring, int, error) {
	// Find user by UUID
	user := &entity.User{}
	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error finding user: %s", err)
	}

	// Perform type assertion to convert foundUser to *entity.User
	userEntity, ok := foundUser.(*entity.User)
	if !ok {
		return nil, http.StatusInternalServerError, fmt.Errorf("error asserting user entity type")
	}

	// Get all reminders for this user
	var monitorings []*entity.Monitoring
	if err := s.repo.Find(&monitorings, "user_id = ?", userEntity.ID); err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error retrieving monitorings: %s", err)
	}

	// Return the monitoring records and the HTTP OK status code if the retrieval is successful
	return monitorings, http.StatusOK, nil
}
