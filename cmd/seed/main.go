package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/HMTCITS/hmtc-backend-2025/config"
	"github.com/HMTCITS/hmtc-backend-2025/seeding"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("Running database seeding...")

	mode := flag.String("mode", "", "Seeding mode: skip, append, or replace")
	continueOnError := flag.Bool("continue", true, "Continue seeding if individual records fail")
	flag.Parse()

	options := seeding.DefaultSeedOptions()

	if *mode != "" {
		switch strings.ToLower(*mode) {
		case "skip":
			options.Mode = seeding.Skip
		case "append":
			options.Mode = seeding.Append
		case "replace":
			options.Mode = seeding.Replace
		default:
			fmt.Printf("Invalid mode: %s. Using default mode: %s\n", *mode, options.Mode)
		}
	}

	fmt.Printf("Seeding with mode: %s\n", options.Mode)
	fmt.Printf("Continue on error: %t\n", *continueOnError)

	db := config.ConnectDatabase()
	defer config.CloseDatabase(db)

	if err := seeding.RunAllSeeders(db, options, seeding.GetSeeders()...); err != nil {
		fmt.Printf("Error seeding database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Database seeding completed successfully")
}
