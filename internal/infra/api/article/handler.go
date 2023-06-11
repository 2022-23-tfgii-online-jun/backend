package article

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// articleHandler type contains an instance of ArticleService.
type articleHandler struct {
	articleService ports.ArticleService
}

// newHandler is a constructor function for initializing articleHandler with the given ArticleService.
// The return is a pointer to an articleHandler instance.
func newHandler(articleService ports.ArticleService) *articleHandler {
	return &articleHandler{
		articleService: articleService,
	}
}

// CreateArticle handles the HTTP request for creating an article.
// It binds the incoming form-data payload to the reqCreate struct.
// If any error occurs during this process, it will return a 400 Bad Request status.
// If the article is created successfully, it will return a 200 OK status with the created article.
func (a *articleHandler) CreateArticle(c *gin.Context) {
	reqCreate := &entity.RequestCreateArticle{}

	// Get user UUID from context
	userUUID, _ := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))

	// Bind incoming form-data payload to the reqCreate struct.
	if err := c.ShouldBind(reqCreate); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Create the article and store it in the database.
	createdArticle, err := a.articleService.CreateArticle(c, userUUID, reqCreate)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while creating the article", err)
		return
	}

	// Return a successful response with the created article.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Article created successfully",
		"data":    createdArticle,
	})
}

// GetAllArticles handles the HTTP request for getting all articles.
// It retrieves all articles from the database.
// If any error occurs during this process, it will return a 500 Internal Server Error status.
// If the articles are retrieved successfully, it will return a 200 OK status with the retrieved articles.
func (a *articleHandler) GetAllArticles(c *gin.Context) {
	// Get all articles from the database.
	articles, err := a.articleService.GetAllArticles()
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while getting the articles", err)
		return
	}

	// Return a successful response with the retrieved articles.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Articles retrieved successfully",
		"data":    articles,
	})
}

// UpdateArticle handles the HTTP request for updating an article.
// It parses the article UUID from the URL parameter and binds the incoming JSON payload to the reqUpdate struct.
// If any error occurs during this process, it will return a 400 Bad Request status.
// If the article is updated successfully, it will return a 200 OK status with the updated article.
func (a *articleHandler) UpdateArticle(c *gin.Context) {
	// Parse the article UUID from the URL parameter.
	articleUUID, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid UUID format", err)
		return
	}

	// Bind the incoming JSON payload to an UpdateArticle struct.
	reqUpdate := &entity.RequestUpdateArticle{}
	if err := c.ShouldBind(reqUpdate); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Update the article in the database.
	updatedArticle, err := a.articleService.UpdateArticle(articleUUID, reqUpdate)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while updating the article", err)
		return
	}

	// Return a successful response with the updated article.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Article updated successfully",
		"data":    updatedArticle,
	})
}

// DeleteArticle handles the HTTP request for deleting an article.
// It gets the article UUID from the path parameter and calls the articleService to delete the article.
// If any error occurs during this process, it will return the corresponding status code and error message.
// If the article is deleted successfully, it will return a 200 OK status.
func (a *articleHandler) DeleteArticle(c *gin.Context) {
	// Get article UUID from path parameter.
	articleUUID, _ := uuid.Parse(fmt.Sprintf("%v", c.Param("uuid")))

	// Delete the article from the database.
	statusCode, err := a.articleService.DeleteArticle(c, articleUUID)
	if err != nil {
		handleError(c, statusCode, "An error occurred while deleting the article", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Article deleted successfully",
	})
}

// AddArticleToCategory handles the HTTP request for adding an article to a category.
// It binds the incoming JSON payload to the req struct and calls the articleService to add the article to the category.
// If any error occurs during this process, it will return the corresponding status code and error message.
// If the article is added to the category successfully, it will return a 200 OK status.
func (a *articleHandler) AddArticleToCategory(ctx *gin.Context) {
	// Declare a variable for the incoming request payload.
	req := &entity.AddArticleToCategoryRequest{}

	// Bind incoming JSON payload to the req struct.
	if err := ctx.ShouldBindJSON(req); err != nil {
		handleError(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Call the service to add the article to the category.
	err := a.articleService.AddArticleToCategory(req.Category, req.Article)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "An error occurred while adding the article to the category", err)
		return
	}

	// Return a successful response.
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Article added to Category successfully",
	})
}

// handleError is a generic error handler that logs the error and responds with the corresponding status code and error message.
func handleError(c *gin.Context, statusCode int, message string, err error) {
	// Log the error message and the error itself.
	log.Printf("[ArticleHandler]: %s, %v", message, err)

	// Send the JSON response with the status code and error message.
	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": message,
		"data":    nil,
	})
}
