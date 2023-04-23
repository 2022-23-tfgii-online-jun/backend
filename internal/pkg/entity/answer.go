// Package entity defines the domain entities (models) for the application.
package entity

import (
	"time"

	"github.com/google/uuid"
)

// TableName returns the name of the table corresponding to the Answer entity in the database.
func (*Answer) TableName() string {
	return "answers"
}

// Answer represents a struct for answers
type Answer struct {
	ID         int       `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	UUID       uuid.UUID `gorm:"Column:uuid" json:"uuid"`
	UserID     int       `gorm:"Column:user_id" json:"-"`
	QuestionID int       `gorm:"Column:question_id" json:"-"`
	Text       string    `gorm:"Column:text" binding:"required" json:"text"`
	CreatedAt  time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
}

// RequestCreateAnswer represents a struct for creating answers
type RequestCreateAnswer struct {
	QuestionUUID uuid.UUID `json:"question_uuid"`
	Text         string    `binding:"required" json:"text"`
}
