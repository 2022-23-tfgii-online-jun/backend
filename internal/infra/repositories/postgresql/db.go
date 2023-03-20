package postgresql

import (
	"errors" // Importing errors package for error handling

	"gorm.io/gorm" // Importing gorm package for database ORM
)

type Client struct {
	db *gorm.DB
}

// NewClient returns a new client object for performing database operations.
// This function initializes a new client object with the gorm DB object.
func NewClient() *Client {
	return &Client{
		db: Db,
	}
}

// Create stores a new record in the database.
// This function creates a new record in the database using the given value and returns an error if the operation fails.
func (c *Client) Create(value interface{}) error {
	if value == nil {
		return errors.New("input value cannot be nil")
	}
	err := c.db.Create(value).Error
	if err != nil {
		return errors.New("failed to create record: " + err.Error())
	}
	return nil
}

// CreateWithOmit stores a new record in the database and omits the specified columns.
// This function creates a new record in the database using the given value and omits the specified columns. It returns an error if the operation fails.
func (c *Client) CreateWithOmit(omitColumns string, value interface{}) error {
	if value == nil {
		return errors.New("input value cannot be nil")
	}
	err := c.db.Omit(omitColumns).Create(value).Error
	if err != nil {
		return errors.New("failed to create record with omitted columns: " + err.Error())
	}
	return nil
}

// Update updates an existing record in the database using the given value.
// This function updates an existing record in the database using the given value and returns an error if the operation fails.
func (c *Client) Update(value interface{}) error {
	if value == nil {
		return errors.New("input value cannot be nil")
	}
	err := c.db.Updates(value).Error
	if err != nil {
		return errors.New("failed to update record: " + err.Error())
	}
	return nil
}
