package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MonitoringRepository is the interface that defines the methods for accessing the monitoring data store.
type MonitoringRepository interface {
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)
	Create(value interface{}) error
	FindItemByIDs(firstID, secondID int, tableName string, column1Name string, column2Name string, dest interface{}) error
}

// MonitoringService is the interface that defines the methods for managing monitorings in the application.
type MonitoringService interface {
	CreateMonitoring(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateMonitoring) (*entity.Monitoring, int, error)
}
