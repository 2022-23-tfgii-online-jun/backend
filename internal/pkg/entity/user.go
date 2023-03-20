// Package entity defines the domain entities (models) for the application.
package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	// RoleUser represents the default role for regular users.
	RoleUser = "user"
)

// TableName returns the name of the table corresponding to the User entity in the database.
func (*User) TableName() string {
	return "users"
}

// User represents a struct for users, defining their properties and relationships
// with other entities in the system.
type User struct {
	ID           int        `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	UUID         uuid.UUID  `gorm:"Column:uuid" json:"uuid"`
	FirstName    string     `gorm:"Column:first_name" json:"first_name" binding:"required,min=3,max=100"`
	LastName     string     `gorm:"Column:last_name" json:"last_name" binding:"required,min=3,max=100"`
	ProfileImage string     `gorm:"Column:profile_image" json:"profile_image"`
	DateOfBirth  string     `gorm:"Column:date_of_birth" json:"date_of_birth" binding:"required"`
	Sex          string     `gorm:"Column:sex" json:"sex" binding:"required"`
	Email        string     `gorm:"Column:email" binding:"required,email" json:"email"`
	Password     string     `gorm:"Column:password" sql:"DEFAULT:NULL" validate:"required" binding:"required" json:"password"`
	UserType     string     `gorm:"Column:user_type" json:"user_type" binding:"required"`
	IsActive     bool       `gorm:"Column:is_active" sql:"DEFAULT:0" json:"is_active"`
	IsBanned     bool       `gorm:"Column:is_banned" sql:"DEFAULT:false" json:"is_banned"`
	City         string     `gorm:"Column:city" json:"city" binding:"required"`
	Country      string     `gorm:"Column:country" json:"country" binding:"required"`
	CreatedAt    time.Time  `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"-"`
	UpdatedAt    time.Time  `gorm:"Column:updated_at" sql:"DEFAULT:current_timestamp" json:"-"`
	DeletedAt    *time.Time `gorm:"Column:deleted_at" sql:"DEFAULT:NULL" json:"-"`
}

// SignUp represents a struct for user registration, including only the necessary fields for registration.
type SignUp struct {
	Email    string `gorm:"Column:email" binding:"required,email" json:"email"`
	Password string `gorm:"Column:password" json:"password" validate:"required" binding:"required"`
}

// UpdateUser represents a struct for updating a user's properties.
type UpdateUser struct {
	ID          int     `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	FirstName   *string `gorm:"Column:first_name" json:"first_name"`
	LastName    *string `gorm:"Column:last_name" json:"last_name"`
	DateOfBirth *string `gorm:"Column:date_of_birth" json:"date_of_birth"`
	Sex         *string `gorm:"Column:sex" json:"sex"`
	UserType    *string `gorm:"Column:user_type" json:"user_type"`
	City        *string `gorm:"Column:city" json:"city"`
	Country     *string `gorm:"Column:country" json:"country"`
}
