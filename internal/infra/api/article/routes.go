package article

import (
	"github.com/emur-uy/backend/internal/infra/api/middlewares"
	"github.com/emur-uy/backend/internal/infra/api/middlewares/constants"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/article"
	"github.com/emur-uy/backend/internal/pkg/service/media"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the article-related routes on the given gin.Engine instance.
// It initializes the necessary components, such as the repository, service, and handler,
// to handle article-related operations in a hexagonal architecture.
func RegisterRoutes(e *gin.Engine) {
	// Initialize the repository by creating a new PostgreSQL client.
	client := postgresql.NewClient()

	// Create new repository instances for each repository interface
	articleRepo := postgresql.NewArticleRepository(client)
	mediaRepo := postgresql.NewMediaRepository(client)
	articleMediaRepo := postgresql.NewArticleMediaRepository(client)

	// Create new services
	mediaService := media.NewService(mediaRepo)
	articleMediaService := article.NewArticleMediaService(articleMediaRepo)
	articleService := article.NewService(articleRepo, mediaService, articleMediaService)

	// Create a new articleHandler instance by injecting the ArticleService.
	handler := newHandler(articleService)

	// Group the article routes together.
	articleRoutes := e.Group("/api/v1/articles")

	// Register admin routes requiring authentication and authorization for admin role.
	adminRoutes := articleRoutes.Group("", middlewares.Authenticate(), middlewares.Authorize(constants.RoleAdmin))
	adminRoutes.POST("", handler.CreateArticle)
	adminRoutes.DELETE("/:uuid", handler.DeleteArticle)
	adminRoutes.PUT("/:uuid", handler.UpdateArticle)
	adminRoutes.POST("/:uuid/categories", handler.AddArticleToCategory)

	// Register route for getting all articles accessible to both admin and user roles.
	allowedRoles := []string{constants.RoleAdmin, constants.RoleUser}
	articleRoutes.GET("", middlewares.Authenticate(), middlewares.Authorize(allowedRoles...), handler.GetAllArticles)
}
