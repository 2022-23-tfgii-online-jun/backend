package postgresql

import (
	"errors" // Importing errors package for error handling
	"fmt"

	"github.com/emur-uy/backend/internal/pkg/entity"
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

func (r *Client) FindByUUID(uuid string) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.Where("uuid = ?", uuid).First(user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return user, nil
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
