// Package entity defines the domain entities (models) for the application.
package entity

import (
	"time"

	"github.com/google/uuid"
)

// TableName returns the name of the table corresponding to the Question entity in the database.
func (*Question) TableName() string {
	return "questions"
}

// Question represents a struct for questions
type Question struct {
	ID        int       `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	UUID      uuid.UUID `gorm:"Column:uuid" json:"uuid"`
	UserID    int       `gorm:"Column:user_id" json:"-"`
	Text      string    `gorm:"Column:text" binding:"required" json:"text"`
	CreatedAt time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
}

// RequestCreateQuestion represents a struct for creating questions
type RequestCreateQuestion struct {
	Text string `binding:"required" json:"text"`
}
