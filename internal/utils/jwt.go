package utils

import (
	"time"
	"user-service/pkg/globals"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
			"exp":    time.Now().Add(time.Hour * time.Duration(globals.Config.SecretKey.ExpireHours)).Unix(),
		})

	tokenString, err := token.SignedString([]byte(globals.Config.SecretKey.AccessToken))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWTToken(tokenString string) (bool, error) {
	secretKey := []byte(globals.Config.SecretKey.AccessToken)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, nil
	}

	return true, nil
}
