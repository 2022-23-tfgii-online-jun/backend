package question

import (
	"errors"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var testUserUuid = uuid.MustParse("24df3f36-ca63-11ed-afa1-0242ac120002")

type mockQuestionRepository struct{}

func (m mockQuestionRepository) FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error) {
	if uuid == testUserUuid {
		usr := &entity.User{
			ID:   1,
			UUID: testUserUuid,
		}
		return usr, nil
	}
	return nil, errors.New("not found")
}

func (m mockQuestionRepository) Create(value interface{}) error {
	return nil
}

func (m mockQuestionRepository) CreateWithOmit(omit string, value interface{}) error {
	return nil
}

func (m mockQuestionRepository) Update(value interface{}) error {
	return nil
}

func (m mockQuestionRepository) First(out interface{}, conditions ...interface{}) error {
	return nil
}

func (m mockQuestionRepository) Find(out interface{}, conditions ...interface{}) error {
	return nil
}

func (m mockQuestionRepository) Delete(out interface{}) error {
	return nil
}

func TestCreateQuestion(t *testing.T) {

	// Create a mock repository
	repo := &mockQuestionRepository{}

	svc := NewService(repo)

	// Create a test context and request
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	testCases := []struct {
		name        string
		request     *entity.RequestCreateQuestion
		expectError bool
		userUUID    uuid.UUID
	}{
		{"question creation successful", &entity.RequestCreateQuestion{Text: "test"}, false, testUserUuid},
		{"question creation failed, user doesn't exist", &entity.RequestCreateQuestion{Text: "test"}, true, uuid.New()},
		{"question creation failed, invalid request", nil, true, testUserUuid},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := svc.CreateQuestion(c, tc.userUUID, tc.request)
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

func TestGetAllQuestions(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockQuestionRepository{}
	s := NewService(mockRepo)

	// questions fetched successfully
	_, err := s.GetAllQuestions()
	assert.Nil(t, err)
}
