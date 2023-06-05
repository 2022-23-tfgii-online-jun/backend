package entity

import (
	"time"
)

// TableName returns the name of the table corresponding to the ReminderMedia entity in the database.
func (*ReminderMedia) TableName() string {
	return "reminder_media"
}

// ReminderMedia represents a struct for reminder_media
type ReminderMedia struct {
	ID         int       `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	ReminderID int       `gorm:"Column:reminder_id" json:"-"`
	MediaID    int       `gorm:"Column:media_id" json:"-"`
	CreatedAt  time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"-"`
}
