package entity

import (
	"time"
)

// TableName returns the name of the table corresponding to the HealthService entity in the database.
func (*HealthService) TableName() string {
	return "health_services"
}

// HealthService represents a struct for health services
type HealthService struct {
	ID        int64     `gorm:"Column:id;PRIMARY_KEY" json:"id"`
	Name      string    `gorm:"Column:name" binding:"required" json:"name"`
	UpdatedAt time.Time `gorm:"Column:updated_at" sql:"DEFAULT:current_timestamp" json:"updated_at"`
}

// RequestCreateHealthService represents a struct for creating health services
type RequestCreateHealthService struct {
	Name string `binding:"required" json:"name"`
}
