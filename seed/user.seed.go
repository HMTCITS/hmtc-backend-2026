package main

import (
	"fmt"
	"log"
	"time"

	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Database configuration
	dsn := "host=localhost user=postgres password=root dbname=backend_hmtc port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// AutoMigrate models
	err = db.AutoMigrate(&model.Departement{}, &model.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database models: %v", err)
	}

	// Seed Departements
	departements := []model.Departement{
		{
			Id:   uuid.New(),
			Name: "Creative Media Information",
			Timestamp: model.Timestamp{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	for _, departement := range departements {
		if err := db.FirstOrCreate(&departement, model.Departement{Name: departement.Name}).Error; err != nil {
			log.Printf("Failed to seed departement '%s': %v", departement.Name, err)
		} else {
			fmt.Printf("Seeded departement: %s\n", departement.Name)
		}
	}

	// Seed Users
	users := []model.User{
		{
			Id:            uuid.New(),
			NRP:           "5025221052",
			DepartementId: &departements[0].Id,
			Timestamp: model.Timestamp{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	for _, user := range users {
		if err := db.FirstOrCreate(&user, model.User{NRP: user.NRP}).Error; err != nil {
			log.Printf("Failed to seed user '%s': %v", user.NRP, err)
		} else {
			fmt.Printf("Seeded user: %s\n", user.NRP)
		}
	}
}
