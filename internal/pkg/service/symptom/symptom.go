package symptom

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
	ErrTypeAssertionFailed     = errors.New("type assertion failed")
	ErrCreatingSymptom         = errors.New("error creating symptom")
	ErrFindingSymptom          = errors.New("error finding symptom")
	ErrAddingUserToSymptom     = errors.New("error adding user to symptom")
	ErrRemovingUserFromSymptom = errors.New("error removing user from symptom")
	ErrFindingSymptomUser      = errors.New("error finding symptom user")
)

// service struct holds the necessary dependencies for the symptom service
type service struct {
	repo ports.SymptomRepository
}

// NewService returns a new instance of the symptom service with the given symptom repository.
func NewService(symptomRepo ports.SymptomRepository) ports.SymptomService {
	return &service{
		repo: symptomRepo,
	}
}

// CreateSymptom is the service for creating a symptom and saving it in the database.
func (s *service) CreateSymptom(c *gin.Context, createReq *entity.RequestCreateSymptom) (*entity.Symptom, int, error) {

	if createReq == nil {
		return nil, http.StatusBadRequest, errors.New("nil payload")
	}

	// Create a new symptom
	symptom := &entity.Symptom{
		Name:     createReq.Name,
		IsActive: createReq.IsActive,
		Scale:    createReq.Scale,
	}

	// Save the symptom to the database
	err := s.repo.CreateWithOmit("uuid", symptom)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("%w: %s", ErrCreatingSymptom, err)
	}

	// Return the symptom and the HTTP OK status code if the create operation is successful
	return symptom, http.StatusOK, nil
}

// GetAllSymptoms returns all symptoms stored in the database.
func (s *service) GetAllSymptoms() ([]*entity.Symptom, error) {
	// Get all symptoms from the database
	var symptoms []*entity.Symptom
	if err := s.repo.Find(&symptoms); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrFindingSymptom, err)
	}

	return symptoms, nil
}

// AddUserToSymptom adds a user to a symptom.
func (s *service) AddUserToSymptom(userUUID uuid.UUID, symptomUser *entity.RequestCreateSymptomUser) (int, error) {
	// Find user by UUID
	user, err := s.repo.FindByUUID(userUUID, &entity.User{})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("%w: %s", ErrFindingSymptom, err)
	}

	// Ensure the found entity is of type *entity.User
	userEntity, ok := user.(*entity.User)
	if !ok {
		return http.StatusInternalServerError, ErrTypeAssertionFailed
	}

	// Find symptom by UUID
	symptom, err := s.repo.FindByUUID(symptomUser.SymptomUUID, &entity.Symptom{})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("%w: %s", ErrFindingSymptom, err)
	}

	// Ensure the found entity is of type *entity.Symptom
	symptomEntity, ok := symptom.(*entity.Symptom)
	if !ok {
		return http.StatusInternalServerError, ErrTypeAssertionFailed
	}

	// Create a new record for the user and symptom
	record := &entity.SymptomUser{
		UserID:    userEntity.ID,
		SymptomID: symptomEntity.ID,
	}

	// Add the user to the symptom
	err = s.repo.Create(record)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("%w: %s", ErrAddingUserToSymptom, err)
	}

	// Return the HTTP OK status code if the operation is successful
	return http.StatusOK, nil
}

// GetSymptomsByUser returns all symptoms related to a user.
func (s *service) GetSymptomsByUser(userUUID uuid.UUID) ([]*entity.Symptom, error) {
	// Find user by UUID
	user, err := s.repo.FindByUUID(userUUID, &entity.User{})
	if err != nil {
		return nil, fmt.Errorf("error finding user: %s", err)
	}

	// Ensure the found entity is of type *entity.User
	userEntity, ok := user.(*entity.User)
	if !ok {
		return nil, fmt.Errorf("error asserting user entity type")
	}

	// Get the symptom user IDs from the repository
	var symptomUserIDs []*entity.SymptomUser
	err = s.repo.Find(&symptomUserIDs, "user_id = ?", userEntity.ID)
	if err != nil {
		return nil, fmt.Errorf("error finding symptom user IDs: %s", err)
	}

	// Get the symptoms based on the obtained IDs
	symptoms := make([]*entity.Symptom, 0, len(symptomUserIDs))
	for _, symptomUser := range symptomUserIDs {
		var symptom entity.Symptom
		err := s.repo.Find(&symptom, "id = ?", symptomUser.SymptomID)
		if err != nil {
			return nil, fmt.Errorf("error finding symptom by ID: %s", err)
		}
		symptoms = append(symptoms, &symptom)
	}

	return symptoms, nil
}

// RemoveUserFromSymptom removes a user from a symptom.
func (s *service) RemoveUserFromSymptom(userUUID uuid.UUID, req *entity.RequestCreateSymptomUser) (int, error) {
	// Find user by UUID
	user, err := s.repo.FindByUUID(userUUID, &entity.User{})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error finding user: %s", err)
	}

	// Ensure the found entity is of type *entity.User
	userEntity, ok := user.(*entity.User)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("error asserting user entity type")
	}

	// Find symptom by UUID
	symptom, err := s.repo.FindByUUID(req.SymptomUUID, &entity.Symptom{})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error finding symptom: %s", err)
	}

	// Ensure the found entity is of type *entity.Symptom
	symptomEntity, ok := symptom.(*entity.Symptom)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("error asserting symptom entity type")
	}

	// Find symptom user by user and symptom IDs
	var symptomUser entity.SymptomUser
	err = s.repo.Find(&symptomUser, "user_id = ? AND symptom_id = ?", userEntity.ID, symptomEntity.ID)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error finding SymptomUser: %s", err)
	}

	// Remove the user from the symptom
	err = s.repo.Delete(&symptomUser)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error removing user from symptom: %s", err)
	}

	// Return the HTTP OK status code if the operation is successful
	return http.StatusOK, nil
}
