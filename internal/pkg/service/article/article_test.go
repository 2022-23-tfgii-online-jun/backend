package article

import (
	"bytes"
	"errors"
	"fmt"
	aws "github.com/emur-uy/backend/internal/infra/repositories/spaces"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"testing"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testUserUuid = uuid.MustParse("24df3f36-ca63-11ed-afa1-0242ac120002")
var testArticleUuid = uuid.MustParse("1a09e86a-4011-4290-85f3-8e2d6f7f0866")
var testCatUuid = uuid.MustParse("bfb23f5c-a664-432b-b6cc-b7cd17bacf5b")

// MockArticleRepository is a mock implementation of the ArticleRepository interface for testing.
type MockArticleRepository struct{}

func (m *MockArticleRepository) Create(value interface{}) error {
	return nil
}

func (m *MockArticleRepository) Find(out interface{}, conditions ...interface{}) error {
	return nil
}

func (m *MockArticleRepository) Delete(out interface{}) error {
	return nil
}

// CreateWithOmit is a mock implementation of the CreateWithOmit method.
func (m *MockArticleRepository) CreateWithOmit(omit string, value interface{}) error {
	return nil
}

// Update is a mock implementation of the Update method.
func (m *MockArticleRepository) Update(value interface{}) error {
	return nil
}

// First is a mock implementation of the First method.
func (m *MockArticleRepository) First(out interface{}, conditions ...interface{}) error {
	return nil
}

// UpdateColumns is a mock implementation of the UpdateColumns method.
func (m *MockArticleRepository) UpdateColumns(value interface{}, column string, updateValue interface{}) error {
	return nil
}

// FindByUUID is a mock implementation of the FindByUUID method.
func (m *MockArticleRepository) FindByUUID(uId uuid.UUID, out interface{}) (interface{}, error) {
	if uId == testUserUuid {
		usr := &entity.User{
			ID:   1,
			UUID: testUserUuid,
		}
		return usr, nil
	}
	if uId == testArticleUuid {
		res := &entity.Article{
			ID:   1,
			UUID: testArticleUuid,
		}
		return res, nil
	}
	if uId == testCatUuid {
		res := &entity.Category{
			ID:   1,
			UUID: testArticleUuid,
		}
		return res, nil
	}
	return nil, errors.New("not found")
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

type MockArticleMediaService struct{}

func (m MockArticleMediaService) CreateArticleMedia(articleMedia *entity.ArticleMedia) error {
	return nil
}

func (m MockArticleMediaService) DeleteArticleMedia(articleMedia *entity.ArticleMedia) error {
	return nil
}

func (m MockArticleMediaService) FindByArticleID(id int, i *[]*entity.ArticleMedia) error {
	if id == 1 {
		return nil
	}
	return errors.New("not found")
}

func mockUploadFileToS3Stream(src io.Reader, uploadPath string, isPublic bool) (string, error) {
	// Mocked implementation of the upload function
	return "mocked-url", nil
}

func TestCreateArticle(t *testing.T) {

	// Set up the mock repository and service.
	mockRepo := &MockArticleRepository{}
	mockMediaSvc := &MockMediaService{}
	mockArticleMediaSvc := &MockArticleMediaService{}
	s := NewService(mockRepo, mockMediaSvc, mockArticleMediaSvc)

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

	// Create a test request for creating an article
	createReq := &entity.RequestCreateArticle{
		Title:   "Test article",
		Content: "Test content",
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
		{"user exist & article creation successful", testUserUuid, false, *c},
		{"user doesn't exist & article creation failed", uuid.New(), true, *ctx},
		{"file not found & article creation failed", testUserUuid, true, *ctx},
	}

	// Execute test cases.
	// Iterate through each test case and run the corresponding test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := s.CreateArticle(&tc.context, createReq)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned and the statusCode is not OK.
				require.Error(t, err)
			} else {
				// If no error is expected, ensure there is no error returned and the statusCode is OK.
				require.NoError(t, err)
			}
		})
	}
}

func TestUpdateArticle(t *testing.T) {
	// Set up the mock repository and service.
	mockRepo := &MockArticleRepository{}
	mockMediaSvc := &MockMediaService{}
	mockArticleMediaSvc := &MockArticleMediaService{}
	s := NewService(mockRepo, mockMediaSvc, mockArticleMediaSvc)

	// Create a test request for creating an article
	req := &entity.RequestUpdateArticle{
		Title:   "Test article",
		Content: "Test content",
	}

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

	// Create a test context without request file
	ctx, _ := gin.CreateTestContext(nil)
	ctx.Request = httptest.NewRequest(http.MethodPost, "/", fileBuf)
	ctx.Request.Header.Set("Content-Type", "multipart/form-data")

	uploadFunc = mockUploadFileToS3Stream
	defer func() {
		// Restore the original upload function after the test
		uploadFunc = aws.UploadFileToS3Stream
	}()

	// Define test cases.
	testCases := []struct {
		name        string
		uId         uuid.UUID
		expectError bool
		request     *entity.RequestUpdateArticle
		context     gin.Context
	}{
		{"article exist & article update successful", testArticleUuid, false, req, *c},
		{"article doesn't exist & article updation failed", uuid.New(), true, nil, *c},
		{"article updation failed, article medias don't exists", testUserUuid, true, req, *c},
		{"article updation failed, invalid request", testUserUuid, true, nil, *ctx},
	}

	// Execute test cases.
	// Iterate through each test case and run the corresponding test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			statusCode, err := s.UpdateArticle(c, tc.uId, tc.request)
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

func TestDeleteArticle(t *testing.T) {
	// Set up the mock repository and service.
	mockRepo := &MockArticleRepository{}
	mockMediaSvc := &MockMediaService{}
	mockArticleMediaSvc := &MockArticleMediaService{}
	s := NewService(mockRepo, mockMediaSvc, mockArticleMediaSvc)

	// Create a test context
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	// Define test cases.
	testCases := []struct {
		name        string
		uId         uuid.UUID
		expectError bool
	}{
		{"article exist & article deletion successful", testArticleUuid, false},
		{"article doesn't exist & article deletion failed", uuid.New(), true},
	}

	// Execute test cases.
	// Iterate through each test case and run the corresponding test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			statusCode, err := s.DeleteArticle(c, tc.uId)
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

func TestGetAllArticles(t *testing.T) {
	// Set up the mock repository and service.
	mockRepo := &MockArticleRepository{}
	mockMediaSvc := &MockMediaService{}
	mockArticleMediaSvc := &MockArticleMediaService{}
	s := NewService(mockRepo, mockMediaSvc, mockArticleMediaSvc)

	// Define test cases.
	testCases := []struct {
		name        string
		expectError bool
	}{
		{"article fetch successful", false},
	}

	// Execute test cases.
	// Iterate through each test case and run the corresponding test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			articles, err := s.GetAllArticles()
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned
				require.Error(t, err)
				assert.Nil(t, articles)
			} else {
				// If no error is expected, ensure there is no error returned
				require.NoError(t, err)
			}
		})
	}
}

func TestAddArticleToCategory(t *testing.T) {
	// Set up the mock repository and service.
	mockRepo := &MockArticleRepository{}
	mockMediaSvc := &MockMediaService{}
	mockArticleMediaSvc := &MockArticleMediaService{}
	s := NewService(mockRepo, mockMediaSvc, mockArticleMediaSvc)

	// Define test cases.
	testCases := []struct {
		name        string
		articleUId  uuid.UUID
		catUid      uuid.UUID
		expectError bool
	}{
		{"article & category exist & category add to article successful", testArticleUuid, testCatUuid, false},
		{"article doesn't exist & category add to article failed", uuid.New(), testCatUuid, true},
		{"category doesn't exist & category add to article failed", testArticleUuid, uuid.New(), true},
	}

	// Execute test cases.
	// Iterate through each test case and run the corresponding test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := s.AddArticleToCategory(tc.catUid, tc.articleUId)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned .
				require.Error(t, err)
			} else {
				// If no error is expected, ensure there is no error returned
				require.NoError(t, err)
			}
		})
	}
}
