package entity

// DefaultCredentials represents email/password combination.
type DefaultCredentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
