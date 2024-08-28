package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(bytes), err
}

func ValidatePassword(pwd string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil, err
}
