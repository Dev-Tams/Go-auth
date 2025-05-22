package auth

import (
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
	//hashes the password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}