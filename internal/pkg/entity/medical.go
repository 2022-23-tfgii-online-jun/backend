// Package entity defines the domain entities (models) for the application.
package entity

import (
	"time"
)

// TableName returns the name of the table corresponding to the Medical entity in the database.
func (*Medical) TableName() string {
	return "medicals"
}

// Medical represents a struct for medical records
type Medical struct {
	ID               int64     `gorm:"Column:id;PRIMARY_KEY" json:"id"`
	FirstName        string    `gorm:"Column:first_name" json:"first_name"`
	LastName         string    `gorm:"Column:last_name" json:"last_name"`
	CjppuNumber      string    `gorm:"Column:cjppu_number" json:"cjppu_number"`
	ProfessionNumber string    `gorm:"Column:profession_number" json:"profession_number"`
	CreatedAt        time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"-"`
}
