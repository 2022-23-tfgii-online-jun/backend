package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// TableName returns the name of the table corresponding to the Reminder entity in the database.
func (*Reminder) TableName() string {
	return "reminders"
}

// NotificationSlice represents a wrapper struct for the Notification slice type.
type NotificationSlice []Notification

// Value returns the database value for the NotificationSlice type.
func (n NotificationSlice) Value() (driver.Value, error) {
	// Convert the NotificationSlice to JSON bytes.
	jsonBytes, err := json.Marshal(n)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

// Scan scans the database value and assigns it to the NotificationSlice type.
func (n *NotificationSlice) Scan(value interface{}) error {
	// Check if the value is nil.
	if value == nil {
		return nil
	}

	// Convert the value to []byte.
	jsonBytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan NotificationSlice: unexpected value type")
	}

	fmt.Println(string(jsonBytes))

	// Unmarshal the JSON bytes to a NotificationSlice.
	if err := json.Unmarshal(jsonBytes, &n); err != nil {
		return err
	}

	return nil
}

// TaskSlice represents a wrapper struct for the Task slice type.
type TaskSlice []Task

// Value returns the database value for the TaskSlice type.
func (t TaskSlice) Value() (driver.Value, error) {
	// Convert the TaskSlice to JSON bytes.
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

// Scan scans the database value and assigns it to the TaskSlice type.
func (t *TaskSlice) Scan(value interface{}) error {
	// Check if the value is nil.
	if value == nil {
		return nil
	}

	// Convert the value to []byte.
	jsonBytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan TaskSlice: unexpected value type")
	}

	// Unmarshal the JSON bytes to a TaskSlice.
	if err := json.Unmarshal(jsonBytes, &t); err != nil {
		return err
	}

	return nil
}

// Reminder represents a struct for reminders
type Reminder struct {
	ID           int               `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	UserID       int               `gorm:"Column:user_id" json:"-"`
	UUID         uuid.UUID         `gorm:"Column:uuid" json:"uuid"`
	Name         string            `gorm:"Column:name" binding:"required" json:"name"`
	Type         string            `gorm:"Column:type" binding:"required" json:"type"`
	Date         time.Time         `gorm:"Column:date" binding:"required" json:"date"`
	Notification NotificationSlice `gorm:"Column:notification;type:json" json:"notification"`
	Task         TaskSlice         `gorm:"Column:task;type:json" json:"task"`
	Note         string            `gorm:"Column:note" json:"note"`
	Medical      int               `gorm:"Column:medical_id" json:"medical_id"`
	IsActive     bool              `gorm:"Column:is_active" sql:"DEFAULT:1" json:"is_active"`
	CreatedAt    time.Time         `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
}

// Notification represents the struct for notifications
type Notification struct {
	DaysOrHours string `json:"days_or_hours"`
	HoursBefore int    `json:"hours_before"`
}

// Task represents the struct for tasks
type Task struct {
	Name    string `json:"name"`
	Checked bool   `json:"checked"`
}

// RequestCreateReminder represents a struct for RequestCreateReminder
type RequestCreateReminder struct {
	Name         string         `gorm:"Column:name" binding:"required" form:"name"`
	Type         string         `gorm:"Column:type" binding:"required" form:"type"`
	Date         time.Time      `gorm:"Column:date" binding:"required" form:"date" time_format:"02/01/2006"`
	Notification []Notification `gorm:"Column:notification" form:"notification"`
	Task         []Task         `gorm:"Column:task" form:"task"`
	Medical      int            `gorm:"Column:medical" form:"medical"`
	Note         string         `gorm:"Column:note" form:"note"`
}

// GetReminderResponse represents a struct for GetReminderResponse
type GetReminderResponse struct {
	UUID         uuid.UUID                  `json:"uuid"`
	Name         string                     `json:"name"`
	Type         string                     `json:"type"`
	Date         time.Time                  `json:"date"`
	Notification NotificationSlice          `json:"notification"`
	Task         TaskSlice                  `json:"task"`
	Note         string                     `json:"note"`
	Medical      int                        `json:"medical"`
	IsActive     bool                       `json:"is_active"`
	Media        []GetReminderMediaResponse `json:"media"`
}

// GetReminderMediaResponse represents a struct for GetReminderMediaResponse
type GetReminderMediaResponse struct {
	MediaURL   string `json:"media_url"`
	MediaThumb string `json:"media_thumb"`
}

// RequestUpdateReminder represents a struct for RequestUpdateReminder
type RequestUpdateReminder struct {
	Name         string         `form:"name"`
	Type         string         `form:"type"`
	Date         time.Time      `form:"date" time_format:"02/01/2006"`
	Notification []Notification `form:"notification"`
	Task         []Task         `form:"task"`
	Note         string         `form:"note"`
	Medical      string         `form:"medical"`
}
