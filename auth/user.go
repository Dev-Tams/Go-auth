package auth

import (
	"errors"
	"github.com/google/uuid"
)

// type User struct {
// 	ID       string
// 	Username string
// 	Email    string
// 	Password string
// }


func RegisterUser(username, email, password string) (User, error) {
	var existing User
	if err := DB.Where("email = ?", email).First(&existing).Error; err == nil {
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

	DB.Create(&user)
	return user, nil
}

func Login(email, password string) (string, error) {
	var user User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := CheckPassword(user.Password, password); err != nil {
		return "", errors.New("invalid email or password")
	}

	return GenerateJWT(user.ID)
}
