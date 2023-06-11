package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MedicalRecordRepository es la interfaz que define los métodos para acceder al repositorio de registros médicos.
type MedicalRecordRepository interface {
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)
	CreateWithOmit(omit string, value interface{}) error
	Update(value interface{}) error
	// First retrieves the first record that matches the given conditions from the database
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error
}

// MedicalRecordService es la interfaz que define los métodos para administrar los registros médicos en la aplicación.
type MedicalRecordService interface {
	CreateMedicalRecord(c *gin.Context, userUUID uuid.UUID, createReq *entity.MedicalRecord) (*entity.MedicalRecord, int, error)
	GetMedicalRecord(c *gin.Context, userUUID uuid.UUID) (*entity.MedicalRecord, int, error)
	UpdateMedicalRecord(c *gin.Context, userUUID uuid.UUID, medicalRecordUUID uuid.UUID, updateReq *entity.MedicalRecord) (*entity.MedicalRecord, int, error)
}
