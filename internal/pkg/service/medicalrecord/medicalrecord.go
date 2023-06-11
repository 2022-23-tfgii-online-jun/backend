package medicalrecord

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type medicalRecordService struct {
	repo ports.MedicalRecordRepository
}

// NewService returns a new instance of the medical record service with the given medical record repository.
func NewService(repo ports.MedicalRecordRepository) ports.MedicalRecordService {
	return &medicalRecordService{
		repo: repo,
	}
}

// CreateMedicalRecord is the service for creating a medical record and saving it in the database.
func (s *medicalRecordService) CreateMedicalRecord(c *gin.Context, userUUID uuid.UUID, createReq *entity.MedicalRecord) (*entity.MedicalRecord, int, error) {
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

	// Create a new medical record entity
	medicalRecord := &entity.MedicalRecord{
		UserID:                  userEntity.ID,
		HealthCareProvider:      createReq.HealthCareProvider,
		EmergencyMedicalService: createReq.EmergencyMedicalService,
		MultipleSclerosisType:   createReq.MultipleSclerosisType,
		LaboralCondition:        createReq.LaboralCondition,
		Conmorbidity:            createReq.Conmorbidity,
		TreatingNeurologist:     createReq.TreatingNeurologist,
		SupportNetwork:          createReq.SupportNetwork,
		IsDisabled:              createReq.IsDisabled,
		EducationalLevel:        createReq.EducationalLevel,
	}

	// Save the medical record to the database
	err = s.repo.CreateWithOmit("uuid", medicalRecord)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error creating medical record: %s", err)
	}

	// Return the created medical record and the HTTP OK status code
	return medicalRecord, http.StatusOK, nil
}

// GetMedicalRecord is the service for retrieving a medical record from the database.
func (s *medicalRecordService) GetMedicalRecord(c *gin.Context, uuid uuid.UUID) (*entity.MedicalRecord, int, error) {
	// Retrieve the user ID based on the provided UUID
	user, err := s.repo.FindByUUID(uuid, &entity.User{})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error finding user: %s", err)
	}

	// Ensure the found entity is of type *entity.User
	userEntity, ok := user.(*entity.User)
	if !ok {
		return nil, http.StatusInternalServerError, fmt.Errorf("error asserting user entity type")
	}

	// Retrieve the medical record for the user from the database
	medicalRecord := &entity.MedicalRecord{}
	err = s.repo.First(medicalRecord, "user_id = ?", userEntity.ID)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error retrieving medical record: %s", err)
	}

	// Return the retrieved medical record and the HTTP OK status code
	return medicalRecord, http.StatusOK, nil
}

// UpdateMedicalRecord is the service for updating a medical record in the database.
func (s *medicalRecordService) UpdateMedicalRecord(c *gin.Context, userUUID uuid.UUID, medicalRecordUUID uuid.UUID, updateReq *entity.MedicalRecord) (*entity.MedicalRecord, int, error) {
	// Find the user by UUID
	user, err := s.repo.FindByUUID(userUUID, &entity.User{})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error finding user: %s", err)
	}

	// Ensure the found entity is of type *entity.User
	userEntity, ok := user.(*entity.User)
	if !ok {
		return nil, http.StatusInternalServerError, fmt.Errorf("error asserting user entity type")
	}

	// Find the medical record by UUID
	medicalRecord, err := s.repo.FindByUUID(medicalRecordUUID, &entity.MedicalRecord{})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error finding medical record: %s", err)
	}

	// Ensure the found entity is of type *entity.MedicalRecord
	medicalRecordEntity, ok := medicalRecord.(*entity.MedicalRecord)
	if !ok {
		return nil, http.StatusInternalServerError, fmt.Errorf("error asserting medical record entity type")
	}

	// Check if the user is the owner of the medical record
	if medicalRecordEntity.UserID != userEntity.ID {
		return nil, http.StatusForbidden, fmt.Errorf("user is not authorized to update the medical record")
	}

	// Update the medical record entity with the new values
	medicalRecordEntity.HealthCareProvider = updateReq.HealthCareProvider
	medicalRecordEntity.EmergencyMedicalService = updateReq.EmergencyMedicalService
	medicalRecordEntity.MultipleSclerosisType = updateReq.MultipleSclerosisType
	medicalRecordEntity.LaboralCondition = updateReq.LaboralCondition
	medicalRecordEntity.Conmorbidity = updateReq.Conmorbidity
	medicalRecordEntity.TreatingNeurologist = updateReq.TreatingNeurologist
	medicalRecordEntity.SupportNetwork = updateReq.SupportNetwork
	medicalRecordEntity.IsDisabled = updateReq.IsDisabled
	medicalRecordEntity.EducationalLevel = updateReq.EducationalLevel

	// Save the updated medical record to the database
	err = s.repo.Update(medicalRecordEntity)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error updating medical record: %s", err)
	}

	// Return the updated medical record and the HTTP OK status code
	return medicalRecordEntity, http.StatusOK, nil
}

// handleError handles errors by sending an appropriate response to the client.
func handleError(c *gin.Context, status int, message string, err error) {
	log.Printf("[MedicalRecordService]: %s, %v", message, err)
	c.JSON(status, gin.H{
		"code":    status,
		"message": message,
	})
}
