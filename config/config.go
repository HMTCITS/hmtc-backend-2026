package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost            string
	DBUser            string
	DBPassword        string
	DBName            string
	DBPort            string
	JWTSecret         string
	JWTRefreshSecret  string
	NodeEnv           string
	ServerPort        string
	OauthClientID     string
	OauthClientSecret string
	RedirectURL       string
	GDriveFolderID    string
	SheetsID          string
	SheetsName        string
}

var AppConfig *Config

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system env")
	}

	AppConfig = &Config{
		DBHost:            getEnv("DB_HOST", ""),
		DBUser:            getEnv("DB_USER", ""),
		DBPassword:        getEnv("DB_PASS", ""),
		DBName:            getEnv("DB_NAME", ""),
		DBPort:            getEnv("DB_PORT", ""),
		JWTSecret:         getEnv("JWT_SECRET", ""),
		JWTRefreshSecret:  getEnv("JWT_REFRESH_SECRET", ""),
		NodeEnv:           getEnv("NODE_ENV", ""),
		ServerPort:        getEnv("SERVER_PORT", ""),
		OauthClientID:     getEnv("OAUTH_CLIENT_ID", ""),
		OauthClientSecret: getEnv("OAUTH_CLIENT_SECRET", ""),
		RedirectURL:       getEnv("REDIRECT_URL", ""),
		GDriveFolderID:    getEnv("GDRIVE_FOLDER_ID", ""),
		SheetsID:          getEnv("SHEETS_ID", ""),
		SheetsName:        getEnv("SHEETS_NAME", ""),
	}

	log.Println("Configuration loaded successfully")
	return AppConfig
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
