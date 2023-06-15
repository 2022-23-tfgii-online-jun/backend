package reminder

import (
	"github.com/emur-uy/backend/internal/infra/api/middlewares"
	"github.com/emur-uy/backend/internal/infra/api/middlewares/constants"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/media"
	"github.com/emur-uy/backend/internal/pkg/service/reminder"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the reminder-related routes on the given gin.Engine instance.
// It initializes the necessary components, such as the repository, service, and handler,
// to handle reminder-related operations in a hexagonal architecture.
func RegisterRoutes(e *gin.Engine) {
	// Initialize the repository by creating a new PostgreSQL client.
	client := postgresql.NewClient()

	// Create new repository instances for each repository interface
	reminderRepo := postgresql.NewReminderRepository(client)
	mediaRepo := postgresql.NewMediaRepository(client)
	reminderMediaRepo := postgresql.NewReminderMediaRepository(client)

	// Create new services
	mediaService := media.NewService(mediaRepo)
	reminderMediaService := reminder.NewReminderMediaService(reminderMediaRepo)
	reminderService := reminder.NewService(reminderRepo, mediaService, reminderMediaService)

	// Create a new reminderHandler instance by injecting the ReminderService.
	handler := newHandler(reminderService)

	// Group the reminder routes together.
	reminderRoutes := e.Group("/api/v1/reminders")
	reminderRoutes.Use(middlewares.Authenticate(), middlewares.Authorize(constants.RoleUser))

	// Register user routes requiring authentication and authorization for user role.
	reminderRoutes.POST("", handler.CreateReminder)
	reminderRoutes.GET("", handler.GetAllReminders)
	reminderRoutes.PUT("", handler.UpdateReminder)
	reminderRoutes.DELETE("", handler.DeleteReminder)
}
