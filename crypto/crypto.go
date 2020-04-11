package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordEncrypt: Convert raw password to hash
func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// CompareHashAndPassword: Compare raw password to its hash
func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
