package category

import (
	"errors"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

var testUserUuid = uuid.MustParse("24df3f36-ca63-11ed-afa1-0242ac120002")
var testArticleUuid = uuid.MustParse("1a09e86a-4011-4290-85f3-8e2d6f7f0866")
var testCatUuid = uuid.MustParse("bfb23f5c-a664-432b-b6cc-b7cd17bacf5b")

type mockCategoryRepository struct{}

func (m mockCategoryRepository) FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error) {
	if uuid == testCatUuid {
		cat := &entity.Category{
			ID:   1,
			UUID: testCatUuid,
			Name: "Test Name",
		}
		return cat, nil
	}
	return nil, errors.New("cat not found")
}

func (m mockCategoryRepository) CreateWithOmit(omit string, value interface{}) error {
	if value == nil {
		return errors.New("input value cannot be nil")
	}
	return nil
}

func (m mockCategoryRepository) Update(value interface{}) error {
	return nil
}

func (m mockCategoryRepository) First(out interface{}, conditions ...interface{}) error {
	return nil
}

func (m mockCategoryRepository) Find(out interface{}, conditions ...interface{}) error {
	return nil
}

func (m mockCategoryRepository) Delete(out interface{}) error {
	return nil
}

func TestCreateCategory(t *testing.T) {

	// Create a mock repository
	repo := &mockCategoryRepository{}

	svc := NewService(repo)

	// Create a test context and request
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	testCases := []struct {
		name        string
		category    *entity.Category
		expectError bool
	}{
		{"category creation successful", &entity.Category{Name: "Test category"}, false},
		{"category creation failed", nil, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			statusCode, err := svc.CreateCategory(c, tc.category)
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

func TestUpdateCategory(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockCategoryRepository{}
	s := NewService(mockRepo)

	updateData := &entity.Category{
		ID:   1,
		UUID: uuid.New(),
		Name: "Test Name",
	}

	// Test case 1: category found and updated successfully
	status, err := s.UpdateCategory(testCatUuid, updateData)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)

	// Test case 2: category not found
	status, err = s.UpdateCategory(uuid.New(), updateData)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, status)
}

func TestGetAllCategory(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockCategoryRepository{}
	s := NewService(mockRepo)

	// Test case 1: category fetched successfully
	_, err := s.GetAllCategories()
	assert.Nil(t, err)
}

func TestDeleteCategory(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockCategoryRepository{}
	s := NewService(mockRepo)

	// Create a test context
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	// Test case 1: category found and deleted successfully
	status, err := s.DeleteCategory(c, testCatUuid)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)

	// Test case 2: category not found
	status, err = s.DeleteCategory(c, uuid.New())
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, status)
}
