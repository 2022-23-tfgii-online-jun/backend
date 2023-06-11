package answer

import (
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// answerHandler type contains an instance of AnswerService.
type answerHandler struct {
	answerService ports.AnswerService
}

// newHandler is a constructor function for initializing answerHandler with the given AnswerService.
// The return is a pointer to an answerHandler instance.
func newHandler(answerService ports.AnswerService) *answerHandler {
	return &answerHandler{
		answerService: answerService,
	}
}

// CreateAnswer handles the HTTP request for creating an answer.
// It validates the userUUID and questionUUID from the request parameters and binds the JSON request body to createReq.
// If any error occurs during this process, it will return a 400 Bad Request status.
// If the Answer is created successfully, it will return a 200 OK status.
func (a *answerHandler) CreateAnswer(c *gin.Context) {
	// Parse the userUUID from the request context
	userUUID, err := uuid.Parse(c.GetString("userUUID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid user UUID",
		})
		return
	}

	// Parse the questionUUID from the request parameters
	questionUUIDStr := c.Param("question_uuid")
	questionUUID, err := uuid.Parse(questionUUIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid question UUID",
		})
		return
	}

	// Bind the JSON request body to createReq
	var createReq entity.RequestCreateAnswer
	if err := c.ShouldBindJSON(&createReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid request body",
		})
		return
	}

	// Call the answerService to create the answer
	statusCode, err := a.answerService.CreateAnswer(c, userUUID, questionUUID, &createReq)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"code":    statusCode,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Answer created successfully",
	})
}
