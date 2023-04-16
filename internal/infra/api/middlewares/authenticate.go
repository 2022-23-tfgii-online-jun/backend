package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/emur-uy/backend/internal/infra/api/middlewares/jwtutils"
	"github.com/gin-gonic/gin"
)

// RegisterAuthMiddlewares is a function that sets up the authentication and authorization middlewares
// on the given gin.RouterGroup instance for the specified role.
func RegisterAuthMiddlewares(role string, r *gin.RouterGroup) {
	r.Use(Authenticate())
	r.Use(Authorize(role))
}

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
		return "", errors.New("missing token")
	}
	jwtToken := strings.ReplaceAll(authorizationHeader, "Bearer", "")
	return strings.TrimSpace(jwtToken), nil
}

// Authorize is Authorization middleware to validate roles for API calls
func Authorize(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user making the request has any of the specified roles
		// If the user has any of the roles, call the next middleware/handler
		// If the user does not have any of the roles, return an error response
		userRole := c.GetString("role")

		authorized := false
		for _, role := range allowedRoles {
			if userRole == role {
				authorized = true
				break
			}
		}

		if !authorized {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to access this resource"})
			return
		}
		c.Next()
	}
}
