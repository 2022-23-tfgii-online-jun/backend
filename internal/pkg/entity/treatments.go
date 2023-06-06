// Package entity defines the domain entities (models) for the application.
package entity

import (
	"time"

	"github.com/google/uuid"
)

// TableName returns the name of the table corresponding to the Treatment entity in the database.
func (*Treatment) TableName() string {
	return "treatments"
}

// Treatment represents a struct for treatments
type Treatment struct {
	ID        int       `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	UUID      uuid.UUID `gorm:"Column:uuid" json:"uuid"`
	UserID    int       `gorm:"Column:user_id" json:"-"`
	Name      string    `gorm:"Column:name" binding:"required" json:"name"`
	Type      string    `gorm:"Column:type" binding:"required" json:"type"`
	Frequency string    `gorm:"Column:frecuency" binding:"required" json:"frequency"`
	Shots     string    `gorm:"Column:shots" binding:"required" json:"shots"`
	DateStart time.Time `gorm:"Column:date_start" binding:"required" json:"date_start"`
	CreatedAt time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
}

// RequestCreateTreatment represents a struct for creating treatments
type RequestCreateTreatment struct {
	Name      string    `form:"name" binding:"required"`
	Type      string    `form:"type" binding:"required"`
	Frequency string    `form:"frequency" binding:"required"`
	Shots     string    `form:"shots" binding:"required"`
	DateStart time.Time `form:"date_start" binding:"required"`
}
