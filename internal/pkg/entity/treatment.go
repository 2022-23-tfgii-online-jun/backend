package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// TableName returns the name of the table corresponding to the Treatment entity in the database.
func (*Treatment) TableName() string {
	return "treatments"
}

// FrequencySlice represents a wrapper struct for the Frequency slice type.
type FrequencySlice []Frequency

// Value returns the database value for the FrequencySlice type.
func (f FrequencySlice) Value() (driver.Value, error) {
	jsonBytes, err := json.Marshal(f)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

// Scan scans the database value and assigns it to the FrequencySlice type.
func (f *FrequencySlice) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	jsonBytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan FrequencySlice: unexpected value type")
	}

	if err := json.Unmarshal(jsonBytes, &f); err != nil {
		return err
	}

	return nil
}

// ShotsSlice represents a wrapper struct for the Shots slice type.
type ShotsSlice []Shot

// Value returns the database value for the ShotsSlice type.
func (s ShotsSlice) Value() (driver.Value, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

// Scan scans the database value and assigns it to the ShotsSlice type.
func (s *ShotsSlice) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	jsonBytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan ShotsSlice: unexpected value type")
	}

	if err := json.Unmarshal(jsonBytes, s); err != nil {
		return err
	}

	return nil
}

// Treatment represents a struct for treatments
type Treatment struct {
	ID        int            `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	UserID    int            `gorm:"Column:user_id" json:"-"`
	UUID      uuid.UUID      `gorm:"Column:uuid" json:"-"`
	Name      string         `gorm:"Column:name" binding:"required" json:"name"`
	Type      string         `gorm:"Column:type" binding:"required" json:"type"`
	Frequency FrequencySlice `gorm:"Column:frequency;type:json" json:"frequency"`
	Shots     ShotsSlice     `gorm:"Column:shots;type:json" json:"shots"`
	DateStart time.Time      `gorm:"Column:date_start" json:"date_start"`
	CreatedAt time.Time      `gorm:"Column:created_at" sql:"DEFAULT:current_timestamp" json:"created_at"`
	Notes     string         `gorm:"Column:notes" json:"notes"`
}

// Frequency represents the struct for frequencies
type Frequency struct {
	Day  string   `json:"day"`
	Time []string `json:"time"`
}

// Shot represents the struct for shots
type Shot struct {
	Name string `json:"name"`
	Dose int    `json:"dose"`
}

// RequestCreateTreatment struct
type RequestCreateTreatment struct {
	Name      string         `json:"name"`
	Frequency FrequencySlice `json:"frequency"`
	Shots     ShotsSlice     `json:"shots"`
	DateStart time.Time      `json:"date_start"`
	Type      string         `json:"type"`
	Notes     string         `json:"notes"`
}

type RequestUpdateTreatment struct {
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	Frequency FrequencySlice `json:"frequency"`
	Shots     ShotsSlice     `json:"shots"`
	Notes     string         `json:"notes"`
	DateStart time.Time      `json:"date_start"`
}
