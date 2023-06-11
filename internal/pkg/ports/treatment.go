package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TreatmentRepository is the interface that defines the methods for accessing the treatment data store.
type TreatmentRepository interface {
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)
	FindItemByIDs(firstID, secondID int, tableName string, column1Name string, column2Name string, dest interface{}) error
	Create(value interface{}) error
	CreateWithOmit(omit string, value interface{}) error
	Update(value interface{}) error
	First(out interface{}, conditions ...interface{}) error
	Find(out interface{}, conditions ...interface{}) error
	Delete(out interface{}) error
}

// TreatmentService is the interface that defines the methods for managing treatments in the application.
type TreatmentService interface {
	CreateTreatment(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateTreatment) (*entity.Treatment, int, error)
	UpdateTreatment(treatmentUUID uuid.UUID, updateReq *entity.RequestUpdateTreatment) (int, error)
	DeleteTreatment(treatmentUUID uuid.UUID) (int, error)
	GetAllTreatments(userUUID uuid.UUID) ([]*entity.Treatment, error)
}
