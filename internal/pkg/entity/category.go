// Package entity defines the domain entities (models) for the application.
package entity

import (
	"time"

	"github.com/google/uuid"
)

// TableName returns the name of the table corresponding to the Role entity in the database.
func (*Category) TableName() string {
	return "categories"
}

// Category represents a struct for categories
type Category struct {
	ID        int       `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	UUID      uuid.UUID `gorm:"Column:uuid" json:"uuid"`
	Name      string    `gorm:"Column:name" binding:"required" json:"name"`
	CreatedAt time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
}

type AddArticleToCategoryRequest struct {
	Category uuid.UUID `json:"category"`
	Article  uuid.UUID `json:"article"`
}
