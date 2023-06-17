package recipe

import (
	"errors"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type mockRecipeMediaRepository struct{}

func (m mockRecipeMediaRepository) Create(value interface{}) error {
	return nil
}

func (m mockRecipeMediaRepository) Delete(value interface{}) error {
	if value == nil || reflect.ValueOf(value).IsNil() {
		return errors.New("nil value")
	}
	return nil
}

func (m mockRecipeMediaRepository) Find(model interface{}, dest interface{}, conditions ...interface{}) error {
	if conditions[0] == "recipe_id = ?" && conditions[1] == 1 {
		return nil
	}
	return errors.New("not found")
}

func TestFindByRecipeID(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockRecipeMediaRepository{}
	s := NewRecipeMediaService(mockRepo)

	medias := []*entity.RecipeMedia{}

	// Test case 1: media found
	err := s.FindByRecipeID(1, &medias)
	assert.Nil(t, err)

	media := []*entity.RecipeMedia{}

	// Test case 2: media not found
	err = s.FindByRecipeID(2, &media)
	assert.NotNil(t, err)
}

func TestCreateRecipeMedia(t *testing.T) {

	// Create a mock repository
	mockRepo := &mockRecipeMediaRepository{}
	s := NewRecipeMediaService(mockRepo)

	testCases := []struct {
		name        string
		media       *entity.RecipeMedia
		expectError bool
	}{
		{"recipe media creation successful", &entity.RecipeMedia{MediaID: 1}, false},
		{"recipe media creation failed", nil, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := s.CreateRecipeMedia(tc.media)
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

func TestDeleteRecipeMedia(t *testing.T) {
	// Initialize the mock repository and service.
	mockRepo := &mockRecipeMediaRepository{}
	s := NewRecipeMediaService(mockRepo)

	testCases := []struct {
		name        string
		req         *entity.RecipeMedia
		expectError bool
	}{
		{"recipe media deletion successful", &entity.RecipeMedia{MediaID: 1}, false},
		{"recipe media deletion failed", nil, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := s.DeleteRecipeMedia(tc.req)
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
