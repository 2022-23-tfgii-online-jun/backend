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
		jwtToken, err := parseToken(token)
		if err != nil || !isValidToken(jwtToken) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 3. Extract claims and set them in the Gin context.
		setClaims(c, jwtToken)
	}
}

// getJwt retrieves the JWT from the Authorization header.
func getJwt(c *gin.Context) string {
	authorizationHeader := c.Request.Header.Get("Authorization")
	jwtToken := strings.ReplaceAll(authorizationHeader, "Bearer", "")
	return strings.TrimSpace(jwtToken)
}

// parseToken parses and returns a JWT token, or an error if the token is invalid.
func parseToken(token string) (*jwt.Token, error) {
	jwtKey := []byte(config.Get().JWTTokenKey)
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
}

// isValidToken checks if the provided JWT token is valid.
func isValidToken(jwtToken *jwt.Token) bool {
	return jwtToken.Valid
}

// setClaims extracts claims from the JWT token and sets them in the Gin context.
func setClaims(c *gin.Context, jwtToken *jwt.Token) {
	claims := jwtToken.Claims.(jwt.MapClaims)

	c.Set("email", claims["email"])
	c.Set("userUUID", claims["user_uuid"])
}
