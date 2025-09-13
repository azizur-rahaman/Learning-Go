package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRETKEY = "your_secret_key"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (int64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(SECRETKEY), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// Check token expiration
	exp, ok := claims["exp"].(float64)
	if !ok {
		return 0, errors.New("invalid token expiration")
	}

	if time.Now().Unix() > int64(exp) {
		return 0, errors.New("token expired")
	}

	userId, ok := claims["userId"].(float64)

	if !ok {
		return 0, errors.New("invalid token user id")
	}

	return int64(userId), nil
}
