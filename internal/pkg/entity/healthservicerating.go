package entity

import (
	"time"
)

// TableName returns the name of the table corresponding to the HealthService entity in the database.
func (*HealthServiceRating) TableName() string {
	return "health_services_ratings"
}

// HealthServiceRating represents a struct for health service ratings
type HealthServiceRating struct {
	HealthServiceID int       `gorm:"Column:health_service_id" json:"health_service_id"`
	ReminderID      int       `gorm:"Column:reminder_id" json:"reminder_id"`
	Rating          int       `gorm:"Column:rating" json:"rating"`
	CreatedAt       time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"-"`
}
