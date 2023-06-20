package media

import (
	"errors"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

var testMediaUuid = uuid.MustParse("bfb23f5c-a664-432b-b6cc-b7cd17bacf5b")

type mockMediaRepository struct{}

func (m mockMediaRepository) CreateWithOmit(omit string, value interface{}) error {
	if value == nil || reflect.ValueOf(value).IsNil() {
		return errors.New("nil value")
	}
	return nil
}

func (m mockMediaRepository) Delete(value interface{}) error {
	if value == nil || reflect.ValueOf(value).IsNil() {
		return errors.New("nil value")
	}
	return nil
}

func (m mockMediaRepository) Find(model interface{}, dest interface{}, conditions ...interface{}) error {
	if conditions[0] == "id = ?" && conditions[1] == 1 {
		return nil
	}
	return errors.New("not found")
}

func TestFindByMediaID(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockMediaRepository{}
	s := NewService(mockRepo)

	// Test case 1: media found
	err := s.FindByMediaID(1, &entity.Media{})
	assert.Nil(t, err)

	// Test case 2: media not found
	err = s.FindByMediaID(2, &entity.Media{})
	assert.NotNil(t, err)
}

func TestCreateMedia(t *testing.T) {

	// Create a mock repository
	repo := &mockMediaRepository{}

	svc := NewService(repo)

	testCases := []struct {
		name        string
		media       *entity.Media
		expectError bool
	}{
		{"media creation successful", &entity.Media{UUID: testMediaUuid}, false},
		{"media creation failed", nil, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := svc.CreateMedia(tc.media)
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

func TestDeleteMedia(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockMediaRepository{}
	s := NewService(mockRepo)

	testCases := []struct {
		name        string
		media       *entity.Media
		expectError bool
	}{
		{"media deletion successful", &entity.Media{UUID: testMediaUuid}, false},
		{"media deletion failed", nil, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := s.DeleteMedia(tc.media)
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
