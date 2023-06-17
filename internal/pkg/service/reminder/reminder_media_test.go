package reminder

import (
	"errors"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type mockReminderMediaRepository struct{}

func (m mockReminderMediaRepository) Create(value interface{}) error {
	return nil
}

func (m mockReminderMediaRepository) Delete(value interface{}) error {
	if value == nil || reflect.ValueOf(value).IsNil() {
		return errors.New("nil value")
	}
	return nil
}

func (m mockReminderMediaRepository) Find(model interface{}, dest interface{}, conditions ...interface{}) error {
	if conditions[0] == "reminder_id = ?" && conditions[1] == 1 {
		return nil
	}
	return errors.New("not found")
}

func TestFindByRecipeID(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockReminderMediaRepository{}
	s := NewReminderMediaService(mockRepo)

	medias := []*entity.ReminderMedia{}

	// Test case 1: media found
	err := s.FindByReminderID(1, &medias)
	assert.Nil(t, err)

	media := []*entity.ReminderMedia{}

	// Test case 2: media not found
	err = s.FindByReminderID(2, &media)
	assert.NotNil(t, err)
}

func TestCreateReminderMedia(t *testing.T) {

	// Create a mock repository
	mockRepo := &mockReminderMediaRepository{}
	s := NewReminderMediaService(mockRepo)

	testCases := []struct {
		name        string
		media       *entity.ReminderMedia
		expectError bool
	}{
		{"reminder media creation successful", &entity.ReminderMedia{MediaID: 1}, false},
		{"reminder media creation failed", nil, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := s.CreateReminderMedia(tc.media)
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

func TestDeleteReminderMedia(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockReminderMediaRepository{}
	s := NewReminderMediaService(mockRepo)

	testCases := []struct {
		name        string
		req         *entity.ReminderMedia
		expectError bool
	}{
		{"reminder media deletion successful", &entity.ReminderMedia{MediaID: 1}, false},
		{"reminder media deletion failed", nil, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := s.DeleteReminderMedia(tc.req)
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
