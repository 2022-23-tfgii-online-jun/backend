// Package entity defines the domain entities (models) for the application.
package entity

import (
	"time"

	"github.com/google/uuid"
)

// TableName sets the table name for the MedicalRecord model.
func (MedicalRecord) TableName() string {
	return "medical_records"
}

type MedicalRecord struct {
	ID                      int       `gorm:"column:id;primary_key" json:"-"`
	UUID                    uuid.UUID `gorm:"column:uuid" json:"uuid"`
	UserID                  int       `gorm:"column:user_id" json:"-"`
	HealthCareProvider      string    `gorm:"column:health_care_provider" json:"health_care_provider"`
	EmergencyMedicalService string    `gorm:"column:emergency_medical_service" json:"emergency_medical_service"`
	MultipleSclerosisType   string    `gorm:"column:multiple_sclerosis_type" json:"multiple_sclerosis_type"`
	LaboralCondition        string    `gorm:"column:laboral_condition" json:"laboral_condition"`
	Conmorbidity            bool      `gorm:"column:conmorbidity" json:"conmorbidity"`
	TreatingNeurologist     string    `gorm:"column:treating_neurologist" json:"treating_neurologist"`
	SupportNetwork          bool      `gorm:"column:support_network" json:"support_network"`
	IsDisabled              bool      `gorm:"column:is_disabled" json:"is_disabled"`
	EducationalLevel        string    `gorm:"column:educational_level" json:"educational_level"`
	CreatedAt               time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt               time.Time `gorm:"column:updated_at" json:"updated_at"`
}
