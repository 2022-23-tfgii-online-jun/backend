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

// questionHandler type contains an instance of QuestionService
type questionHandler struct {
	questionService ports.QuestionService
}

// newHandler is a constructor function for initializing questionHandler with the given QuestionService.
// The return is a pointer to an questionHandler instance.
func newHandler(questionService ports.QuestionService) *questionHandler {
	return &questionHandler{
		questionService: questionService,
	}
}

// CreateQuestion handles the HTTP request for creating a question.
// It binds the incoming form-data payload to the reqCreate struct and calls the question service to create the question.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the question is created successfully, it returns a 200 OK status with the created question.
func (q *questionHandler) CreateQuestion(c *gin.Context) {
	reqCreate := &entity.RequestCreateQuestion{}

	// Get user UUID from context
	userUUID, _ := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))

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
// It retrieves all questions from the database.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the questions are retrieved successfully, it returns a 200 OK status with the retrieved questions.
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

// handleError is a generic error handler that logs the error and responds.
func handleError(c *gin.Context, statusCode int, message string, err error) {
	// Log the error message and the error itself.
	log.Printf("[QuestionHandler]: %s, %v", message, err)

	// Send the JSON response with the status code and error message.
	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": message,
		"data":    nil,
	})
}
