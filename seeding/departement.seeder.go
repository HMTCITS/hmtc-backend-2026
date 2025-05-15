package seeding

import (
	"errors"
	"fmt"
	"time"

	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DepartementSeeder struct {
	departements []model.Departement
}

func NewDepartementSeeder() *DepartementSeeder {
	return &DepartementSeeder{
		departements: []model.Departement{
			{
				Id:   uuid.New(),
				Name: "Creative Media Information",
				Timestamp: model.Timestamp{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			{
				Id:   uuid.New(),
				Name: "External Affairs",
				Timestamp: model.Timestamp{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			{
				Id:   uuid.New(),
				Name: "Internal Affairs",
				Timestamp: model.Timestamp{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			{
				Id:   uuid.New(),
				Name: "Board Of Directors",
				Timestamp: model.Timestamp{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			{
				Id:   uuid.New(),
				Name: "Entrepreneurship Development Departement",
				Timestamp: model.Timestamp{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			{
				Id:   uuid.New(),
				Name: "Student Resources Development",
				Timestamp: model.Timestamp{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			{
				Id:   uuid.New(),
				Name: "Student Social Development",
				Timestamp: model.Timestamp{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			{
				Id:   uuid.New(),
				Name: "Student Talenta and Interests",
				Timestamp: model.Timestamp{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			{
				Id:   uuid.New(),
				Name: "Student Welfare, Research, and Technology",
				Timestamp: model.Timestamp{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
		},
	}
}

func (s *DepartementSeeder) GetName() string {
	return "DepartementSeeder"
}

func (s *DepartementSeeder) Seed(db *gorm.DB, options SeedOptions) error {
	var count int64
	db.Model(&model.Departement{}).Count(&count)

	switch options.Mode {
	case Skip:
		if count > 0 {
			fmt.Println("Departments already exist, skipping...")
			return nil
		}

	case Replace:
		if count > 0 {
			fmt.Println("Clearing existing departments...")
			if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Departement{}).Error; err != nil {
				return fmt.Errorf("failed to clear departments: %w", err)
			}
		}

	case Append:
		fmt.Println("Appending departments...")
	}

	var seeded, skipped, failed int

	for _, department := range s.departements {
		var existingDept model.Departement
		result := db.Where("name = ?", department.Name).First(&existingDept)

		if result.Error == nil {
			fmt.Printf("Department already exists: %s (skipping)\n", department.Name)
			skipped++
			continue
		}

		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Printf("Error checking for existing department '%s': %v\n", department.Name, result.Error)
			failed++
			continue
		}

		createResult := db.Create(&department)
		if createResult.Error != nil {
			fmt.Printf("Failed to create department '%s': %v\n", department.Name, createResult.Error)
			failed++
		} else {
			fmt.Printf("Created new department: %s\n", department.Name)
			seeded++
		}
	}

	fmt.Printf("Departments: %d created, %d skipped, %d failed\n", seeded, skipped, failed)

	if failed > 0 && seeded == 0 && len(s.departements) > 0 {
		return fmt.Errorf("all department creation attempts failed")
	}

	return nil
}
