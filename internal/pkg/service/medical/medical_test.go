package medical

import (
	"bytes"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

type mockMedicalRepository struct{}

func (m mockMedicalRepository) Create(value interface{}) error {
	return nil
}

func (m mockMedicalRepository) Update(value interface{}) error {
	return nil
}

func (m mockMedicalRepository) Find(out interface{}, conditions ...interface{}) error {
	return nil
}

func TestCreateRecordFromFile(t *testing.T) {

	// Create a mock repository
	repo := &mockMedicalRepository{}

	svc := NewService(repo)

	// Create a test context with file and request
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	// Prepare the CSV file data
	fileContent := `First Name;column1;Last Name;column2;Cjppu Number;Profession Number
John;column1;Doe;column2;num1;num2
Jane;column1;Smith;column2;num1;num2`

	// Create a temporary file for the CSV data
	tempFile, err := os.CreateTemp("", "test.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Write the CSV data to the temporary file
	_, err = tempFile.WriteString(fileContent)
	if err != nil {
		t.Fatal(err)
	}
	tempFile.Close()

	// Open the temporary file
	file, err := os.Open(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// Create a multipart writer for the form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file part using the CSV file data
	part, err := writer.CreateFormFile("file", "test.csv")
	if err != nil {
		t.Fatal(err)
	}

	// Copy the file data into the form file part
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatal(err)
	}

	// Close the multipart writer
	writer.Close()

	// Set the request's content type and body
	c.Request, err = http.NewRequest(http.MethodPost, "/create", body)
	if err != nil {
		t.Fatal(err)
	}
	c.Request.Header.Set("Content-Type", writer.FormDataContentType())

	//create a context without file
	ctx, _ := gin.CreateTestContext(nil)
	ctx.Request, err = http.NewRequest(http.MethodPost, "/create", body)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name        string
		expectError bool
		context     *gin.Context
	}{
		{"medical creation from file successful", false, c},
		{"medical creation from file failed, bad request", true, ctx},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			statusCode, err := svc.CreateRecordFromFile(tc.context)
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

func TestGetAllMedicalRecords(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockMedicalRepository{}
	s := NewService(mockRepo)

	testCases := []struct {
		name        string
		expectError bool
	}{
		{"fetch medical successful", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := s.GetAllMedicalRecords()
			// Check the result based on the test case's expected error state.
			if tc.expectError {
				// If an error is expected, ensure there is an error returned and the healthServices is nil.
				require.Error(t, err)
			} else {
				// If no error is expected, ensure there is no error returned and the healthServices is not nil.
				require.NoError(t, err)
			}
		})
	}
}

func TestAddRatingToMedical(t *testing.T) {

	// Create a mock repository
	repo := &mockMedicalRepository{}

	svc := NewService(repo)

	testCases := []struct {
		name        string
		request     *entity.MedicalRating
		expectError bool
	}{
		{"medical rating successful", &entity.MedicalRating{MedicalID: 1, ReminderID: 1}, false},
		{"medical rating failed, bad request", &entity.MedicalRating{MedicalID: 0, ReminderID: 0}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			statusCode, err := svc.AddRatingToMedical(tc.request)
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
