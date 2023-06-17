package article

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
	ErrTypeAssertion    = errors.New("type assertion failed")
	ErrCreatingArticle  = errors.New("error creating article")
	ErrUpdatingArticle  = errors.New("error updating article")
	ErrDeletingArticle  = errors.New("failed to delete article")
	ErrAddingCategory   = errors.New("error adding article to category")
	ErrProcessingUpload = errors.New("error processing content upload file")
)

// service defines dependencies for HTTP server.
type service struct {
	repo ports.ArticleRepository
}

// NewService creates and returns a new Article service instance
func NewService(articleRepo ports.ArticleRepository) ports.ArticleService {
	return &service{
		repo: articleRepo,
	}
}

// CreateArticle is the service for creating an article and saving it in the database
func (s *service) CreateArticle(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateArticle) (int, error) {
	user := &entity.User{}

	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		return http.StatusNotFound, err
	}

	user, ok := foundUser.(*entity.User)
	if !ok {
		return http.StatusInternalServerError, ErrTypeAssertion
	}

	fileProcessCode, fileUrl, err := processUploadRequestFile(s, c)
	if err != nil || fileProcessCode != http.StatusOK {
		return http.StatusInternalServerError, ErrProcessingUpload
	}

	article := &entity.Article{
		Title:   createReq.Title,
		Image:   fileUrl,
		Content: createReq.Content,
		UserID:  user.ID,
	}

	err = s.repo.CreateWithOmit("uuid", article)
	if err != nil {
		return http.StatusInternalServerError, ErrCreatingArticle
	}

	return http.StatusOK, nil
}

// UpdateArticle is the service for updating an article in the database
func (s *service) UpdateArticle(articleUUID uuid.UUID, updateReq *entity.RequestUpdateArticle) (int, error) {
	article := &entity.Article{}
	foundArticle, err := s.repo.FindByUUID(articleUUID, article)
	if err != nil {
		return http.StatusNotFound, err
	}

	article, ok := foundArticle.(*entity.Article)
	if !ok {
		return http.StatusInternalServerError, ErrTypeAssertion
	}

	article.Title = updateReq.Title
	article.Content = updateReq.Content

	err = s.repo.Update(article)
	if err != nil {
		return http.StatusInternalServerError, ErrUpdatingArticle
	}

	return http.StatusOK, nil
}

// DeleteArticle deletes an article from the database by its UUID.
func (s *service) DeleteArticle(c *gin.Context, articleUUID uuid.UUID) (int, error) {
	article := &entity.Article{}
	foundArticle, err := s.repo.FindByUUID(articleUUID, article)
	if err != nil {
		return http.StatusNotFound, err
	}

	article, ok := foundArticle.(*entity.Article)
	if !ok {
		return http.StatusInternalServerError, ErrTypeAssertion
	}

	err = s.repo.Delete(article)
	if err != nil {
		return http.StatusInternalServerError, ErrDeletingArticle
	}

	return http.StatusOK, nil
}

// GetAllArticles returns all articles stored in the database
func (s *service) GetAllArticles() ([]*entity.Article, error) {
	var articles []*entity.Article
	if err := s.repo.Find(&articles); err != nil {
		return nil, err
	}

	return articles, nil
}

// AddArticleToCategory is the service for adding an article to a category and saving the relationship in the database
func (s *service) AddArticleToCategory(categoryUUID uuid.UUID, articleUUID uuid.UUID) error {
	article := &entity.Article{}
	category := &entity.Category{}

	_, err := s.repo.FindByUUID(articleUUID, article)
	if err != nil {
		return err
	}

	_, err = s.repo.FindByUUID(categoryUUID, category)
	if err != nil {
		return err
	}

	articleCategory := entity.ArticleCategory{
		ArticleID:  article.ID,
		CategoryID: category.ID,
	}

	err = s.repo.Create(&articleCategory)
	if err != nil {
		return ErrAddingCategory
	}

	return nil
}

var uploadFunc = aws.UploadFileToS3Stream

// processUploadRequestFile processes the file upload request
func processUploadRequestFile(s *service, c *gin.Context) (int, string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return http.StatusBadRequest, "", ErrProcessingUpload
	}
	files := form.File["file"]
	if files == nil || len(files) < 1 {
		return http.StatusBadRequest, "", ErrProcessingUpload
	}

	file := files[0]
	src, err := file.Open()
	if err != nil {
		return http.StatusInternalServerError, "", ErrProcessingUpload
	}

	fileType := file.Header.Get("Content-Type")
	if fileType != "image/png" && fileType != "image/jpeg" {
		return http.StatusBadRequest, "", ErrProcessingUpload
	}

	fileExt := path.Ext(file.Filename)
	fileNameUuid := uuid.New()

	uploadPath := fmt.Sprintf("%s/%s", config.Get().AwsFolderName, fmt.Sprintf("%s%s", fileNameUuid.String(), fileExt))
	url, err := uploadFunc(src, uploadPath, true)
	if err != nil || url == "" {
		return http.StatusInternalServerError, "", ErrProcessingUpload
	}

	defer src.Close()

	return http.StatusOK, url, nil
}
