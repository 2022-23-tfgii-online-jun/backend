package middlewares

import (
	"net/http"
	"strings"

	"github.com/emur-uy/backend/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Authenticate is a middleware to validate JWT tokens and extract claims for authenticated requests.
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Retrieve the JWT from the Authorization header.
		token := getJwt(c)

		// 2. Parse and validate the JWT.
		jwtToken, err := ParseToken(token)
		if err != nil || !IsValidToken(jwtToken) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 3. Extract claims and set them in the Gin context.
		SetClaims(c, jwtToken)
	}
}

// getJwt retrieves the JWT from the Authorization header.
func getJwt(c *gin.Context) string {
	authorizationHeader := c.Request.Header.Get("Authorization")
	jwtToken := strings.ReplaceAll(authorizationHeader, "Bearer", "")
	return strings.TrimSpace(jwtToken)
}

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

// Authorize is Authorization middleware to validate roles for API calls
func Authorize(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user making the request has the specified role
		// If the user has the role, call the next middleware/handler
		// If the user does not have the role, return an error response
		if c.GetString("role") != role {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to access this resource"})
			return
		}
		c.Next()
	}
}
