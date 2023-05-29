package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ReminderRepository is the interface that defines the methods for accessing the reminder data store.
type ReminderRepository interface {
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	CreateWithOmit(omit string, value interface{}) error

	Find(model interface{}, dest interface{}, conditions ...interface{}) error
	Update(value interface{}) error

	// 	First(out interface{}, conditions ...interface{}) error

	// 	Find(out interface{}, conditions ...interface{}) error

	Delete(out interface{}) error
	//	}
}

// ReminderService is the interface that defines the methods for managing reminders in the application.
type ReminderService interface {
	CreateReminder(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateReminder) (int, error)
	GetAllReminders(c *gin.Context, userUUID uuid.UUID) ([]*entity.GetReminderResponse, error)
	UpdateReminder(c *gin.Context, reminderUUID uuid.UUID, updateReq *entity.RequestUpdateReminder) (int, error)
	// DeleteReminder(c *gin.Context, reminderUUID uuid.UUID) (int, error)
	// GetAllReminders() ([]*entity.Reminder, error)
}
