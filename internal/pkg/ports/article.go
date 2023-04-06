package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ArticleRepository is the interface that defines the methods for accessing the article data store.
type ArticleRepository interface {
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	// CreateWithOmit creates a new user record while omitting specific fields.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// UpdateUser updates an existing user record with the provided user data.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first record that matches the given conditions from the database
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error

	//Delete(article *entity.Article) error
	//Find(articles *[]entity.Article, query interface{}, args ...interface{}) error
}

// ArticleService is the interface that defines the methods for managing articles in the application.
type ArticleService interface {
	CreateArticle(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateArticle) (int, error)
	//UpdateArticle(updateReq *entity.UpdateArticle) error
	//DeleteArticle(articleID uint) error
	//GetArticleByID(articleID uint) (*entity.Article, error)
	//GetAllArticles() ([]entity.Article, error)
}
