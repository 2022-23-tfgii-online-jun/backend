// Package entity defines the domain entities (models) for the application.
package entity

import (
	"time"
)

const (
	// RoleUser represents the default role for regular users.
	RoleUser = "user"
	// RoleAdmin represents the admin role.
	RoleAdmin = "admin"
)

// TableName returns the name of the table corresponding to the UserRole entity in the database.
func (*UserRole) TableName() string {
	return "role_user"
}

// UserRole represents a struct for users role
type UserRole struct {
	ID        int       `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	UserID    int       `gorm:"Column:user_id" json:"user_id"`
	RoleID    int       `gorm:"Column:role_id" json:"role_id"`
	CreatedAt time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"-"`
}
