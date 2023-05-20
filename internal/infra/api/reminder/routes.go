package reminder

import (
	"github.com/emur-uy/backend/internal/infra/api/middlewares"
	"github.com/emur-uy/backend/internal/infra/api/middlewares/constants"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/media"
	"github.com/emur-uy/backend/internal/pkg/service/reminder"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	// Initialize the repository by creating a new PostgreSQL client.
	client := postgresql.NewClient()

	// Create new repository instances for each repository interface
	reminderRepo := postgresql.NewReminderRepository(client)
	mediaRepo := postgresql.NewMediaRepository(client)
	reminderMediaRepo := postgresql.NewReminderMediaRepository(client)

	// Create new services
	mediaService := media.NewService(mediaRepo)
	reminderMediaService := reminder.NewReminderMediaService(reminderMediaRepo)              // Corrección aquí
	reminderService := reminder.NewService(reminderRepo, mediaService, reminderMediaService) // Corrección aquí

	//	reminderService := reminder.NewService(reminderRepo, mediaService, reminderMediaService)

	// Create a new reminderHandler instance by injecting the ReminderService.
	handler := newHandler(reminderService)

	// Group the reminder routes together.
	reminderRoutes := e.Group("/api/v1/reminders")

	// Register route for creating reminders accessible only to user role.
	allowedRolesCreate := []string{constants.RoleUser}
	reminderRoutes.POST("", middlewares.Authenticate(), middlewares.Authorize(allowedRolesCreate...), handler.CreateReminder)
}
