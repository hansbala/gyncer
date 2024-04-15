package core

import (
	"crypto/sha256"
	"fmt"
)

func HashString(data string) (string, error) {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(data))
	if err != nil {
		return "", err
	}
	// return human readable string
	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
