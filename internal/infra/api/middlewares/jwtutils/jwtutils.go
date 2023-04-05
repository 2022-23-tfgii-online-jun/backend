package jwtutils

import (
	"github.com/emur-uy/backend/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// ParseToken parses and returns a JWT token, or an error if the token is invalid.
func ParseToken(token string) (*jwt.Token, error) {
	jwtKey := []byte(config.Get().JWTTokenKey)
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
}

// IsValidToken checks if the provided JWT token is valid.
func IsValidToken(jwtToken *jwt.Token) bool {
	return jwtToken.Valid
}

// SetClaims extracts claims from the JWT token and sets them in the Gin context.
func SetClaims(c *gin.Context, jwtToken *jwt.Token) {
	claims := jwtToken.Claims.(jwt.MapClaims)

	c.Set("email", claims["email"])
	c.Set("userUUID", claims["user_uuid"])
	c.Set("role", claims["role"])
}
