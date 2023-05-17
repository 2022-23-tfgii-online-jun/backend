package entity

import (
	"time"

	"github.com/google/uuid"
)

// TableName returns the name of the table corresponding to the Reminder entity in the database.
func (*Reminder) TableName() string {
	return "reminders"
}

// Reminder represents a struct for reminders
type Reminder struct {
	UserID       int       `gorm:"Column:user_id" json:"-"`
	FileID       int       `gorm:"Column:file_id" json:"-"`
	UUID         uuid.UUID `gorm:"Column:uuid" json:"uuid"`
	Name         string    `gorm:"Column:name" binding:"required" json:"name"`
	Type         string    `gorm:"Column:type" binding:"required" json:"type"`
	Date         time.Time `gorm:"Column:date" binding:"required" json:"date"`
	Notification int       `gorm:"Column:notification" json:"notification"`
	Task         []Task    `gorm:"Column:task" json:"task"`
	Note         string    `gorm:"Column:note" json:"note"`
	IsActive     bool      `gorm:"Column:is_active" sql:"DEFAULT:1" json:"is_active"`
	CreatedAt    time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
}

// Notification represents the struct for notifications
type Notification struct {
	DaysOrHours string `gorm:"Column:days_or_hours" json:"days_or_hours"`
	HoursBefore int    `gorm:"Column:hours_before" json:"hours_before"`
}

// Task represents the struct for tasks
type Task struct {
	Name    string `gorm:"Column:name" json:"name"`
	Checked bool   `gorm:"Column:checked" sql:"DEFAULT:0" json:"checked"`
}
