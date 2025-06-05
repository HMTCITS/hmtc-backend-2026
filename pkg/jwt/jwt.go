package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	accessToken  TokenType = "Access"
	refreshToken TokenType = "Refresh"
)

func GenerateToken(payload map[string]string, tokenType TokenType) (string, error) {
	var expiredAt time.Time

	switch tokenType {
	case accessToken:
		expiredAt = time.Now().Add(72 * time.Hour)
	case refreshToken:
		expiredAt = time.Now().Add(168 * time.Hour)
	}

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt.Unix()
	claims["iss"] = "hmtc"

	for k, v := range payload {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(getSecretKey()))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GetPayload(tokenString string) (map[string]string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(getSecretKey()), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token expired")
		}
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	payload := make(map[string]string)
	for key, value := range claims {
		if strValue, ok := value.(string); ok {
			payload[key] = strValue
		}
	}

	return payload, nil
}

func getSecretKey() string {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "apahayo"
	}
	return secretKey
}
