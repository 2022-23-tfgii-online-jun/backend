package reminder

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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

	// Parse individual form-data fields
	reqCreate.Name = c.PostForm("name")
	reqCreate.Type = c.PostForm("type")
	dateStr := c.PostForm("date")
	layout := "02/01/2006"
	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid date format", err)
		return
	}
	reqCreate.Date = parsedDate
	reqCreate.Note = c.PostForm("note")

	// Parse 'notification' form-data field
	notificationStr := c.PostForm("notification")
	var notifications []entity.Notification
	if err := json.Unmarshal([]byte(notificationStr), &notifications); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input for notification", err)
		return
	}
	reqCreate.Notification = notifications

	// Parse 'task' form-data field
	taskStr := c.PostForm("task")
	var tasks []entity.Task
	if err := json.Unmarshal([]byte(taskStr), &tasks); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input for task", err)
		return
	}
	reqCreate.Task = tasks

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

// GetAllReminders handler for getting all reminders
func (r *reminderHandler) GetAllReminders(c *gin.Context) {
	// Get user uuid from context
	userUUID, _ := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))

	// Fetch all the reminders from the service.
	reminders, err := r.reminderService.GetAllReminders(c, userUUID)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while fetching the reminders", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Reminders fetched successfully",
		"data":    reminders,
	})
}

// UpdateReminder handler for updating a reminder
func (r *reminderHandler) UpdateReminder(c *gin.Context) {
	// Parse the reminder UUID from the URL parameter.
	reminderUUID, err := uuid.Parse(c.Query("uuid"))
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid UUID format", err)
		return
	}

	// Bind the incoming JSON payload to an UpdateReminder struct.
	reqUpdate := &entity.RequestUpdateReminder{}
	// Parse individual form-data fields
	reqUpdate.Name = c.PostForm("name")
	reqUpdate.Type = c.PostForm("type")
	dateStr := c.PostForm("date")
	layout := "02/01/2006"
	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid date format", err)
		return
	}
	reqUpdate.Date = parsedDate
	reqUpdate.Note = c.PostForm("note")

	// Parse 'notification' form-data field
	notificationStr := c.PostForm("notification")
	var notifications []entity.Notification
	if err := json.Unmarshal([]byte(notificationStr), &notifications); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input for notification", err)
		return
	}
	reqUpdate.Notification = notifications

	// Parse 'task' form-data field
	taskStr := c.PostForm("task")
	var tasks []entity.Task
	if err := json.Unmarshal([]byte(taskStr), &tasks); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input for task", err)
		return
	}
	reqUpdate.Task = tasks

	// Update the reminder in the database.
	updatedReminder, err := r.reminderService.UpdateReminder(c, reminderUUID, reqUpdate)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while updating the reminder", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Reminder updated successfully",
		"data":    updatedReminder,
	})
}
