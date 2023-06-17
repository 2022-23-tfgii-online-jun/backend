package recipe

import (
	"errors"
	"fmt"
	"net/http"
	"path"

	"github.com/emur-uy/backend/config"
	aws "github.com/emur-uy/backend/internal/infra/repositories/spaces"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	ErrTypeAssertionFailed = errors.New("type assertion failed")
	ErrCreatingRecipe      = errors.New("error creating recipe")
	ErrUpdatingRecipe      = errors.New("error updating recipe")
	ErrDeletingRecipe      = errors.New("error deleting recipe")
	ErrCreatingMedia       = errors.New("error creating media")
	ErrFindingRecipeMedia  = errors.New("error finding recipe media")
	ErrCreatingRecipeMedia = errors.New("error creating recipe media association")
	ErrDeletingRecipeMedia = errors.New("error deleting recipe media association")
	ErrDeletingMedia       = errors.New("error deleting media")
	ErrUnsupportedFileType = errors.New("unsupported file type")
	ErrFileNotFound        = errors.New("file not found")
)

const (
	PNG  = "image/png"
	JPEG = "image/jpeg"
)

// service struct holds the necessary dependencies for the recipe service
type service struct {
	repo               ports.RecipeRepository
	mediaService       ports.MediaService
	recipeMediaService ports.RecipeMediaService
}

// NewService returns a new instance of the recipe service with the given recipe repository, media service, and recipe media service.
func NewService(recipeRepo ports.RecipeRepository, mediaService ports.MediaService, recipeMediaService ports.RecipeMediaService) ports.RecipeService {
	return &service{
		repo:               recipeRepo,
		mediaService:       mediaService,
		recipeMediaService: recipeMediaService,
	}
}

// CreateRecipe is the service for creating a recipe and saving it in the database.
func (s *service) CreateRecipe(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateRecipe) (*entity.Recipe, error) {
	user := &entity.User{}

	// Find user by UUID
	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		// Return error if the user is not found
		return nil, err
	}

	// Perform type assertion to convert foundUser to *entity.User
	user, ok := foundUser.(*entity.User)
	if !ok {
		return nil, ErrTypeAssertionFailed
	}

	// Call the processUploadRequestFile function to handle the image upload and create the media entry
	fileProcessCode, fileUrls, err := processUploadRequestFiles(s, c)
	if err != nil || fileProcessCode != http.StatusOK {
		return nil, fmt.Errorf("error processing content upload file: %s", err)
	}

	// Create a new recipe
	recipe := &entity.Recipe{
		Name:        createReq.Name,
		Ingredients: createReq.Ingredients,
		Elaboration: createReq.Elaboration,
		Time:        createReq.Time,
		Category:    createReq.Category,
	}

	// Save the recipe to the database
	err = s.repo.CreateWithOmit("uuid", recipe)
	if err != nil {
		return nil, ErrCreatingRecipe
	}

	// For each uploaded file, create a new media entry and then a new RecipeMedia entry
	for _, fileUrl := range fileUrls {
		media := &entity.Media{
			MediaURL: fileUrl,
		}

		err = s.mediaService.CreateMedia(media)
		if err != nil {
			return nil, ErrCreatingMedia
		}

		recipeMedia := &entity.RecipeMedia{
			RecipeID: recipe.ID,
			MediaID:  media.ID,
		}
		err = s.recipeMediaService.CreateRecipeMedia(recipeMedia)
		if err != nil {
			return nil, fmt.Errorf("error creating RecipeMedia: %s", err)
		}
	}

	return recipe, nil
}

// UpdateRecipe is the service for updating a recipe in the database.
func (s *service) UpdateRecipe(c *gin.Context, recipeUUID uuid.UUID, updateReq *entity.RequestUpdateRecipe) (int, error) {
	// Find the existing recipe by UUID
	recipe := &entity.Recipe{}
	foundRecipe, err := s.repo.FindByUUID(recipeUUID, recipe)
	if err != nil {
		// Return error if the recipe is not found
		return http.StatusNotFound, err
	}

	// Perform type assertion to convert foundRecipe to *entity.Recipe
	recipe, ok := foundRecipe.(*entity.Recipe)
	if !ok {
		return http.StatusInternalServerError, ErrTypeAssertionFailed
	}

	if updateReq == nil {
		return http.StatusBadRequest, errors.New("nil payload")
	}

	// Update the recipe fields with the new data from the update request
	recipe.Name = updateReq.Name

	// Update the recipe in the database
	err = s.repo.Update(recipe)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error updating recipe: %s", err)
	}

	fileProcessCode, fileUrls, err := processUploadRequestFiles(s, c)
	if err != nil || fileProcessCode != http.StatusOK {
		return http.StatusInternalServerError, fmt.Errorf("error processing content upload file: %s", err)
	}

	// Get existing recipe media data
	recipeMedias := []*entity.RecipeMedia{}
	err = s.recipeMediaService.FindByRecipeID(recipe.ID, &recipeMedias)
	if err != nil {
		return http.StatusInternalServerError, ErrFindingRecipeMedia
	}

	// For each uploaded file, create a new media entry and a new recipe_media association
	for _, fileUrl := range fileUrls {
		media := &entity.Media{
			MediaURL: fileUrl,
		}
		err = s.mediaService.CreateMedia(media)
		if err != nil {
			return http.StatusInternalServerError, ErrCreatingMedia
		}
		recipeMedia := &entity.RecipeMedia{
			RecipeID: recipe.ID,
			MediaID:  media.ID,
		}
		err = s.recipeMediaService.CreateRecipeMedia(recipeMedia)
		if err != nil {
			return http.StatusInternalServerError, ErrCreatingRecipeMedia
		}
	}

	// Delete old media entries
	for _, recipeMedia := range recipeMedias {
		mediaID := recipeMedia.MediaID

		// Delete the recipeMedia entry
		err = s.recipeMediaService.DeleteRecipeMedia(recipeMedia)
		if err != nil {
			return http.StatusInternalServerError, ErrDeletingRecipeMedia
		}

		// Delete the media from the repository
		media := entity.Media{
			ID: mediaID,
		}
		err = s.mediaService.DeleteMedia(&media)
		if err != nil {
			return http.StatusInternalServerError, ErrDeletingMedia
		}
	}

	// Return the HTTP OK status code if the update is successful
	return http.StatusOK, nil
}

// DeleteRecipe deletes a recipe from the database by its UUID.
func (s *service) DeleteRecipe(c *gin.Context, recipeUUID uuid.UUID) (int, error) {
	// Retrieve the recipe from the repository by its UUID.
	recipe := &entity.Recipe{}
	foundRecipe, err := s.repo.FindByUUID(recipeUUID, recipe)
	if err != nil {
		// Return an error response if the recipe is not found.
		return http.StatusNotFound, err
	}

	// Perform type assertion to convert foundRecipe to *entity.Recipe.
	recipe, ok := foundRecipe.(*entity.Recipe)
	if !ok {
		return http.StatusInternalServerError, ErrTypeAssertionFailed
	}

	// Get recipe media associations by recipe ID
	recipeMedias := []*entity.RecipeMedia{}
	err = s.recipeMediaService.FindByRecipeID(recipe.ID, &recipeMedias)
	if err != nil {
		return http.StatusInternalServerError, ErrFindingRecipeMedia
	}

	// Iterate over each recipe_media association
	for _, recipeMedia := range recipeMedias {
		media := &entity.Media{}
		// Find the media by ID
		err := s.mediaService.FindByMediaID(recipeMedia.MediaID, media)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		// Delete the recipe_media association from the repository
		err = s.recipeMediaService.DeleteRecipeMedia(recipeMedia)
		if err != nil {
			return http.StatusInternalServerError, ErrDeletingRecipeMedia
		}

		// Delete the media from the repository
		err = s.mediaService.DeleteMedia(media)
		if err != nil {
			return http.StatusInternalServerError, ErrDeletingMedia
		}
	}

	// Delete the recipe from the repository
	err = s.repo.Delete(recipe)
	if err != nil {
		return http.StatusInternalServerError, ErrDeletingRecipe
	}

	// Return a success response.
	return http.StatusOK, nil
}

// GetAllRecipes returns all recipes stored in the database with associated image URLs.
func (s *service) GetAllRecipes() ([]*entity.RecipeWithMediaURLs, error) {
	// Get all recipes from the database
	var recipes []*entity.Recipe
	if err := s.repo.Find(&recipes); err != nil {
		return nil, err
	}

	recipesWithMediaURLs := make([]*entity.RecipeWithMediaURLs, len(recipes))

	for i, recipe := range recipes {
		// Get associated media for the recipe
		recipeMedias := []*entity.RecipeMedia{}
		err := s.recipeMediaService.FindByRecipeID(recipe.ID, &recipeMedias)
		if err != nil {
			return nil, err
		}

		mediaURLs := make([]string, len(recipeMedias))

		// Get media URLs
		for j, recipeMedia := range recipeMedias {
			media := &entity.Media{}
			err = s.mediaService.FindByMediaID(recipeMedia.MediaID, media)
			if err != nil {
				return nil, err
			}

			mediaURLs[j] = media.MediaURL
		}

		recipesWithMediaURLs[i] = &entity.RecipeWithMediaURLs{
			Recipe:    recipe,
			MediaURLs: mediaURLs,
		}
	}

	return recipesWithMediaURLs, nil
}

// VoteRecipe is the service for voting a recipe in the database.
func (s *service) VoteRecipe(c *gin.Context, userUUID uuid.UUID, recipeUUID uuid.UUID, vote int) (int, error) {
	if vote < 1 || vote > 5 {
		return http.StatusBadRequest, errors.New("invalid vote value, must be between 1 and 5")
	}

	// Get user by UUID
	user := &entity.User{}
	_, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		return http.StatusNotFound, fmt.Errorf("error finding user: %s", err)
	}

	// Get recipe by UUID
	recipe := &entity.Recipe{}
	_, err = s.repo.FindByUUID(recipeUUID, recipe)
	if err != nil {
		return http.StatusNotFound, fmt.Errorf("error finding recipe: %s", err)
	}

	type RatingRecipe struct {
		ID       int
		UserID   int
		RecipeID int
		Level    int
	}

	// Check if the user has already voted for this recipe
	existingVote := &RatingRecipe{}
	err = s.repo.FindItemByIDs(user.ID, recipe.ID, "rating_recipes", "user_id", "recipe_id", existingVote)
	if err == nil {
		// Update existing vote value
		existingVote.Level = vote
		return http.StatusOK, s.repo.Update(existingVote)
	}

	// Create a new vote record
	recipeVote := &RatingRecipe{
		UserID:   user.ID,
		RecipeID: recipe.ID,
		Level:    vote,
	}
	return http.StatusOK, s.repo.Create(recipeVote)
}

var uploadFunc = aws.UploadFileToS3Stream

// processUploadRequestFiles processes the file upload request
func processUploadRequestFiles(s *service, c *gin.Context) (int, []string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return http.StatusBadRequest, nil, fmt.Errorf("get form err: %s", err.Error())
	}
	files := form.File["file"]
	if files == nil {
		return http.StatusBadRequest, nil, ErrFileNotFound
	}

	var fileUrls []string

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return http.StatusInternalServerError, nil, fmt.Errorf("failed to open file: %s", err)
		}
		defer src.Close()

		fileType := file.Header.Get("Content-Type")
		if fileType != "image/png" && fileType != "image/jpeg" {
			return http.StatusBadRequest, nil, ErrUnsupportedFileType
		}

		fileExt := path.Ext(file.Filename)
		fileNameUuid := uuid.New()

		uploadPath := fmt.Sprintf("%s/%s", config.Get().AwsFolderName, fmt.Sprintf("%s%s", fileNameUuid.String(), fileExt))
		url, err := uploadFunc(src, uploadPath, true)
		if err != nil || url == "" {
			return http.StatusInternalServerError, nil, fmt.Errorf("s3 upload error: %s", err.Error())
		}

		fileUrls = append(fileUrls, url)
	}

	return http.StatusOK, fileUrls, nil
}
