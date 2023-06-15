package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ArticleRepository is an interface that acts as a contract for the data access layer,
// requiring implementations to provide methods for querying and modifying article data.
type ArticleRepository interface {
	// FindByUUID retrieves an Article based on its UUID.
	// Returns the article and an error if any occurred.
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	// Create takes a new Article value and adds it to the data store.
	// Returns an error if the operation fails.
	Create(value interface{}) error

	// CreateWithOmit creates a new article record while omitting specific fields.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// Update updates an existing Article value in the data store.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first record that matches the given conditions from the database.
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error

	// Find retrieves all Article values that match the given conditions from the database.
	// Returns an error if the operation fails.
	Find(out interface{}, conditions ...interface{}) error

	// Delete removes an existing Article from the data store.
	// Returns an error if the operation fails.
	Delete(out interface{}) error
}

// ArticleService is an interface defining a contract for business logic operators related to Articles.
// It works with the entity layer to manipulate Article data.
type ArticleService interface {
	// CreateArticle takes the user's UUID and a request to create an Article.
	// Returns the status and an error if any occurred.
	CreateArticle(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateArticle) (int, error)

	// UpdateArticle updates an existing Article using the provided update request.
	// Returns the status and an error if any occurred.
	UpdateArticle(articleUUID uuid.UUID, updateReq *entity.RequestUpdateArticle) (int, error)

	// DeleteArticle removes an existing Article using the provided article UUID.
	// Returns the status and an error if any occurred.
	DeleteArticle(c *gin.Context, articleUUID uuid.UUID) (int, error)

	// GetAllArticles retrieves all articles from the data store.
	// Returns a slice of Article entities and an error if any occurred.
	GetAllArticles() ([]*entity.Article, error)

	// AddArticleToCategory associates an article with a category using their UUIDs.
	// Returns an error if the operation fails.
	AddArticleToCategory(articleUUID uuid.UUID, categoryUUID uuid.UUID) error
}
