package entity

import (
	"time"

	"github.com/google/uuid"
)

// TableName returns the name of the table corresponding to the Media entity in the database.
func (*Media) TableName() string {
	return "media"
}

// Media represents a struct for media
type Media struct {
	ID         int       `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	UUID       uuid.UUID `gorm:"Column:uuid" json:"uuid"`
	MediaURL   string    `gorm:"Column:media_url" binding:"required" json:"media_url"`
	MediaThumb string    `gorm:"Column:media_thumb" binding:"required" json:"media_thumb"`
	CreatedAt  time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
}
