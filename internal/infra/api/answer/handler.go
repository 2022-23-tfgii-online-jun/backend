package answer

import (
	"fmt"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type answerHandler struct {
	answerService ports.AnswerService
}

func newHandler(answerService ports.AnswerService) *answerHandler {
	return &answerHandler{
		answerService: answerService,
	}
}

// CreateAnswer handles the HTTP request for creating an answer.
func (a *answerHandler) CreateAnswer(c *gin.Context) {
	userUUID, err := uuid.Parse(c.GetString("userUUID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid user UUID",
		})
		return
	}

	questionUUIDStr := c.Param("question_uuid")
	questionUUID, err := uuid.Parse(questionUUIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid question UUID",
		})
		return
	}

	var createReq entity.RequestCreateAnswer
	if err := c.ShouldBindJSON(&createReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid request body",
		})
		return
	}

	fmt.Println(questionUUID)

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
