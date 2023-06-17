package question

import (
	"errors"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	ErrTypeAssertionFailed = errors.New("type assertion failed")
	ErrCreatingQuestion    = errors.New("error creating question")
)

// service struct holds the necessary dependencies for the question service
type service struct {
	repo ports.QuestionRepository
}

// NewService returns a new instance of the question service with the given question repository.
func NewService(questionRepo ports.QuestionRepository) ports.QuestionService {
	return &service{
		repo: questionRepo,
	}
}

// CreateQuestion is the service for creating a question and saving it in the database.
func (s *service) CreateQuestion(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateQuestion) (*entity.Question, error) {
	user := &entity.User{}

	// Find user by UUID
	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		// Return error if the user is not found
		return nil, err
	}

	// Perform type assertion to convert foundUser to *entity.User
	user, ok := foundUser.(*entity.User)
	if !ok {
		return nil, ErrTypeAssertionFailed
	}

	if createReq == nil {
		return nil, errors.New("nil payload")
	}

	// Create a new question
	question := &entity.Question{
		UserID: user.ID,
		Text:   createReq.Text,
	}

	// Save the question to the database
	err = s.repo.CreateWithOmit("uuid", question)
	if err != nil {
		return nil, ErrCreatingQuestion
	}

	return question, nil
}

// GetAllQuestions returns all questions stored in the database.
func (s *service) GetAllQuestions() ([]*entity.Question, error) {
	// Get all questions from the database
	var questions []*entity.Question
	if err := s.repo.Find(&questions); err != nil {
		return nil, err
	}

	return questions, nil
}

func (s *service) GetAllQuestionsAndAnswers() ([]*entity.QuestionAndAnswers, error) {
	// Get all questions from the database
	var questions []*entity.QuestionAndAnswers
	if err := s.repo.Find(&questions); err != nil {
		return nil, err
	}

	// Get answers for each question
	for _, question := range questions {
		var answers []*entity.Answer
		if err := s.repo.Find(&answers, question.ID); err != nil {
			return nil, err
		}
		question.Answers = answers
	}

	return questions, nil
}
