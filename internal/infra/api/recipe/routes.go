package recipe

import (
	"github.com/emur-uy/backend/internal/infra/api/middlewares"
	"github.com/emur-uy/backend/internal/infra/api/middlewares/constants"
	"github.com/emur-uy/backend/internal/infra/repositories/postgresql"
	"github.com/emur-uy/backend/internal/pkg/service/media"
	"github.com/emur-uy/backend/internal/pkg/service/recipe"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the recipe-related routes on the given gin.Engine instance.
// It initializes the necessary components, such as the repository, service, and handler,
// to handle recipe-related operations in a hexagonal architecture.
func RegisterRoutes(e *gin.Engine) {
	// Initialize the repository by creating a new PostgreSQL client.
	client := postgresql.NewClient()

	// Create new repository instances for each repository interface
	recipeRepo := postgresql.NewRecipeRepository(client)
	mediaRepo := postgresql.NewMediaRepository(client)
	recipeMediaRepo := postgresql.NewRecipeMediaRepository(client)

	// Create new services
	mediaService := media.NewService(mediaRepo)
	recipeMediaService := recipe.NewRecipeMediaService(recipeMediaRepo)
	recipeService := recipe.NewService(recipeRepo, mediaService, recipeMediaService)

	// Create a new recipeHandler instance by injecting the RecipeService.
	handler := newHandler(recipeService)

	// Group the recipe routes together.
	recipeRoutes := e.Group("/api/v1/recipes")

	// Register admin routes requiring authentication and authorization for admin role.
	adminRoutes := recipeRoutes.Group("", middlewares.Authenticate(), middlewares.Authorize(constants.RoleAdmin))
	adminRoutes.POST("", handler.CreateRecipe)
	adminRoutes.DELETE("/:uuid", handler.DeleteRecipe)
	adminRoutes.PUT("/:uuid", handler.UpdateRecipe)

	// Register route for getting all recipes accessible to both admin and user roles.
	allowedRoles := []string{constants.RoleAdmin, constants.RoleUser}
	recipeRoutes.GET("", middlewares.Authenticate(), middlewares.Authorize(allowedRoles...), handler.GetAllRecipes)
	recipeRoutes.POST("/:uuid/vote", middlewares.Authenticate(), middlewares.Authorize(allowedRoles...), handler.VoteRecipe)
}
