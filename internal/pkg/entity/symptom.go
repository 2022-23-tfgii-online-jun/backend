package entity

import (
	"time"

	"github.com/google/uuid"
)

// TableName returns the name of the table corresponding to the Symptom entity in the database.
func (*Symptom) TableName() string {
	return "symptoms"
}

// Symptom represents a struct for symptoms
type Symptom struct {
	ID        int       `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	UUID      uuid.UUID `gorm:"Column:uuid" json:"uuid"`
	Name      string    `gorm:"Column:name" binding:"required" json:"name"`
	IsActive  bool      `gorm:"Column:is_active" json:"is_active"`
	Scale     int       `gorm:"Column:scale" json:"scale"`
	CreatedAt time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
}

// RequestCreateSymptom represents a struct for creating symptoms
type RequestCreateSymptom struct {
	Name     string `binding:"required" json:"name"`
	IsActive bool   `json:"is_active"`
	Scale    int    `json:"scale"`
}
