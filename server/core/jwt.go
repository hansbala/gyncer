package core

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/hansbala/gyncer/config"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// generates a new JWT auth token
func GenerateJWT(email string) (string, error) {
	config := config.GetConfig()
	secretToken := config.Server.JwtSecret
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSting, err := token.SignedString([]byte(secretToken))
	if err != nil {
		return "", err
	}

	return tokenSting, nil
}

// validates a signed JWT auth token
func ValidateJWT(signedToken string) (*Claims, error) {
	config := config.GetConfig()
	secretToken := config.Server.JwtSecret
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretToken), nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		fmt.Println("could not parse claims")
		return nil, errors.New("could not parse claims")
	}
	return claims, nil
}
