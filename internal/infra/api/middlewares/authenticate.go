package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/emur-uy/backend/internal/infra/api/middlewares/jwtutils"
	"github.com/gin-gonic/gin"
)

// Authenticate is a middleware to validate JWT tokens and extract claims for authenticated requests.
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Retrieve the JWT from the Authorization header.
		token, err := getJwt(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// 2. Parse and validate the JWT.
		jwtToken, err := jwtutils.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if !jwtutils.IsValidToken(jwtToken) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
			return
		}

		// 3. Extract claims and set them in the Gin context.
		jwtutils.SetClaims(c, jwtToken)
	}
}

// getJwt retrieves the JWT from the Authorization header.
func getJwt(c *gin.Context) (string, error) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader == "" {
		return "", errors.New("Missing token")
	}
	jwtToken := strings.ReplaceAll(authorizationHeader, "Bearer", "")
	return strings.TrimSpace(jwtToken), nil
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
