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

type articleHandler struct {
	articleService ports.ArticleService
}

func newHandler(articleService ports.ArticleService) *articleHandler {
	return &articleHandler{
		articleService: articleService,
	}
}

// CreateArticle handler for creating an article
func (a *articleHandler) CreateArticle(c *gin.Context) {
	reqCreate := &entity.RequestCreateArticle{}

	//  Get user uuid from context
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

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Article created successfully",
		"data":    createdArticle,
	})
}

// DeleteArticle handler for deleting an article
func (a *articleHandler) DeleteArticle(c *gin.Context) {
	// Get article uuid from path parameter
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

// handleError is a generic error handler that logs the error and responds
func handleError(c *gin.Context, statusCode int, message string, err error) {
	// Log the error message and the error itself
	log.Printf("[ArticleHandler]: %s, %v", message, err)

	// Send the JSON response with the status code and error message
	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": message,
		"data":    nil,
	})
}
