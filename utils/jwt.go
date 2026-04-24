package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(userId uuid.UUID, role string, departement string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":         userId,
		"role":        role,
		"departement": departement,
		"exp":         jwt.NewNumericDate(time.Now().Add(time.Minute * 120)),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func GenerateRefreshToken(userId uuid.UUID, role string, departement string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":         userId,
		"role":        role,
		"departement": departement,
		"exp":         jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenStr string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
