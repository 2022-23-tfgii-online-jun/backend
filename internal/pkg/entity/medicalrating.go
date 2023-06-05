// Package entity defines the domain entities (models) for the application.
package entity

import (
	"time"
)

// TableName returns the name of the table corresponding to the MedicalRating entity in the database.
func (*MedicalRating) TableName() string {
	return "medical_ratings"
}

// MedicalRating represents a struct for medical rating records
type MedicalRating struct {
	ID         int64     `gorm:"Column:id;PRIMARY_KEY" json:"id"`
	MedicalID  int64     `gorm:"Column:medical_id" json:"medical_id"`
	ReminderID int64     `gorm:"Column:reminder_id" json:"reminder_id"`
	Rating     int64     `gorm:"Column:rating" json:"rating"`
	CreatedAt  time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"-"`
}
