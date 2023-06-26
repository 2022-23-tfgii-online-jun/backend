package entity

import (
	"time"
)

// TableName returns the name of the table corresponding to the ArticleMedia entity in the database.
func (*ArticleMedia) TableName() string {
	return "article_media"
}

// ArticleMedia represents a struct for article_media
type ArticleMedia struct {
	ID        int       `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	ArticleID int       `gorm:"Column:article_id" json:"-"`
	MediaID   int       `gorm:"Column:media_id" json:"-"`
	CreatedAt time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"-"`
}
