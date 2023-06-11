package entity

import (
	"time"

	"github.com/google/uuid"
)

// TableName returns the name of the table corresponding to the SymptomUser entity in the database.
func (*SymptomUser) TableName() string {
	return "symptom_user"
}

// SymptomUser represents a struct for the symptom_user table
type SymptomUser struct {
	ID        int       `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	UserID    int       `gorm:"Column:user_id" json:"user_id"`
	SymptomID int       `gorm:"Column:symptom_id" json:"symptom"`
	CreatedAt time.Time `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
}

type RequestCreateSymptomUser struct {
	SymptomUUID uuid.UUID `json:"sypmtom"`
}

type RemoveUserFromSymptom struct {
	UserID    int   `gorm:"Column:user_id"`
	SymptomID int64 `gorm:"Column:symptom_id"`
}
