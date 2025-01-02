package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

func GenerateToken() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", errors.New("Failed to generate token")
	}
	return hex.EncodeToString(bytes), nil
}
