package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const SALT = 14

func GeneratePasswordHash(password string) (*string, error) {
	password_hash_bytes, err := bcrypt.GenerateFromPassword([]byte(password), SALT)
	if err != nil {
		return nil, errors.New("Fail hashing the user password")
	}
	hash := string(password_hash_bytes)
	return &hash, nil
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

