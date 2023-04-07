package postgresql

import (
	"errors" // Importing errors package for error handling
	"fmt"

	"github.com/google/uuid"
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

// FindByUUID retrieves a record from the database based on the provided UUID and assigns the result to the provided out interface.
// Returns the out interface and an error if the record is not found or if there is any issue during the query execution.
func (r *Client) FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error) {
	// Query the database for the record with the specified UUID and store the result in the out interface
	err := r.db.Where("uuid = ?", uuid).First(out).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("record not found")
		}
		return nil, err
	}
	return out, nil
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
	err := c.db.Save(value).Error
	if err != nil {
		return errors.New("failed to update record: " + err.Error())
	}
	return nil
}

// UpdateColumns updates a specific column of a record in the database.
// Returns an error if something goes wrong.
func (c *Client) UpdateColumns(value interface{}, column string, updateValue interface{}) error {
	err := c.db.Model(value).Update(column, updateValue).Error
	return err
}

// First returns the first record that matches the given conditions.
// This function returns the first record that matches the given conditions and returns an error if the operation fails.
func (c *Client) First(dest interface{}, conditions ...interface{}) error {
	if dest == nil {
		return errors.New("destination value cannot be nil")
	}
	err := c.db.First(dest, conditions...).Error
	if err != nil {
		return errors.New("failed to retrieve first record: " + err.Error())
	}
	return nil
}

// Delete deletes a record from the database based on the provided interface{}.
// This function deletes a record from the database using the given interface{} and returns an error if the operation fails.
func (c *Client) Delete(out interface{}) error {
	err := c.db.Delete(out).Error
	if err != nil {
		return errors.New("failed to delete record: " + err.Error())
	}
	return nil
}
