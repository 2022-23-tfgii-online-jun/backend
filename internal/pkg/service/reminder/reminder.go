package reminder

import (
	"fmt"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type service struct {
	repo ports.ReminderRepository
}

// NewService returns a new instance of the reminder service with the given reminder repository.
func NewService(reminderRepo ports.ReminderRepository) ports.ReminderService {
	return &service{
		repo: reminderRepo,
	}
}

// CreateReminder is the service for creating a reminder and saving it in the database
func (s *service) CreateReminder(c *gin.Context, userUUID uuid.UUID, fileUUID uuid.UUID, createReq *entity.RequestCreateReminder) (int, error) {
	user := &entity.User{}

	// Find user by UUID
	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		// Return error if the user is not found
		return http.StatusNotFound, err
	}
	user, ok := foundUser.(*entity.User)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("type assertion failed")
	}

	file := &entity.File{}

	// Find file by UUID
	foundFile, err := s.repo.FindByUUID(fileUUID, file)
	if err != nil {
		// Return error if the file is not found
		return http.StatusNotFound, err
	}
	file, ok = foundFile.(*entity.File)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("type assertion failed")
	}

	// Create a new reminder
	reminder := &entity.Reminder{
		UserID:        user.ID,
		FileID:        file.ID,
		Name:          createReq.Name,
		Type:          createReq.Type,
		Date:          createReq.Date,
		Notifications: 0,
		Tasks:         0,
		IsActive:      true,
	}

	// Save the reminder to the database
	err = s.repo.CreateWithOmit("uuid", reminder)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error creating reminder: %s", err)
	}

	// Return the HTTP OK status code if the update is successful
	return http.StatusOK, nil
}
