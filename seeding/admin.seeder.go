package seeding

import (
	"errors"
	"fmt"
	"time"

	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/HMTCITS/hmtc-backend-2025/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminSeeder struct{}

func NewAdminSeeder() *AdminSeeder {
	return &AdminSeeder{}
}

func (s *AdminSeeder) GetName() string {
	return "AdminSeeder"
}

func (s *AdminSeeder) Seed(db *gorm.DB, options SeedOptions) error {
	var count int64
	db.Model(&model.User{}).Where("role = ?", "admin").Count(&count)

	switch options.Mode {
	case Skip:
		if count > 0 {
			fmt.Println("admin already exist, skipping...")
			return nil
		}

	case Replace:
		if count > 0 {
			fmt.Println("Clearing existing admin...")
			if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.User{}).Error; err != nil {
				return fmt.Errorf("failed to clear users: %w", err)
			}
		}

	case Append:
		fmt.Println("Appending admin...")
	}

	var departements []model.Departement
	if err := db.Find(&departements).Error; err != nil {
		return fmt.Errorf("failed to fetch departments: %w", err)
	}

	if len(departements) == 0 {
		return fmt.Errorf("no departments found, make sure to run DepartementSeeder first")
	}

	departmentMap := make(map[string]uuid.UUID)
	for _, dept := range departements {
		departmentMap[dept.Name] = dept.Id
	}

	adminToSeeder := []struct {
		Email          string
		Password       string
		NRP            string
		DepartmentName string
		Role           string
	}{
		{
			Email:          "admin@hmtc-its.com",
			Password:       "password_admin",
			NRP:            "5053231014",
			DepartmentName: "Creative Media Information",
			Role:           "admin",
		},
	}
	var seeded, skipped, failed int

	for _, adminData := range adminToSeeder {
		deptId, ok := departmentMap[adminData.DepartmentName]
		if !ok {
			fmt.Printf("Warning: Department '%s' not found, skipping user '%s'\n",
				adminData.DepartmentName, adminData.Email)
			skipped++
			continue
		}

		var existingUser model.User
		result := db.Where("email = ?", adminData.Email).First(&existingUser)

		if result.Error == nil {
			fmt.Printf("User already exists with Email '%s' (skipping)\n", adminData.Email)
			skipped++
			continue
		}

		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Printf("Error checking for existing user with Email '%s': %v\n", adminData.Email, result.Error)
			failed++
			continue
		}

		hashedPassword, err := utils.HashPassword(adminData.Password)
		if err != nil {
			fmt.Printf("Failed to hash password for user '%s': %v\n", adminData.Email, err)
			failed++
			continue
		}

		nrp := adminData.NRP

		user := model.User{
			Id:            uuid.New(),
			NRP:           &nrp,
			Email:         adminData.Email,
			PasswordHash:  hashedPassword,
			Role:          model.UserRole(adminData.Role),
			DepartementId: &deptId,
			Timestamp: model.Timestamp{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		createResult := db.Create(&user)
		if createResult.Error != nil {
			fmt.Printf("Failed to create user '%s': %v\n", user.Email, createResult.Error)
			failed++
		} else {
			fmt.Printf("Created new admin: %s\n", user.Email)
			seeded++
		}
	}

	fmt.Printf("Users: %d created, %d skipped, %d failed\n", seeded, skipped, failed)

	if failed > 0 && seeded == 0 && len(adminToSeeder) > 0 {
		return fmt.Errorf("all user creation attempts failed")
	}

	return nil
}
