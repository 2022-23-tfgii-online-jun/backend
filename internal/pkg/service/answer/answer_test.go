package answer

import (
	"errors"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"testing"
)

type mockAnswerRepository struct{}

var testUserUuid = uuid.MustParse("24df3f36-ca63-11ed-afa1-0242ac120002")
var testQuestionUuid = uuid.MustParse("1a09e86a-4011-4290-85f3-8e2d6f7f0866")

func (m *mockAnswerRepository) FindByUUID(uId uuid.UUID, out interface{}) (interface{}, error) {
	return nil, nil
}

func (m *mockAnswerRepository) CreateWithOmit(string, interface{}) error {
	return nil
}

type mockRepositoryOverride struct {
	mockAnswerRepository
	findByUUIDFunc func(uuid.UUID, interface{}) (interface{}, error)
}

func (m *mockRepositoryOverride) Create(value interface{}) error {
	return nil
}

func (m *mockRepositoryOverride) Update(value interface{}) error {
	return nil
}

func (m *mockRepositoryOverride) First(out interface{}, conditions ...interface{}) error {
	return nil
}

func (m *mockRepositoryOverride) Find(out interface{}, conditions ...interface{}) error {
	return nil
}

func (m *mockRepositoryOverride) Delete(out interface{}) error {
	return nil
}

func (m *mockRepositoryOverride) FindByUUID(uuid uuid.UUID, entity interface{}) (interface{}, error) {
	return m.findByUUIDFunc(uuid, entity)
}

func TestCreateAnswer(t *testing.T) {
	// Create a mock repository
	repo := &mockRepositoryOverride{
		findByUUIDFunc: func(uId uuid.UUID, out interface{}) (interface{}, error) {
			if uId == testUserUuid {
				user := &entity.User{
					ID:   1,
					UUID: testUserUuid,
				}
				return user, nil
			}
			if uId == testQuestionUuid {
				question := &entity.Question{
					ID:     1,
					UUID:   uId,
					UserID: 1,
				}
				return question, nil
			}
			return nil, nil
		},
	}

	// Create the answer service with the mock repository
	svc := NewService(repo)

	// Create a test context and request
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	// Create a test UUID for the user and question
	userUUID := testUserUuid
	questionUUID := testQuestionUuid

	// Create a test request for creating an answer
	createReq := &entity.RequestCreateAnswer{
		Text: "Test answer",
	}

	// Positive case: User and question are found, answer is created successfully
	statusCode, err := svc.CreateAnswer(c, userUUID, questionUUID, createReq)

	// Check the result
	if err != nil {
		t.Errorf("CreateAnswer returned an error: %v", err)
	}

	if statusCode != http.StatusOK {
		t.Errorf("CreateAnswer returned an unexpected status code: %d", statusCode)
	}

	// Negative case: User is not found
	repo.findByUUIDFunc = func(uuid.UUID, interface{}) (interface{}, error) {
		return nil, errors.New("user not found")
	}

	statusCode, err = svc.CreateAnswer(c, userUUID, questionUUID, createReq)

	// Check the result
	expectedStatusCode := http.StatusNotFound
	if err == nil {
		t.Error("CreateAnswer did not return an error for user not found")
	}

	if statusCode != expectedStatusCode {
		t.Errorf("CreateAnswer returned an unexpected status code for user not found. Expected: %d, Got: %d", expectedStatusCode, statusCode)
	}

	// Negative case: Question is not found
	repo.findByUUIDFunc = func(uuid.UUID, interface{}) (interface{}, error) {
		return &entity.User{}, nil
	}

	statusCode, err = svc.CreateAnswer(c, userUUID, questionUUID, createReq)

	// Check the result
	expectedStatusCode = http.StatusInternalServerError
	if err == nil {
		t.Error("CreateAnswer did not return an error for question not found")
	}

	if statusCode != expectedStatusCode {
		t.Errorf("CreateAnswer returned an unexpected status code for question not found. Expected: %d, Got: %d", expectedStatusCode, statusCode)
	}
}
