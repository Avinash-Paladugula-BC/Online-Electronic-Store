package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key")

type user struct {
	UserID uint
	Email  string
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint, email string) (string, error) {
	expiry := time.Now().Add(24 * 7 * time.Hour)

	claims := &user{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiry),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (*user, error) {
	usr := &user{}
	token, err := jwt.ParseWithClaims(tokenStr, usr, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return usr, nil
}
