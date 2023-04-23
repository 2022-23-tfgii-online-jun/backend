package question

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type questionHandler struct {
	questionService ports.QuestionService
}

func newHandler(questionService ports.QuestionService) *questionHandler {
	return &questionHandler{
		questionService: questionService,
	}
}

// CreateQuestion handler for creating a question
func (q *questionHandler) CreateQuestion(c *gin.Context) {
	reqCreate := &entity.RequestCreateQuestion{}

	//  Get user uuid from context
	userUUID, _ := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))

	fmt.Println(userUUID)

	// Bind incoming form-data payload to the reqCreate struct.
	if err := c.ShouldBind(reqCreate); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Create the question and store it in the database.
	createdQuestion, err := q.questionService.CreateQuestion(c, userUUID, reqCreate)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while creating the question", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Question created successfully",
		"data":    createdQuestion,
	})
}

// GetAllQuestions handles the HTTP request for getting all questions.
func (q *questionHandler) GetAllQuestions(c *gin.Context) {
	// Get all questions from the database.
	questions, err := q.questionService.GetAllQuestions()
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while getting the questions", err)
		return
	}

	// Return a successful response with the retrieved questions.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Questions retrieved successfully",
		"data":    questions,
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
