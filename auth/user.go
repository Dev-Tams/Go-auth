package auth

import (
	"errors"
	"github.com/google/uuid"
)

type User struct {
	ID       string
	Username string
	Email    string
	Password string
}

var users = map[string]User{}

func RegisterUser(username, email, password string) (User, error) {
	if _, exists := users[email]; exists {
		return User{}, errors.New("user already exists")
	}

	hashed, err := HashPassword(password)
	if err != nil {
		return User{}, err
	}

	user := User{
		ID:       uuid.New().String(),
		Username: username,
		Email:    email,
		Password: hashed,
	}

	users[email] = user
	return user, nil
}

func Login(email, password string) (string, error) {
	user, exists := users[email]
	if !exists {
		return "", errors.New("invalid email or password")
	}

	if err := CheckPassword(user.Password, password); err != nil {
		return "", errors.New("invalid email or password")
	}

	return GenerateJWT(user.ID)
}
