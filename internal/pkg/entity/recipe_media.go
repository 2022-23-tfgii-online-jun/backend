package entity

import (
	"time"
)

// TableName returns the name of the table corresponding to the RecipeMedia entity in the database.
func (*RecipeMedia) TableName() string {
	return "recipe_media"
}

// RecipeMedia represents a struct for recipe_media
type RecipeMedia struct {
	ID        int       `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	RecipeID  int       `gorm:"Column:recipe_id" json:"-"`
	MediaID   int       `gorm:"Column:media_id" json:"-"`
	CreatedAt time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"-"`
}
