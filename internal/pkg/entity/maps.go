package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// PhoneSlice represents a slice of phone details.
type PhoneSlice []Phone

// Phone represents the details of a phone.
type Phone struct {
	Number string `gorm:"Column:number" json:"number"`
}

// HoursAvailabilitySlice represents a slice of hours availability.
type HoursAvailabilitySlice []HoursAvailability

// HoursAvailability represents the hours availability.
type HoursAvailability struct {
	Day       string `gorm:"Column:day" json:"day"`
	OpenTime  string `gorm:"Column:open_time" json:"open_time"`
	CloseTime string `gorm:"Column:close_time" json:"close_time"`
}

type Map struct {
	ID                int64                  `gorm:"Column:id" json:"-"`
	UUID              uuid.UUID              `gorm:"Column:uuid" json:"uuid"`
	Name              string                 `gorm:"Column:name" json:"name"`
	Latitude          string                 `gorm:"Column:latitude" json:"latitude"`
	Longitude         string                 `gorm:"Column:longitude" json:"longitude"`
	Type              int                    `gorm:"Column:type" json:"type"`
	HoursAvailability HoursAvailabilitySlice `gorm:"Column:hours_availability" json:"hours_availability"`
	Phone             PhoneSlice             `gorm:"Column:phone" json:"phone"`
	IsPublished       bool                   `gorm:"Column:is_published" json:"is_published"`
	CreatedAt         time.Time              `gorm:"Column:created_at" json:"created_at"`
	UpdatedAt         time.Time              `gorm:"Column:updated_at" json:"updated_at"`
}

// HoursAvailabilitySlice implements the sql.Scanner and driver.Valuer interfaces.
func (h *HoursAvailabilitySlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan HoursAvailabilitySlice: unexpected value type")
	}
	return json.Unmarshal(bytes, &h)
}

// Value implements the driver.Valuer interface.
func (h HoursAvailabilitySlice) Value() (driver.Value, error) {
	bytes, err := json.Marshal(h)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// PhoneSlice implements the sql.Scanner and driver.Valuer interfaces.
func (p *PhoneSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan PhoneSlice: unexpected value type")
	}
	return json.Unmarshal(bytes, &p)
}

// Value implements the driver.Valuer interface.
func (p PhoneSlice) Value() (driver.Value, error) {
	bytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

type RequestCreateUpdateMap struct {
	Name              string                 `json:"name"`
	Latitude          string                 `json:"latitude"`
	Longitude         string                 `json:"longitude"`
	Type              int                    `json:"type"`
	HoursAvailability HoursAvailabilitySlice `json:"hours_availability"`
	Phone             PhoneSlice             `json:"phone"`
	IsPublished       bool                   `json:"is_published"`
}
