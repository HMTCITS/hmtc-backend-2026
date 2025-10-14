package repository

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/HMTCITS/hmtc-backend-2025/model"
)

const tokenFile = "token.json"

func SaveRefreshToken(refreshToken string) error {
	t := model.TokenFile{RefreshToken: strings.TrimSpace(refreshToken)}
	b, _ := json.MarshalIndent(t, "", "  ")
	return os.WriteFile(tokenFile, b, 0644)
}

func LoadRefreshToken() (string, error) {
	b, err := os.ReadFile(tokenFile)
	if err != nil {
		return "", err
	}
	var t model.TokenFile
	if err := json.Unmarshal(b, &t); err != nil {
		return "", err
	}
	return strings.TrimSpace(t.RefreshToken), nil
}
