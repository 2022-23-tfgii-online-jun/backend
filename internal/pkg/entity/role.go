// Package entity defines the domain entities (models) for the application.
package entity

import (
	"time"
)

// TableName returns the name of the table corresponding to the Role entity in the database.
func (*Role) TableName() string {
	return "roles"
}

// Role represents a struct for roles
type Role struct {
	ID        int       `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	Role      string    `gorm:"Column:role" json:"role"`
	CreatedAt time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"-"`
}
