package article

import (
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

type service struct {
	repo ports.ArticleRepository
}

// NewService returns a new instance of the article service with the given article repository.
func NewService(articleRepo ports.ArticleRepository) ports.ArticleService {
	return &service{
		repo: articleRepo,
	}
}

// CreateArticle is the service for creating an article and saving it in the database
func (s *service) CreateArticle(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateArticle) (int, error) {
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

	fileProcessCode, fileUrl, err := processUploadRequestFile(s, c)
	if err != nil || fileProcessCode != http.StatusOK {
		return http.StatusInternalServerError, fmt.Errorf("error processing content upload file, %s", err)
	}

	// Create a new article
	article := &entity.Article{
		Title:   createReq.Title,
		Image:   fileUrl,
		Content: createReq.Content,
		UserID:  user.ID,
	}

	// Save the article to the database
	err = s.repo.CreateWithOmit("uuid", article)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error creating article: %s", err)
	}

	// Return the HTTP OK status code if the update is successful
	return http.StatusOK, nil
}

// DeleteArticle deletes an article from the database by its UUID.
func (s *service) DeleteArticle(c *gin.Context, articleUUID uuid.UUID) (int, error) {
	// Retrieve the article from the repository by its UUID.
	article := &entity.Article{}
	foundArticle, err := s.repo.FindByUUID(articleUUID, article)
	if err != nil {
		// Return an error response if the article is not found.
		return http.StatusNotFound, fmt.Errorf("article not found")
	}

	// Perform type assertion to convert foundArticle to *entity.Article.
	article, ok := foundArticle.(*entity.Article)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("type assertion failed")
	}

	// Delete the article from the repository.
	err = s.repo.Delete(article)
	if err != nil {
		// Return an error response if there was an issue deleting the article.
		return http.StatusInternalServerError, fmt.Errorf("failed to delete article")
	}

	// Return a success response.
	return http.StatusOK, nil
}

// processUploadRequestFile processes the file upload request
func processUploadRequestFile(s *service, c *gin.Context) (int, string, error) {
	// Parse the multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return http.StatusBadRequest, "", fmt.Errorf("get form err: %s", err.Error())
	}
	files := form.File["file"]
	if files == nil || len(files) < 1 {
		return http.StatusBadRequest, "", fmt.Errorf("file not found, %s", err)
	}

	// Access the first file
	file := files[0]

	// Open the source file
	src, err := file.Open()
	if err != nil {
		return http.StatusInternalServerError, "", fmt.Errorf("failed to open file, %s", err)
	}

	// Validate the file type (PNG and JPEG are supported)
	fileType := file.Header.Get("Content-Type")
	if fileType != "image/png" && fileType != "image/jpeg" {
		return http.StatusBadRequest, "", fmt.Errorf("unsupported file type, %s", fileType)
	}

	// Generate a unique filename with UUID and proper extension
	fileExt := path.Ext(file.Filename)
	fileNameUuid := uuid.New()

	// Upload the image to S3 directly from the source stream
	uploadPath := fmt.Sprintf("%s/%s", config.Get().AwsFolderName, fmt.Sprintf("%s%s", fileNameUuid.String(), fileExt))
	url, err := aws.UploadFileToS3Stream(src, uploadPath, true)
	if err != nil || url == "" {
		return http.StatusInternalServerError, "", fmt.Errorf("s3 upload error: %s", err.Error())
	}

	// Close the source file
	defer src.Close()

	return http.StatusOK, url, nil
}
