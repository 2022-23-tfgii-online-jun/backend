package treatment

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	ErrTypeAssertionFailed = errors.New("type assertion failed")
	ErrCreatingTreatment   = errors.New("error creating treatment")
	ErrFindingUser         = errors.New("error finding user")
	ErrUpdatingTreatment   = errors.New("error updating treatment")
	ErrDeletingTreatment   = errors.New("error deleting treatment")
)

// service struct holds the necessary dependencies for the treatment service
type service struct {
	repo ports.TreatmentRepository
}

// NewService returns a new instance of the treatment service with the given treatment repository.
func NewService(treatmentRepo ports.TreatmentRepository) ports.TreatmentService {
	return &service{
		repo: treatmentRepo,
	}
}

// CreateTreatment is the service for creating a treatment and saving it in the database.
func (s *service) CreateTreatment(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateTreatment) (*entity.Treatment, int, error) {
	// Find user by UUID
	user := &entity.User{}
	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	// Perform type assertion to convert foundUser to *entity.User
	user, ok := foundUser.(*entity.User)
	if !ok {
		return nil, http.StatusInternalServerError, ErrTypeAssertionFailed
	}

	// Create a new treatment
	treatment := &entity.Treatment{
		Name:      createReq.Name,
		Type:      createReq.Type,
		Frequency: createReq.Frequency,
		Shots:     createReq.Shots,
		DateStart: createReq.DateStart,
		Notes:     createReq.Notes,
		UserID:    user.ID,
	}

	// Save the treatment to the database
	err = s.repo.CreateWithOmit("uuid", treatment)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error creating treatment: %s", err)
	}

	return treatment, http.StatusOK, nil
}

// GetAllTreatments returns all treatments of a specific user stored in the database.
func (s *service) GetAllTreatments(userUUID uuid.UUID) ([]*entity.Treatment, error) {
	// Find user by UUID
	user := &entity.User{}
	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		return nil, err
	}

	// Perform type assertion to convert foundUser to *entity.User
	user, ok := foundUser.(*entity.User)
	if !ok {
		return nil, ErrTypeAssertionFailed
	}

	// Get all treatments from the database for the user
	var treatments []*entity.Treatment
	if err := s.repo.Find(&treatments, "user_id = ?", user.ID); err != nil {
		return nil, err
	}

	return treatments, nil
}

// UpdateTreatment is the service for updating a treatment in the database.
func (s *service) UpdateTreatment(treatmentUUID uuid.UUID, updateReq *entity.RequestUpdateTreatment) (int, error) {
	// Find the existing treatment by UUID
	treatment := &entity.Treatment{}
	foundTreatment, err := s.repo.FindByUUID(treatmentUUID, treatment)
	if err != nil {
		// Return error if the treatment is not found
		return http.StatusNotFound, err
	}
	// Perform type assertion to convert foundTreatment to *entity.Treatment
	treatment, ok := foundTreatment.(*entity.Treatment)
	if !ok {
		return http.StatusInternalServerError, ErrTypeAssertionFailed
	}

	// Update the treatment fields with the new data from the update request
	treatment.Name = updateReq.Name

	// Update the treatment in the database
	err = s.repo.Update(treatment)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error updating treatment: %s", err)
	}

	// Return the HTTP OK status code if the update is successful
	return http.StatusOK, nil
}

// DeleteTreatment is the service for deleting a treatment from the database.
func (s *service) DeleteTreatment(treatmentUUID uuid.UUID) (int, error) {
	// Find the existing treatment by UUID
	treatment := &entity.Treatment{}
	foundTreatment, err := s.repo.FindByUUID(treatmentUUID, treatment)
	if err != nil {
		// Return error if the treatment is not found
		return http.StatusNotFound, err
	}
	// Perform type assertion to convert foundTreatment to *entity.Treatment
	treatment, ok := foundTreatment.(*entity.Treatment)
	if !ok {
		return http.StatusInternalServerError, ErrTypeAssertionFailed
	}

	// Delete the treatment from the database
	err = s.repo.Delete(treatment)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error deleting treatment: %s", err)
	}

	// Return the HTTP OK status code if the delete is successful
	return http.StatusOK, nil
}
