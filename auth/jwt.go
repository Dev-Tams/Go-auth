package auth

import (
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

//added a secret key for jwt

var jwtSecret = []byte("secret-key")

func GenerateJWT(UserID string) (string, error){
	claims := jwt.MapClaims{
		"user_id" : UserID,
		"exp" : time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
func ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, ok := claims["user_id"].(string)
		if !ok {
			return "", errors.New("invalid user_id in token")
		}
		return userID, nil
	}

	return "", errors.New("invalid token")
}
