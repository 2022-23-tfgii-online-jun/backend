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
	ErrTypeAssertionFailed  = errors.New("type assertion failed")
	ErrCreatingArticle      = errors.New("error creating article")
	ErrUpdatingArticle      = errors.New("error updating article")
	ErrDeletingArticle      = errors.New("error deleting article")
	ErrCreatingMedia        = errors.New("error creating media")
	ErrFindingArticleMedia  = errors.New("error finding article media")
	ErrCreatingArticleMedia = errors.New("error creating article media association")
	ErrDeletingArticleMedia = errors.New("error deleting article media association")
	ErrDeletingMedia        = errors.New("error deleting media")
	ErrUnsupportedFileType  = errors.New("unsupported file type")
	ErrAddingCategory       = errors.New("error adding article to category")
	ErrFileNotFound         = errors.New("file not found")
)

const (
	PNG  = "image/png"
	JPEG = "image/jpeg"
)

// service struct holds the necessary dependencies for the article service
type service struct {
	repo                ports.ArticleRepository
	mediaService        ports.MediaService
	articleMediaService ports.ArticleMediaService
}

// NewService returns a new instance of the article service with the given article repository, media service, and article media service.
func NewService(articleRepo ports.ArticleRepository, mediaService ports.MediaService, articleMediaService ports.ArticleMediaService) ports.ArticleService {
	return &service{
		repo:                articleRepo,
		mediaService:        mediaService,
		articleMediaService: articleMediaService,
	}
}

// CreateArticle is the service for creating an article and saving it in the database.
func (s *service) CreateArticle(c *gin.Context, createReq *entity.RequestCreateArticle) (*entity.Article, error) {
	// Call the processUploadRequestFiles function to handle the image upload and create the media entry
	fileProcessCode, fileUrls, err := processUploadRequestFiles(s, c)
	if err != nil || fileProcessCode != http.StatusOK {
		return nil, fmt.Errorf("error processing content upload file: %s", err)
	}

	// Create a new article
	article := &entity.Article{
		Title:   createReq.Title,
		Content: createReq.Content,
	}

	// Save the article to the database
	err = s.repo.CreateWithOmit("uuid", article)
	if err != nil {
		return nil, ErrCreatingArticle
	}

	// For each uploaded file, create a new media entry and then a new ArticleMedia entry
	for _, fileUrl := range fileUrls {
		media := &entity.Media{
			MediaURL: fileUrl,
		}

		err = s.mediaService.CreateMedia(media)
		if err != nil {
			return nil, ErrCreatingMedia
		}

		articleMedia := &entity.ArticleMedia{
			ArticleID: article.ID,
			MediaID:   media.ID,
		}
		err = s.articleMediaService.CreateArticleMedia(articleMedia)
		if err != nil {
			return nil, fmt.Errorf("error creating ArticleMedia: %s", err)
		}
	}

	return article, nil
}

// UpdateArticle is the service for updating an article in the database.
func (s *service) UpdateArticle(c *gin.Context, articleUUID uuid.UUID, updateReq *entity.RequestUpdateArticle) (int, error) {
	// Find the existing article by UUID
	article := &entity.Article{}
	foundArticle, err := s.repo.FindByUUID(articleUUID, article)
	if err != nil {
		// Return error if the article is not found
		return http.StatusNotFound, err
	}

	// Perform type assertion to convert foundArticle to *entity.Article
	article, ok := foundArticle.(*entity.Article)
	if !ok {
		return http.StatusInternalServerError, ErrTypeAssertionFailed
	}

	if updateReq == nil {
		return http.StatusBadRequest, errors.New("nil payload")
	}

	// Update the article fields with the new data from the update request
	article.Title = updateReq.Title
	article.Content = updateReq.Content
	article.Content = updateReq.Content

	// Update the article in the database
	err = s.repo.Update(article)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error updating article: %s", err)
	}

	fileProcessCode, fileUrls, err := processUploadRequestFiles(s, c)
	if err != nil || fileProcessCode != http.StatusOK {
		return http.StatusInternalServerError, fmt.Errorf("error processing content upload file: %s", err)
	}

	// Get existing article media data
	articleMedias := []*entity.ArticleMedia{}
	err = s.articleMediaService.FindByArticleID(article.ID, &articleMedias)
	if err != nil {
		return http.StatusInternalServerError, ErrFindingArticleMedia
	}

	// For each uploaded file, create a new media entry and a new article_media association
	for _, fileUrl := range fileUrls {
		media := &entity.Media{
			MediaURL: fileUrl,
		}
		err = s.mediaService.CreateMedia(media)
		if err != nil {
			return http.StatusInternalServerError, ErrCreatingMedia
		}
		articleMedia := &entity.ArticleMedia{
			ArticleID: article.ID,
			MediaID:   media.ID,
		}
		err = s.articleMediaService.CreateArticleMedia(articleMedia)
		if err != nil {
			return http.StatusInternalServerError, ErrCreatingArticleMedia
		}
	}

	// Delete old media entries
	for _, articleMedia := range articleMedias {
		mediaID := articleMedia.MediaID

		// Delete the articleMedia entry
		err = s.articleMediaService.DeleteArticleMedia(articleMedia)
		if err != nil {
			return http.StatusInternalServerError, ErrDeletingArticleMedia
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

// DeleteArticle deletes an article from the database by its UUID.
func (s *service) DeleteArticle(c *gin.Context, articleUUID uuid.UUID) (int, error) {
	// Retrieve the article from the repository by its UUID.
	article := &entity.Article{}
	foundArticle, err := s.repo.FindByUUID(articleUUID, article)
	if err != nil {
		// Return an error response if the article is not found.
		return http.StatusNotFound, err
	}

	// Perform type assertion to convert foundArticle to *entity.Article.
	article, ok := foundArticle.(*entity.Article)
	if !ok {
		return http.StatusInternalServerError, ErrTypeAssertionFailed
	}

	// Get article media associations by article ID
	articleMedias := []*entity.ArticleMedia{}
	err = s.articleMediaService.FindByArticleID(article.ID, &articleMedias)
	if err != nil {
		return http.StatusInternalServerError, ErrFindingArticleMedia
	}

	// Iterate over each article_media association
	for _, articleMedia := range articleMedias {
		media := &entity.Media{}
		// Find the media by ID
		err := s.mediaService.FindByMediaID(articleMedia.MediaID, media)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		// Delete the article_media association from the repository
		err = s.articleMediaService.DeleteArticleMedia(articleMedia)
		if err != nil {
			return http.StatusInternalServerError, ErrDeletingArticleMedia
		}

		// Delete the media from the repository
		err = s.mediaService.DeleteMedia(media)
		if err != nil {
			return http.StatusInternalServerError, ErrDeletingMedia
		}
	}

	// Delete the article from the repository
	err = s.repo.Delete(article)
	if err != nil {
		return http.StatusInternalServerError, ErrDeletingArticle
	}

	// Return a success response.
	return http.StatusOK, nil
}

// GetAllArticles returns all articles stored in the database with associated image URLs.
func (s *service) GetAllArticles() ([]*entity.ArticleWithMediaURLs, error) {
	// Get all articles from the database
	var articles []*entity.Article
	if err := s.repo.Find(&articles); err != nil {
		return nil, err
	}

	articlesWithMediaURLs := make([]*entity.ArticleWithMediaURLs, len(articles))

	for i, article := range articles {
		// Get associated media for the article
		articleMedias := []*entity.ArticleMedia{}
		err := s.articleMediaService.FindByArticleID(article.ID, &articleMedias)
		if err != nil {
			return nil, err
		}

		mediaURLs := make([]string, len(articleMedias))

		// Get media URLs
		for j, articleMedia := range articleMedias {
			media := &entity.Media{}
			err = s.mediaService.FindByMediaID(articleMedia.MediaID, media)
			if err != nil {
				return nil, err
			}

			mediaURLs[j] = media.MediaURL
		}

		articlesWithMediaURLs[i] = &entity.ArticleWithMediaURLs{
			Article:   article,
			MediaURLs: mediaURLs,
		}
	}

	return articlesWithMediaURLs, nil
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
