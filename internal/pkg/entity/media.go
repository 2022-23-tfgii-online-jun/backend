// Package entity defines the domain entities (models) for the application.
package entity

import (
	"time"
)

// TableName returns the name of the table corresponding to the Mediaentity in the database.
func (*Media) TableName() string {
	return "media"
}

// Media represents a struct
type Media struct {
	ID         int       `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	RecipeID   int       `gorm:"Column:recipe_id" json:"recipe_id"`
	MediaType  string    `gorm:"Column:media_type" binding:"media_type" json:"media_type"`
	MediaURL   string    `gorm:"Column:media_url" binding:"media_url" json:"media_url"`
	MediaThumb string    `gorm:"Column:media_thumb" binding:"media_thumb" json:"media_thumb"`
	CreatedAt  time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
}
