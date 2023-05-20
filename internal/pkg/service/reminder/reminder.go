package reminder

import (
	"fmt"
	"net/http"
	"path"

	"github.com/emur-uy/backend/config"
	aws "github.com/emur-uy/backend/internal/infra/repositories/spaces"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type service struct {
	repo                 ports.ReminderRepository
	mediaService         ports.MediaService
	reminderMediaService ports.ReminderMediaService
}

// NewService returns a new instance of the reminder service with the given reminder repository.
func NewService(reminderRepo ports.ReminderRepository, mediaService ports.MediaService,
	reminderMediaService ports.ReminderMediaService) ports.ReminderService {
	return &service{
		repo:                 reminderRepo,
		mediaService:         mediaService,
		reminderMediaService: reminderMediaService,
	}
}

// CreateReminder is the service for creating a reminder and saving it in the database
func (s *service) CreateReminder(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateReminder) (int, error) {
	user := &entity.User{}

	// Find user by ID
	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		// Return error if the user is not found
		return http.StatusNotFound, err
	}
	// Perform type assertion to convert foundUser to *entity.User
	user, ok := foundUser.(*entity.User)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("type assertion failed")
	}

	fileProcessCode, fileUrls, err := processUploadRequestFiles(s, c) // This now processes multiple files
	if err != nil || fileProcessCode != http.StatusOK {
		return http.StatusInternalServerError, fmt.Errorf("error processing content upload file, %s", err)
	}

	// Create a new reminder
	reminder := &entity.Reminder{
		UserID:       user.ID,
		Name:         createReq.Name,
		Type:         createReq.Type,
		Date:         createReq.Date,
		Note:         createReq.Note,
		Notification: createReq.Notification,
		Task:         createReq.Task,
		IsActive:     true,
	}

	// Save the reminder to the database
	err = s.repo.CreateWithOmit("uuid", reminder)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error creating reminder: %s", err)
	}

	// For each uploaded file, create a new media entry and a new reminder_media association
	for _, fileUrl := range fileUrls {
		media := &entity.Media{
			MediaURL: fileUrl,
		}
		err = s.mediaService.CreateMedia(media)
		if err != nil {
			return http.StatusInternalServerError, fmt.Errorf("error creating media: %s", err)
		}
		reminderMedia := &entity.ReminderMedia{
			ReminderID: reminder.ID,
			MediaID:    media.ID,
		}
		err = s.reminderMediaService.CreateReminderMedia(reminderMedia)
		if err != nil {
			return http.StatusInternalServerError, fmt.Errorf("error creating reminder_media association: %s", err)
		}
	}

	// Return the HTTP OK status code if the update is successful
	return http.StatusOK, nil
}

// processUploadRequestFiles processes the file upload request
func processUploadRequestFiles(s *service, c *gin.Context) (int, []string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return http.StatusBadRequest, nil, fmt.Errorf("get form err: %s", err.Error())
	}
	files := form.File["file"]
	if files == nil {
		return http.StatusBadRequest, nil, fmt.Errorf("file not found")
	}

	var fileUrls []string

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return http.StatusInternalServerError, nil, fmt.Errorf("failed to open file, %s", err)
		}
		defer src.Close()

		fileType := file.Header.Get("Content-Type")
		if fileType != "image/png" && fileType != "image/jpeg" {
			return http.StatusBadRequest, nil, fmt.Errorf("unsupported file type, %s", fileType)
		}

		fileExt := path.Ext(file.Filename)
		fileNameUuid := uuid.New()

		uploadPath := fmt.Sprintf("%s/%s", config.Get().AwsFolderName, fmt.Sprintf("%s%s", fileNameUuid.String(), fileExt))
		url, err := aws.UploadFileToS3Stream(src, uploadPath, true)
		if err != nil || url == "" {
			return http.StatusInternalServerError, nil, fmt.Errorf("s3 upload error: %s", err.Error())
		}

		fileUrls = append(fileUrls, url)
	}

	return http.StatusOK, fileUrls, nil
}
