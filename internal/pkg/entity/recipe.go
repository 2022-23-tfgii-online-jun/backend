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
	UserID      int       `gorm:"Column:user_id" json:"-"`
	Name        string    `gorm:"Column:name" binding:"required" json:"name"`
	Description string    `gorm:"Column:description" binding:"required" json:"description"`
	Ingredients string    `gorm:"Column:ingredients" binding:"required" json:"ingredients"`
	Elaboration string    `gorm:"Column:elaboration" binding:"required" json:"elaboration"`
	PrepTime    int       `gorm:"Column:prep_time" binding:"required" json:"prep_time"`
	CookTime    int       `gorm:"Column:cook_time" binding:"required" json:"cook_time"`
	Serving     int       `gorm:"Column:serving" binding:"required" json:"servings"`
	Difficulty  int       `gorm:"Column:difficulty" binding:"required" json:"difficulty"`
	Nutrition   string    `gorm:"Column:nutrition" binding:"required" json:"nutrition"`
	IsPublished bool      `gorm:"Column:is_published" sql:"DEFAULT:0" json:"is_published"`
	CreatedAt   time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
}

// RequestCreateRecipe represents a struct for creating articles
type RequestCreateRecipe struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" binding:"required"`
	Ingredients string `form:"ingredients" binding:"required"`
	Elaboration string `form:"elaboration" binding:"required"`
	PrepTime    int    `form:"prep_time" binding:"required"`
	CookTime    int    `form:"cook_time" binding:"required"`
	Serving     int    `form:"serving" binding:"required"`
	Difficulty  int    `form:"difficulty" binding:"required"`
	Nutrition   string `form:"nutrition" binding:"required"`
}

// RequestUpdateRecipe represents a struct for creating articles
type RequestUpdateRecipe struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" binding:"required"`
	Ingredients string `form:"ingredients" binding:"required"`
	Elaboration string `form:"elaboration" binding:"required"`
	PrepTime    int    `form:"prep_time" binding:"required"`
	CookTime    int    `form:"cook_time" binding:"required"`
	Serving     int    `form:"serving" binding:"required"`
	Difficulty  int    `form:"difficulty" binding:"required"`
	Nutrition   string `form:"nutrition" binding:"required"`
}
