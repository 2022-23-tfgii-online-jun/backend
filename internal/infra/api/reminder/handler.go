package reminder

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type reminderHandler struct {
	reminderService ports.ReminderService
}

func newHandler(reminderService ports.ReminderService) *reminderHandler {
	return &reminderHandler{
		reminderService: reminderService,
	}
}

// CreateReminder handler for creating a reminder
func (r *reminderHandler) CreateReminder(c *gin.Context) {
	reqCreate := &entity.RequestCreateReminder{}

	// Get user uuid from context
	userUUID, _ := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))

	// Bind incoming JSON payload to the reqCreate struct.
	if err := c.ShouldBindJSON(reqCreate); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Create the reminder and store it in the database.
	createdReminder, err := r.reminderService.CreateReminder(c, userUUID, reqCreate)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while creating the reminder", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Reminder created successfully",
		"data":    createdReminder,
	})
}

func handleError(c *gin.Context, status int, msg string, err error) {
	log.Println(err)
	c.JSON(status, gin.H{
		"code":    status,
		"message": msg,
		"error":   err.Error(),
	})
}
