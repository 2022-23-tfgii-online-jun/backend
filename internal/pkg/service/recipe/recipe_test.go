package recipe

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
var testRecipeUuid = uuid.MustParse("1a09e86a-4011-4290-85f3-8e2d6f7f0866")
var testRecipeUuidWithoutMedias = uuid.MustParse("bfb23f5c-a664-432b-b6cc-b7cd17bacf5b")

// MockRecipeRepository is a mock implementation of the RecipeRepository interface for testing.
type MockRecipeRepository struct{}

func (m *MockRecipeRepository) FindItemByIDs(firstID, secondID int, tableName string, column1Name string, column2Name string, dest interface{}) error {
	return nil
}

func (m *MockRecipeRepository) Create(value interface{}) error {
	return nil
}

func (m *MockRecipeRepository) CreateWithOmit(omit string, value interface{}) error {
	return nil
}

func (m *MockRecipeRepository) Update(value interface{}) error {
	return nil
}

func (m *MockRecipeRepository) First(out interface{}, conditions ...interface{}) error {
	return nil
}

func (m *MockRecipeRepository) Find(out interface{}, conditions ...interface{}) error {
	return nil
}

func (m *MockRecipeRepository) Delete(out interface{}) error {
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

type MockRecipeMediaService struct{}

func (m MockRecipeMediaService) CreateRecipeMedia(recipeMedia *entity.RecipeMedia) error {
	return nil
}

func (m MockRecipeMediaService) DeleteRecipeMedia(recipeMedia *entity.RecipeMedia) error {
	return nil
}

func (m MockRecipeMediaService) FindByRecipeID(id int, i *[]*entity.RecipeMedia) error {
	if id == 1 {
		return nil
	}
	return errors.New("not found")
}

// FindByUUID is a mock implementation of the FindByUUID method.
func (m *MockRecipeRepository) FindByUUID(uId uuid.UUID, out interface{}) (interface{}, error) {
	if uId == testUserUuid {
		usr := &entity.User{
			ID:   1,
			UUID: testUserUuid,
		}
		return usr, nil
	}
	if uId == testRecipeUuid {
		res := &entity.Recipe{
			ID:   1,
			UUID: testRecipeUuid,
		}
		return res, nil
	}
	if uId == testRecipeUuidWithoutMedias {
		res := &entity.Recipe{
			ID:   2,
			UUID: testRecipeUuidWithoutMedias,
		}
		return res, nil
	}
	return nil, errors.New("not found")
}

func mockUploadFileToS3Stream(src io.Reader, uploadPath string, isPublic bool) (string, error) {
	// Mocked implementation of the upload function
	return "mocked-url", nil
}

func TestCreateRecipe(t *testing.T) {

	// Set up the mock repository and service.
	mockRepo := &MockRecipeRepository{}
	mockMediaSvc := &MockMediaService{}
	mockRecipeMediaSvc := &MockRecipeMediaService{}
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

	createReq := &entity.RequestCreateRecipe{
		Name:     "Test",
		Category: 1,
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
		{"user exist & recipe creation successful", testUserUuid, false, *c},
		{"user doesn't exist & recipe creation failed", uuid.New(), true, *ctx},
		{"file not found & recipe creation failed", testUserUuid, true, *ctx},
	}

	// Execute test cases.
	// Iterate through each test case and run the corresponding test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			res, err := s.CreateRecipe(&tc.context, tc.uId, createReq)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned
				require.Error(t, err)
				assert.Nil(t, res)
			} else {
				// If no error is expected, ensure there is no error returned
				require.NoError(t, err)
				assert.NotNil(t, res)
			}
		})
	}
}

func TestUpdateRecipe(t *testing.T) {
	// Set up the mock repository and service.
	mockRepo := &MockRecipeRepository{}
	mockMediaSvc := &MockMediaService{}
	mockRecipeMediaSvc := &MockRecipeMediaService{}
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

	req := &entity.RequestUpdateRecipe{
		Name:     "Test",
		Category: 1,
	}

	// Define test cases.
	testCases := []struct {
		name        string
		uId         uuid.UUID
		expectError bool
		request     *entity.RequestUpdateRecipe
		context     gin.Context
	}{
		{"recipe exist & recipe update successful", testRecipeUuid, false, req, *c},
		{"recipe doesn't exist & recipe updation failed", uuid.New(), true, nil, *c},
		{"recipe updation failed, recipe medias don't exists", testRecipeUuidWithoutMedias, true, req, *c},
		{"recipe updation failed, invalid request", testRecipeUuid, true, nil, *ctx},
	}

	// Execute test cases.
	// Iterate through each test case and run the corresponding test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			statusCode, err := s.UpdateRecipe(c, tc.uId, tc.request)
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

func TestDeleteRecipe(t *testing.T) {
	// Set up the mock repository and service.
	mockRepo := &MockRecipeRepository{}
	mockMediaSvc := &MockMediaService{}
	mockRecipeMediaSvc := &MockRecipeMediaService{}
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
		{"recipe exist & recipe deletion successful", testRecipeUuid, false},
		{"recipe doesn't exist & recipe deletion failed", uuid.New(), true},
	}

	// Execute test cases.
	// Iterate through each test case and run the corresponding test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			statusCode, err := s.DeleteRecipe(c, tc.uId)
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

func TestGetAllRecipes(t *testing.T) {
	// Set up the mock repository and service.
	mockRepo := &MockRecipeRepository{}
	mockMediaSvc := &MockMediaService{}
	mockRecipeMediaSvc := &MockRecipeMediaService{}
	s := NewService(mockRepo, mockMediaSvc, mockRecipeMediaSvc)

	// Define test cases.
	testCases := []struct {
		name        string
		expectError bool
	}{
		{"recipes fetch successful", false},
	}

	// Execute test cases.
	// Iterate through each test case and run the corresponding test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			res, err := s.GetAllRecipes()
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

func TestVoteRecipe(t *testing.T) {
	// Set up the mock repository and service.
	mockRepo := &MockRecipeRepository{}
	mockMediaSvc := &MockMediaService{}
	mockRecipeMediaSvc := &MockRecipeMediaService{}
	s := NewService(mockRepo, mockMediaSvc, mockRecipeMediaSvc)

	// Create a test context
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	// Define test cases.
	testCases := []struct {
		name        string
		recipeUid   uuid.UUID
		userUid     uuid.UUID
		expectError bool
		vote        int
	}{
		{"vote recipe successful, recipe & user exist", testRecipeUuid, testUserUuid, false, 2},
		{"vote recipe failed, user doesn't exist", testRecipeUuid, uuid.New(), true, 2},
		{"vote recipe failed, recipe doesn't exist", uuid.New(), testUserUuid, true, 2},
		{"vote recipe failed, invalid value of vote", testRecipeUuid, testUserUuid, true, 0},
	}

	// Execute test cases.
	// Iterate through each test case and run the corresponding test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			statusCode, err := s.VoteRecipe(c, tc.userUid, tc.recipeUid, tc.vote)
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned .
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
