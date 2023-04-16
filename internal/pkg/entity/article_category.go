package entity

import (
	"time"

	"github.com/google/uuid"
)

// TableName - returns name of the table
func (*ArticleCategory) TableName() string {
	return "article_category"
}

type ArticleCategory struct {
	ArticleID  int       `gorm:"Column:article_id"`
	CategoryID int       `gorm:"Column:category_id"`
	CreatedAt  time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp"`
}

// CreateArticleCategoryRequest defines the request structure for creating a relation between article and category.
type CreateArticleCategoryRequest struct {
	ArticleUUID  uuid.UUID `json:"article_uuid" binding:"required"`
	CategoryUUID uuid.UUID `json:"category_uuid" binding:"required"`
}
