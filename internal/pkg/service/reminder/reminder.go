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

// GetAllReminders retrieves all reminders from the database
func (s *service) GetAllReminders(c *gin.Context, userUUID uuid.UUID) ([]*entity.GetReminderResponse, error) {
	user := &entity.User{}

	// Find user by UUID
	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		// Return error if the user is not found
		return nil, err
	}
	// Perform type assertion to convert foundUser to *entity.User
	user, ok := foundUser.(*entity.User)
	if !ok {
		return nil, fmt.Errorf("type assertion failed")
	}

	reminders := []*entity.Reminder{}
	// Get all reminders for this user
	//err = s.repo.Find(user.ID, reminder)
	err = s.repo.Find(&entity.Reminder{}, &reminders, "user_id = ?", user.ID)

	response := []*entity.GetReminderResponse{}

	// Get media for each reminder and prepare response
	for _, reminder := range reminders {

		getReminderResponse := &entity.GetReminderResponse{
			UUID:         reminder.UUID,
			Name:         reminder.Name,
			Type:         reminder.Type,
			Date:         reminder.Date,
			Notification: reminder.Notification,
			Task:         reminder.Task,
			Note:         reminder.Note,
			IsActive:     reminder.IsActive,
		}

		//Get reminder medias
		reminderMedias := []*entity.ReminderMedia{}
		err = s.repo.Find(&entity.ReminderMedia{}, &reminderMedias, "reminder_id = ?", reminder.ID)
		if err != nil {
			// Return error if the user is not found
			return nil, err
		}

		reminderMediaResponses := []entity.GetReminderMediaResponse{}

		//Get media details
		for _, reminderMedia := range reminderMedias {
			//Get media details
			media := entity.Media{}
			err = s.repo.Find(&entity.Media{}, &media, "id = ?", reminderMedia.MediaID)
			if err != nil {
				// Return error if the user is not found
				return nil, err
			}

			//add details in response
			reminderMediaResponse := &entity.GetReminderMediaResponse{
				MediaURL:   media.MediaURL,
				MediaThumb: media.MediaThumb,
			}
			reminderMediaResponses = append(reminderMediaResponses, *reminderMediaResponse)
		}
		getReminderResponse.Media = reminderMediaResponses
		response = append(response, getReminderResponse)
	}

	if err != nil {
		// Return error if the user is not found
		return nil, err
	}

	return response, nil
}

// UpdateReminder is the service for updating a reminder in the database
func (s *service) UpdateReminder(c *gin.Context, reminderUUID uuid.UUID, updateReq *entity.RequestUpdateReminder) (int, error) {
	// Find the existing reminder by UUID
	reminder := &entity.Reminder{}
	foundReminder, err := s.repo.FindByUUID(reminderUUID, reminder)
	if err != nil {
		// Return error if the article is not found
		return http.StatusNotFound, err
	}
	// Perform type assertion to convert foundReminder to *entity.Reminder
	reminder, ok := foundReminder.(*entity.Reminder)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("type assertion failed")
	}

	// Update the reminder fields with the new data from the update request
	reminder.Name = updateReq.Name
	reminder.Type = updateReq.Type
	reminder.Date = updateReq.Date
	reminder.Notification = updateReq.Notification
	reminder.Task = updateReq.Task
	reminder.Note = updateReq.Note

	// Update the reminder in the database
	err = s.repo.Update(reminder)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error updating reminder: %s", err)
	}

	fileProcessCode, fileUrls, err := processUploadRequestFiles(s, c) // This now processes multiple files
	if err != nil || fileProcessCode != http.StatusOK {
		return http.StatusInternalServerError, fmt.Errorf("error processing content upload file, %s", err)
	}

	//get existing reminder media data
	reminderMedias := []*entity.ReminderMedia{}
	err = s.repo.Find(&entity.ReminderMedia{}, &reminderMedias, "reminder_id = ?", reminder.ID)
	if err != nil {
		// Return error if the media is not found
		return http.StatusInternalServerError, fmt.Errorf("error updating reminder media, %s", err)
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

	//delete old media entries
	for _, reminderMedia := range reminderMedias {

		mediaID := reminderMedia.MediaID

		//delete the reminderMedia entry
		err = s.repo.Delete(reminderMedia)
		if err != nil {
			// Return an error response if there was an issue deleting the reminder_media.
			return http.StatusInternalServerError, fmt.Errorf("failed to update media")
		}

		// Delete the media from the repository.
		media := entity.Media{
			ID: mediaID,
		}
		err = s.repo.Delete(media)
		if err != nil {
			// Return an error response if there was an issue deleting the media.
			return http.StatusInternalServerError, fmt.Errorf("failed to update media")
		}
	}

	// Return the HTTP OK status code if the update is successful
	return http.StatusOK, nil
}

func (s *service) DeleteReminder(c *gin.Context, reminderUUID uuid.UUID) error {
	reminder := &entity.Reminder{}

	// 1. Find reminder by UUID
	foundReminder, err := s.repo.FindByUUID(reminderUUID, reminder)
	if err != nil {
		// Return error if the reminder is not found
		return err
	}
	// Perform type assertion to convert foundReminder to *entity.Reminder
	reminder, ok := foundReminder.(*entity.Reminder)
	if !ok {
		return fmt.Errorf("type assertion failed")
	}

	// 2. Find reminder_media associations by reminder ID
	reminderMedias := []*entity.ReminderMedia{}
	err = s.reminderMediaService.FindByReminderID(reminder.ID, &reminderMedias)
	if err != nil {
		return err
	}

	// 3. Iterate over each reminder_media association
	for _, reminderMedia := range reminderMedias {
		media := &entity.Media{}
		// Find the media by ID
		err := s.mediaService.FindByMediaID(reminderMedia.MediaID, media)
		if err != nil {
			return err
		}

		// Delete file from S3
		err = aws.DeleteObjectFromS3(media.MediaURL)
		if err != nil {
			return fmt.Errorf("failed to delete uploaded files from s3, error: %s", err.Error())
		}

		// Delete reminder_media association from db
		err = s.reminderMediaService.DeleteReminderMedia(reminderMedia)
		if err != nil {
			return err
		}

		// Delete media from db
		err = s.mediaService.DeleteMedia(media)
		if err != nil {
			return err
		}
	}

	// 4. Delete reminder from db
	err = s.repo.Delete(reminder)
	if err != nil {
		return err
	}

	return nil
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
