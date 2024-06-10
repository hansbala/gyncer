package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hansbala/gyncer/core"
)

func GetTokenFromRequest(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("authorization header is not in Bearer token format")
	}

	return parts[1], nil
}

func JWTTokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := GetTokenFromRequest(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		// Token validation logic
		if _, err := core.ValidateJWT(token); err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "JWT token is invalid"})
			return
		}

		// set the token in the context and move on
		c.Set("JWT_TOKEN", token)
		c.Next()
	}
}
