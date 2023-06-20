package reminder

import (
	"bytes"
	"errors"
	"fmt"
	aws "github.com/emur-uy/backend/internal/infra/repositories/spaces"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"testing"
)

var testUserUuid = uuid.MustParse("24df3f36-ca63-11ed-afa1-0242ac120002")
var testReminderUuid = uuid.MustParse("1a09e86a-4011-4290-85f3-8e2d6f7f0866")
var testReminderUuidWithoutMedias = uuid.MustParse("bfb23f5c-a664-432b-b6cc-b7cd17bacf5b")
var testUserUuidWithoutReminders = uuid.MustParse("d9cb60d9-1134-4de0-bb03-f4dbc5329496")

type MockReminderRepository struct{}

func (m *MockReminderRepository) CreateWithOmit(omit string, value interface{}) error {
	return nil
}

func (m *MockReminderRepository) Find(model interface{}, dest interface{}, conditions ...interface{}) error {
	if conditions[0] == "user_id = ?" && conditions[1] == 1 {
		return nil
	}
	return errors.New("not found")
}

func (m *MockReminderRepository) Update(value interface{}) error {
	return nil
}

func (m *MockReminderRepository) Delete(out interface{}) error {
	return nil
}

type MockMediaService struct{}

func (m MockMediaService) CreateMedia(media *entity.Media) error {
	return nil
}

func (m MockMediaService) DeleteMedia(media *entity.Media) error {
	return nil
}

func (m MockMediaService) FindByMediaID(id int, i *entity.Media) error {
	return nil
}

type MockReminderMediaService struct{}

func (m MockReminderMediaService) CreateReminderMedia(reminderMedia *entity.ReminderMedia) error {
	return nil
}

func (m MockReminderMediaService) DeleteReminderMedia(reminderMedia *entity.ReminderMedia) error {
	return nil
}

func (m MockReminderMediaService) FindByReminderID(id int, i *[]*entity.ReminderMedia) error {
	if id == 1 {
		return nil
	}
	return errors.New("not found")
}

// FindByUUID is a mock implementation of the FindByUUID method.
func (m *MockReminderRepository) FindByUUID(uId uuid.UUID, out interface{}) (interface{}, error) {
	if uId == testUserUuid {
		usr := &entity.User{
			ID:   1,
			UUID: testUserUuid,
		}
		return usr, nil
	}
	if uId == testUserUuidWithoutReminders {
		usr := &entity.User{
			ID:   2,
			UUID: testUserUuidWithoutReminders,
		}
		return usr, nil
	}
	if uId == testReminderUuid {
		res := &entity.Reminder{
			ID:   1,
			UUID: testReminderUuid,
		}
		return res, nil
	}
	if uId == testReminderUuidWithoutMedias {
		res := &entity.Reminder{
			ID:   2,
			UUID: testReminderUuidWithoutMedias,
		}
		return res, nil
	}
	return nil, errors.New("not found")
}

func mockUploadFileToS3Stream(src io.Reader, uploadPath string, isPublic bool) (string, error) {
	// Mocked implementation of the upload function
	return "mocked-url", nil
}

func TestCreateReminder(t *testing.T) {

	// Set up the mock repository and service.
	mockRepo := &MockReminderRepository{}
	mockMediaSvc := &MockMediaService{}
	mockRecipeMediaSvc := &MockReminderMediaService{}
	s := NewService(mockRepo, mockMediaSvc, mockRecipeMediaSvc)

	gin.SetMode(gin.TestMode)

	// Create a test context
	c, _ := gin.CreateTestContext(nil)

	// Create a sample image file data
	imageData := []byte("sample image data")

	// Create a byte buffer to simulate the image file
	fileBuf := &bytes.Buffer{}
	fileWriter := multipart.NewWriter(fileBuf)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name=file; filename="image.jpg"`))
	h.Set("Content-Type", "image/jpeg")
	part, err := fileWriter.CreatePart(h)

	// Write the image data to the form field
	_, err = part.Write(imageData)
	if err != nil {
		t.Fatalf("Failed to write image data: %v", err)
	}

	// Close the multipart writer
	err = fileWriter.Close()
	if err != nil {
		t.Fatalf("Failed to close multipart writer: %v", err)
	}

	// Set the request body and headers
	c.Request = httptest.NewRequest(http.MethodPost, "/", fileBuf)
	c.Request.Header.Set("Content-Type", "multipart/form-data; boundary="+fileWriter.Boundary())

	createReq := &entity.RequestCreateReminder{
		Name:    "Test",
		Medical: 1,
	}

	uploadFunc = mockUploadFileToS3Stream
	defer func() {
		// Restore the original upload function after the test
		uploadFunc = aws.UploadFileToS3Stream
	}()

	// Create a test context without request file
	ctx, _ := gin.CreateTestContext(nil)
	ctx.Request = httptest.NewRequest(http.MethodPost, "/", fileBuf)
	ctx.Request.Header.Set("Content-Type", "multipart/form-data")

	// Define test cases.
	testCases := []struct {
		name        string
		uId         uuid.UUID
		expectError bool
		context     gin.Context
	}{
		{"user exist & reminder creation successful", testUserUuid, false, *c},
		{"user doesn't exist & reminder creation failed", uuid.New(), true, *ctx},
		{"file not found & reminder creation failed", testUserUuid, true, *ctx},
	}

	// Execute test cases.
	// Iterate through each test case and run the corresponding test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			statusCode, err := s.CreateReminder(&tc.context, tc.uId, createReq)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned
				require.Error(t, err)
				assert.NotEqual(t, http.StatusOK, statusCode)
			} else {
				// If no error is expected, ensure there is no error returned
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, statusCode)
			}
		})
	}
}

func TestUpdateReminder(t *testing.T) {
	// Set up the mock repository and service.
	mockRepo := &MockReminderRepository{}
	mockMediaSvc := &MockMediaService{}
	mockRecipeMediaSvc := &MockReminderMediaService{}
	s := NewService(mockRepo, mockMediaSvc, mockRecipeMediaSvc)

	gin.SetMode(gin.TestMode)

	// Create a test context
	c, _ := gin.CreateTestContext(nil)

	// Create a sample image file data
	imageData := []byte("sample image data")

	// Create a byte buffer to simulate the image file
	fileBuf := &bytes.Buffer{}
	fileWriter := multipart.NewWriter(fileBuf)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name=file; filename="image.jpg"`))
	h.Set("Content-Type", "image/jpeg")
	part, err := fileWriter.CreatePart(h)

	// Write the image data to the form field
	_, err = part.Write(imageData)
	if err != nil {
		t.Fatalf("Failed to write image data: %v", err)
	}

	// Close the multipart writer
	err = fileWriter.Close()
	if err != nil {
		t.Fatalf("Failed to close multipart writer: %v", err)
	}

	// Set the request body and headers
	c.Request = httptest.NewRequest(http.MethodPost, "/", fileBuf)
	c.Request.Header.Set("Content-Type", "multipart/form-data; boundary="+fileWriter.Boundary())
	//c.Request.Header.Set("Content-Type", "image/jpeg")

	uploadFunc = mockUploadFileToS3Stream
	defer func() {
		// Restore the original upload function after the test
		uploadFunc = aws.UploadFileToS3Stream
	}()

	// Create a test context without request file
	ctx, _ := gin.CreateTestContext(nil)
	ctx.Request = httptest.NewRequest(http.MethodPost, "/", fileBuf)
	ctx.Request.Header.Set("Content-Type", "multipart/form-data")

	req := &entity.RequestUpdateReminder{
		Name: "Test",
		Type: "test",
	}

	// Define test cases.
	testCases := []struct {
		name        string
		uId         uuid.UUID
		expectError bool
		request     *entity.RequestUpdateReminder
		context     gin.Context
	}{
		{"reminder exist & recipe update successful", testReminderUuid, false, req, *c},
		{"reminder doesn't exist & reminder updation failed", uuid.New(), true, nil, *c},
		{"reminder updation failed, reminder medias don't exists", testReminderUuidWithoutMedias, true, req, *c},
		{"reminder updation failed, invalid request", testReminderUuid, true, nil, *ctx},
	}

	// Execute test cases.
	// Iterate through each test case and run the corresponding test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			statusCode, err := s.UpdateReminder(c, tc.uId, tc.request)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned and the statusCode is not OK.
				require.Error(t, err)
				assert.NotEqual(t, http.StatusOK, statusCode)
			} else {
				// If no error is expected, ensure there is no error returned and the statusCode is OK.
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, statusCode)
			}
		})
	}
}

func TestDeleteReminder(t *testing.T) {
	// Set up the mock repository and service.
	mockRepo := &MockReminderRepository{}
	mockMediaSvc := &MockMediaService{}
	mockRecipeMediaSvc := &MockReminderMediaService{}
	s := NewService(mockRepo, mockMediaSvc, mockRecipeMediaSvc)

	// Create a test context
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	// Define test cases.
	testCases := []struct {
		name        string
		uId         uuid.UUID
		expectError bool
	}{
		{"reminder exist & reminder deletion successful", testReminderUuid, false},
		{"reminder doesn't exist & reminder deletion failed", uuid.New(), true},
	}

	// Execute test cases.
	// Iterate through each test case and run the corresponding test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := s.DeleteReminder(c, tc.uId)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned
				require.Error(t, err)
			} else {
				// If no error is expected, ensure there is no error returned
				require.NoError(t, err)
			}
		})
	}
}

func TestGetAllReminders(t *testing.T) {
	// Set up the mock repository and service.
	mockRepo := &MockReminderRepository{}
	mockMediaSvc := &MockMediaService{}
	mockRecipeMediaSvc := &MockReminderMediaService{}
	s := NewService(mockRepo, mockMediaSvc, mockRecipeMediaSvc)

	// Create a test context
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	// Define test cases.
	testCases := []struct {
		name        string
		expectError bool
		uId         uuid.UUID
	}{
		{"reminders fetch successful", false, testUserUuid},
		{"reminders fetch failed, user exist but no reminders", true, testUserUuidWithoutReminders},
		{"reminders fetch failed, user doesn't exist", true, uuid.New()},
	}

	// Execute test cases.
	// Iterate through each test case and run the corresponding test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			res, err := s.GetAllReminders(c, tc.uId)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned
				require.Error(t, err)
				assert.Nil(t, res)
			} else {
				// If no error is expected, ensure there is no error returned
				require.NoError(t, err)
			}
		})
	}
}
