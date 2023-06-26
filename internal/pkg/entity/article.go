// Package entity defines the domain entities (models) for the application.
package entity

import (
	"time"

	"github.com/google/uuid"
)

// TableName returns the name of the table corresponding to the Role entity in the database.
func (*Article) TableName() string {
	return "articles"
}

// Article represents a struct for articles
type Article struct {
	ID          int       `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	UUID        uuid.UUID `gorm:"Column:uuid" json:"uuid"`
	Title       string    `gorm:"Column:title" binding:"required" json:"title"`
	Content     string    `gorm:"Column:content" binding:"required" json:"content"`
	IsPublished bool      `gorm:"Column:is_published" sql:"DEFAULT:0" json:"is_published"`
	CreatedAt   time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
}

// RequestCreateArticle represents a struct for creating articles
type RequestCreateArticle struct {
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}

// RequestUpdateArticle represents a struct for creating articles
type RequestUpdateArticle struct {
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}

// ArticleWithMediaURLs represents a recipe with associated media URLs.
type ArticleWithMediaURLs struct {
	Article   *Article `json:"article"`
	MediaURLs []string `json:"media"`
}
