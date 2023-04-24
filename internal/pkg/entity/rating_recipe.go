// Package entity defines the domain entities (models) for the application.
package entity

// TableName returns the name of the table corresponding to the Role entity in the database.
func (*Vote) TableName() string {
	return "rating_recipes"
}

// Recipe represents a struct for articles
type Vote struct {
	ID       int `gorm:"Column:id;PRIMARY_KEY" json:"-"`
	UserID   int `gorm:"Column:user_id" json:"-"`
	RecipeID int `gorm:"Column:recipe_id" json:"-"`
	Level    int `gorm:"Column:level" binding:"required" json:"level"`
}

// RequestVoteRecipe is a struct for handling a voting request for a recipe.
type RequestVoteRecipe struct {
	Vote int `json:"vote" binding:"required,min=1,max=5"`
}
