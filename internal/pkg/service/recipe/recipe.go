package recipe

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/emur-uy/backend/config"
	aws "github.com/emur-uy/backend/internal/infra/repositories/spaces"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	PNG  = "image/png"
	JPEG = "image/jpeg"
)

type service struct {
	repo ports.RecipeRepository
}

// NewService returns a new instance of the recipe service with the given recipe repository.
func NewService(recipeRepo ports.RecipeRepository) ports.RecipeService {
	return &service{
		repo: recipeRepo,
	}
}

// CreateRecipe is the service for creating a recipe and saving it in the database
func (s *service) CreateRecipe(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateRecipe) (int, error) {
	user := &entity.User{}

	// Find user by UUID
	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		// Return error if the user is not found
		return http.StatusNotFound, err
	}
	// Perform type assertion to convert foundUser to *entity.User
	user, ok := foundUser.(*entity.User)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("type assertion failed")
	}

	// Create a new recipe
	recipe := &entity.Recipe{
		Name:        createReq.Name,
		Description: createReq.Description,
		Ingredients: createReq.Ingredients,
		Elaboration: createReq.Elaboration,
		PrepTime:    createReq.PrepTime,
		CookTime:    createReq.CookTime,
		Serving:     createReq.Serving,
		Difficulty:  createReq.Difficulty,
		Nutrition:   createReq.Nutrition,
		UserID:      user.ID,
	}

	// Save the recipe to the database
	err = s.repo.CreateWithOmit("uuid", recipe)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error creating recipe: %s", err)
	}

	// Call the processUploadRequestFile function to handle the image upload and create the media entry
	fileProcessCode, err := processUploadRequestFile(s, c, recipe.ID)
	if err != nil || fileProcessCode != http.StatusOK {
		return http.StatusInternalServerError, fmt.Errorf("error processing content upload file, %s", err)
	}

	// Return the HTTP OK status code if the update is successful
	return http.StatusOK, nil
}

// GetAllRecipes returns all recipes stored in the database
func (s *service) GetAllRecipes() ([]*entity.Recipe, error) {
	// Get all recipes from the database
	var recipes []*entity.Recipe
	if err := s.repo.Find(&recipes); err != nil { // Removed .Error after s.repo.Find()
		return nil, err
	}

	return recipes, nil
}

// UpdateRecipe is the service for updating a recipe in the database
func (s *service) UpdateRecipe(recipeUUID uuid.UUID, updateReq *entity.RequestUpdateRecipe) (int, error) {
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
		return http.StatusInternalServerError, fmt.Errorf("type assertion failed")
	}

	// Update the recipe fields with the new data from the update request
	recipe.Name = updateReq.Name
	recipe.Description = updateReq.Description

	// Update the recipe in the database
	err = s.repo.Update(recipe)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error updating recipe: %s", err)
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
		return http.StatusNotFound, fmt.Errorf("recipe not found")
	}

	// Perform type assertion to convert foundRecipe to *entity.Recipe.
	recipe, ok := foundRecipe.(*entity.Recipe)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("type assertion failed")
	}

	// Delete the recipe from the repository.
	err = s.repo.Delete(recipe)
	if err != nil {
		// Return an error response if there was an issue deleting the recipe.
		return http.StatusInternalServerError, fmt.Errorf("failed to delete recipe")
	}

	// Return a success response.
	return http.StatusOK, nil
}

// VoteRecipe is the service for voting a recipe in the database.
func (s *service) VoteRecipe(c *gin.Context, userUUID uuid.UUID, recipeUUID uuid.UUID, vote int) (int, error) {
	// Validate vote value
	if vote < 1 || vote > 5 {
		return http.StatusNotFound, errors.New("invalid vote value, must be between 1 and 5")
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

	// Crear un nuevo registro de voto
	recipeVote := &RatingRecipe{
		UserID:   user.ID,
		RecipeID: recipe.ID,
		Level:    vote,
	}
	return http.StatusOK, s.repo.Create(recipeVote)
}

func processUploadRequestFile(s *service, c *gin.Context, recipeID int) (int, error) {
	// Source
	form, err := c.MultipartForm()
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("get form err: %s", err.Error())
	}

	files := form.File["file"]
	if files == nil || len(files) < 1 {
		return http.StatusBadRequest, fmt.Errorf("file not found, %s", err)
	}

	// Destination
	tempDirPath := fmt.Sprintf("/tmp/uploads/%s", uuid.New().String())
	if _, err = os.Stat(tempDirPath); os.IsNotExist(err) {
		os.MkdirAll(tempDirPath, 0700)
	}

	file := files[0] // only one file expected
	src, err := file.Open()
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to open file, %s", err)
	}

	//Checking file-type, supported types are PNG and JPEG
	fileType := file.Header.Get("Content-Type")
	if fileType != PNG && fileType != JPEG {
		return http.StatusBadRequest, fmt.Errorf("unsupported file type, %s", fileType)
	}

	fileExt := path.Ext(file.Filename)

	fileNameUuid := uuid.New()
	filename := fmt.Sprintf("%s/%s%s", tempDirPath, fileNameUuid.String(), fileExt)
	dst, err := os.Create(filename)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to create file, %s", err)
	}
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to copy file, %s", err)
	}

	// Upload file to S3
	uploadPath := fmt.Sprintf("%s/%s/%s", config.Get().AwsFolderName, fileNameUuid.String(), fmt.Sprintf("%s%s", fileNameUuid.String(), fileExt))
	url, err := aws.UploadFileToS3(filename, uploadPath)
	if err != nil || url == "" {
		return http.StatusInternalServerError, fmt.Errorf("s3 upload error: %s", err.Error())
	}

	// Get and upload thumbnail
	thumbnailUrl, err := getAndUploadThumbnail(filename, fileType, tempDirPath, fileNameUuid.String())
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to process thumbnail for %s, %s", filename, err)
	}

	// Create media entry in the database
	mediaEntry := &entity.Media{
		RecipeID:   recipeID,
		MediaType:  strings.Split(fileType, "/")[0],
		MediaURL:   url,
		MediaThumb: thumbnailUrl,
	}

	err = s.repo.Create(mediaEntry)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error creating media entry in db %s", err)
	}

	defer src.Close()
	defer dst.Close()

	return http.StatusOK, nil
}

func getAndUploadThumbnail(filename, fileType, tempDirPath, fileNameUuid string) (string, error) {
	imgFile, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open image file: %s", err)
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return "", fmt.Errorf("failed to decode image file: %s", err)
	}

	// Create thumbnail
	thumbnail := imaging.Thumbnail(img, 200, 200, imaging.CatmullRom)

	// Save thumbnail to a temporary file
	thumbnailFilename := fmt.Sprintf("%s/%s_thumb%s", tempDirPath, fileNameUuid, filepath.Ext(filename))
	thumbnailFile, err := os.Create(thumbnailFilename)
	if err != nil {
		return "", fmt.Errorf("failed to create thumbnail file: %s", err)
	}
	defer thumbnailFile.Close()

	// Encode thumbnail and write it to the file
	switch fileType {
	case PNG:
		err = png.Encode(thumbnailFile, thumbnail)
	case JPEG:
		err = jpeg.Encode(thumbnailFile, thumbnail, &jpeg.Options{Quality: 80})
	default:
		return "", fmt.Errorf("unsupported file type for thumbnail: %s", fileType)
	}
	if err != nil {
		return "", fmt.Errorf("failed to encode thumbnail: %s", err)
	}

	// Upload thumbnail to S3
	thumbnailUploadPath := fmt.Sprintf("%s/%s/%s_thumb%s", config.Get().AwsFolderName, fileNameUuid, fileNameUuid, filepath.Ext(filename))
	thumbnailUrl, err := aws.UploadFileToS3(thumbnailFilename, thumbnailUploadPath)
	if err != nil || thumbnailUrl == "" {
		return "", fmt.Errorf("s3 upload error for thumbnail: %s", err.Error())
	}

	return thumbnailUrl, nil
}
