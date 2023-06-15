package entity

import (
	"time"

	"github.com/google/uuid"
)

// TableName returns the name of the table corresponding to the Monitoring entity in the database.
func (*Monitoring) TableName() string {
	return "monitorings"
}

// Monitoring represents a struct for the monitorings table
type Monitoring struct {
	ID        int64     `gorm:"column:id;primary_key" json:"-"`
	UserID    int       `gorm:"column:user_id" json:"-"`
	SymptomID int       `gorm:"column:symptom_id" json:"symptom"`
	Scale     int       `gorm:"Column:scale" json:"scale"`
	Date      time.Time `gorm:"column:date;default:current_timestamp" json:"date"`
}

type RequestCreateMonitoring struct {
	SymptomUUID uuid.UUID `json:"symptom"`
	Scale       int       `json:"scale"`
}
