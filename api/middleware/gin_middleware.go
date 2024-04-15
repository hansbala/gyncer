package middleware

import (
	"fmt"
	"gyncer/core"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTTokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is missing"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is not in Bearer token format"})
			return
		}

		// Token validation logic
		token := parts[1]
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
