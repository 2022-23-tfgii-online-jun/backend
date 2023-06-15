// Package entity defines the domain entities (models) for the application.
package entity

import (
	"time"

	"github.com/google/uuid"
)

// TableName returns the name of the table corresponding to the Role entity in the database.
func (*Recipe) TableName() string {
	return "recipes"
}

// Recipe represents a struct for articles
type Recipe struct {
	ID          int       `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	UUID        uuid.UUID `gorm:"Column:uuid" json:"uuid"`
	Name        string    `gorm:"Column:name" binding:"required" json:"name"`
	Ingredients string    `gorm:"Column:ingredients" binding:"required" json:"ingredients"`
	Elaboration string    `gorm:"Column:elaboration" binding:"required" json:"elaboration"`
	Category    int       `gorm:"Column:category_id" binding:"required" json:"category"`
	Time        int       `gorm:"Column:time" binding:"required" json:"time"`
	IsPublished bool      `gorm:"Column:is_published" sql:"DEFAULT:0" json:"is_published"`
	CreatedAt   time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
}

// RequestCreateRecipe represents a struct for creating articles
type RequestCreateRecipe struct {
	Name        string `form:"name" binding:"required"`
	Ingredients string `form:"ingredients" binding:"required"`
	Elaboration string `form:"elaboration" binding:"required"`
	Time        int    `form:"time" binding:"required"`
	Category    int    `form:"category" binding:"required"`
}

// RequestUpdateRecipe represents a struct for creating articles
type RequestUpdateRecipe struct {
	Name        string `form:"name" binding:"required"`
	Ingredients string `form:"ingredients" binding:"required"`
	Elaboration string `form:"elaboration" binding:"required"`
	Time        int    `form:"time" binding:"required"`
	Category    int    `form:"category" binding:"required"`
}

// RecipeWithMediaURLs represents a recipe with associated media URLs.
type RecipeWithMediaURLs struct {
	Recipe    *Recipe  `json:"recipe"`
	MediaURLs []string `json:"media"`
}
