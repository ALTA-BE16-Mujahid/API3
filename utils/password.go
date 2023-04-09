package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		// return user, err
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}
