package bcrypt_utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


func HashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password))
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword checks if provided password is correct when
// compared to the hashed one.
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}