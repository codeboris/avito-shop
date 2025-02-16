package jwtutil

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey string

func InitJWT(JWTSecret string) {
	secretKey = JWTSecret
}

func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ValidateToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("неверный метод подписи")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("недействительный токен")
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return 0, errors.New("неверный формат токена")
	}

	return int(userID), nil
}
