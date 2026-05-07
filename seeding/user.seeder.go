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

type UserSeeder struct{}

func NewUserSeeder() *UserSeeder {
	return &UserSeeder{}
}

func (s *UserSeeder) GetName() string {
	return "UserSeeder"
}

func (s *UserSeeder) Seed(db *gorm.DB, options SeedOptions) error {
	var count int64
	db.Model(&model.User{}).Count(&count)

	switch options.Mode {
	case Skip:
		if count > 0 {
			fmt.Println("Users already exist, skipping...")
			return nil
		}

	case Replace:
		if count > 0 {
			fmt.Println("Clearing existing users...")
			if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.User{}).Error; err != nil {
				return fmt.Errorf("failed to clear users: %w", err)
			}
		}

	case Append:
		fmt.Println("Appending users...")
	}

	usersToSeed := []struct {
		Email          string
		DepartmentName string
		Password       string
	}{
		{Email: "cmi@hmtc-its.com", DepartmentName: "Creative Media Information", Password: "password_hardcoded"},
		{Email: "external@hmtc-its.com", DepartmentName: "External Affairs", Password: "password_hardcoded"},
		{Email: "internal@hmtc-its.com", DepartmentName: "Internal Affairs", Password: "password_hardcoded"},
		{Email: "sekretariat@hmtc-its.com", DepartmentName: "Board Of Directors", Password: "password_hardcoded"},
		{Email: "edd@hmtc-its.com", DepartmentName: "Entrepreneurship Development Departement", Password: "password_hardcoded"},
		{Email: "srd@hmtc-its.com", DepartmentName: "Student Resources Development", Password: "password_hardcoded"},
		{Email: "ssd@hmtc-its.com", DepartmentName: "Student Social Development", Password: "password_hardcoded"},
		{Email: "sti@hmtc-its.com", DepartmentName: "Student Talenta and Interests", Password: "password_hardcoded"},
		{Email: "sw@hmtc-its.com", DepartmentName: "Student Welfare", Password: "password_hardcoded"},
		{Email: "ristek@hmtc-its.com", DepartmentName: "Research and Technology", Password: "password_hardcoded"},
	}
	var seeded, skipped, failed int

	for _, userData := range usersToSeed {
		var existingUser model.User
		result := db.Where("email = ?", userData.Email).First(&existingUser)

		if result.Error == nil {
			fmt.Printf("User already exists with Email '%s' (skipping)\n", userData.Email)
			skipped++
			continue
		}

		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Printf("Error checking for existing user with Email '%s': %v\n", userData.Email, result.Error)
			failed++
			continue
		}

		hashedPassword, err := utils.HashPassword(userData.Password)
		if err != nil {
			fmt.Printf("Failed to hash password for user '%s': %v\n", userData.Email, err)
			failed++
			continue
		}

		user := model.User{
			Id:             uuid.New(),
			Email:          userData.Email,
			PasswordHash:   hashedPassword,
			DepartmentName: userData.DepartmentName,
			Role:           model.Admin,
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
			fmt.Printf("Created new user: %s\n", user.Email)
			seeded++
		}
	}

	fmt.Printf("Users: %d created, %d skipped, %d failed\n", seeded, skipped, failed)

	if failed > 0 && seeded == 0 && len(usersToSeed) > 0 {
		return fmt.Errorf("all user creation attempts failed")
	}

	return nil
}
