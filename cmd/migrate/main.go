package main

import (
	"fmt"
	"os"

	"github.com/HMTCITS/hmtc-backend-2025/config"
	"github.com/HMTCITS/hmtc-backend-2025/migration"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("Running database migration...")

	appConfig := config.LoadConfig()

	db := config.ConnectDatabase(appConfig)
	defer config.CloseDatabase(db)

	if err := migration.Migrate(db); err != nil {
		fmt.Printf("Error migrating database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Database migration completed successfully")
}
