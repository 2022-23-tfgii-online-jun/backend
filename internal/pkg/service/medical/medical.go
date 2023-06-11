package medical

import (
	"encoding/csv"
	"errors"
	"net/http"
	"strings"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/charmap"
)

var (
	ErrGettingFile    = errors.New("error getting the csv file from the request")
	ErrReadingFile    = errors.New("error reading the csv file")
	ErrCreatingRecord = errors.New("error creating record")
	ErrIDsRequired    = errors.New("medical and reminder IDs are required")
	ErrAddingRating   = errors.New("error adding rating to medical record")
)

// service struct holds the necessary dependencies for the medical service
type service struct {
	repo ports.MedicalRepository
}

// NewService returns a new instance of the medical service with the given medical repository.
func NewService(medicalRepo ports.MedicalRepository) ports.MedicalService {
	return &service{
		repo: medicalRepo,
	}
}

// CreateRecordFromFile is the service for creating a medical record and saving it in the database
func (s *service) CreateRecordFromFile(c *gin.Context) (int, error) {
	// Get the file from the request
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		return http.StatusInternalServerError, ErrGettingFile
	}

	// Create a CSV reader with the appropriate encoding
	reader := csv.NewReader(charmap.ISO8859_1.NewDecoder().Reader(file))
	reader.Comma = ';' // Set the delimiter if necessary
	// Read the records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		return http.StatusInternalServerError, ErrReadingFile
	}

	// Skip the header row
	dataRows := records[1:]

	for _, row := range dataRows {
		// Create a new medical record for each row in the CSV
		record := &entity.Medical{
			FirstName:        strings.TrimSpace(row[0]),
			LastName:         strings.TrimSpace(row[2]),
			CjppuNumber:      strings.TrimSpace(row[4]),
			ProfessionNumber: strings.TrimSpace(row[5]),
		}

		// Save the record to the database
		err = s.repo.Create(record)
		if err != nil {
			return http.StatusInternalServerError, ErrCreatingRecord
		}
	}

	// Return the HTTP OK status code if the update is successful
	return http.StatusOK, nil
}

// GetAllMedicalRecords returns all medical records stored in the database
func (s *service) GetAllMedicalRecords() ([]*entity.Medical, error) {
	// Get all medical records from the database
	var medicals []*entity.Medical
	if err := s.repo.Find(&medicals); err != nil {
		return nil, err
	}

	return medicals, nil
}

// AddRatingToMedical is the service for adding a rating to a medical record.
func (m *service) AddRatingToMedical(rating *entity.MedicalRating) (int, error) {
	// Validate the input parameters
	if rating.MedicalID == 0 || rating.ReminderID == 0 {
		return http.StatusBadRequest, ErrIDsRequired
	}

	// Save the medical rating to the database
	err := m.repo.Create(rating)
	if err != nil {
		return http.StatusInternalServerError, ErrAddingRating
	}

	// Return the HTTP OK status code if the operation is successful
	return http.StatusOK, nil
}
