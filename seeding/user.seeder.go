package seeding

import (
	"errors"
	"fmt"
	"time"

	"github.com/HMTCITS/hmtc-backend-2025/model"
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

	usersToSeed := []struct {
		NRP            string
		DepartmentName string
	}{
		{
			NRP:            "5025221113",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5053231015",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025221010",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025221261",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025221285",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025231310",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025231254",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025231037",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5053231010",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025221302",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025221052",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025221115",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025221155",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025231022",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025231119",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025231240",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025231073",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025231098",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5053231002",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5054241028",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025231010",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025231067",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025231004",
			DepartmentName: "Creative Media Information",
		},
		{
			NRP:            "5025221114",
			DepartmentName: "External Affairs",
		},
		{
			NRP:            "5025231007",
			DepartmentName: "External Affairs",
		},
		{
			NRP:            "5025231189",
			DepartmentName: "External Affairs",
		},
		{
			NRP:            "5025221096",
			DepartmentName: "External Affairs",
		},
		{
			NRP:            "5025221040",
			DepartmentName: "External Affairs",
		},
		{
			NRP:            "5025231195",
			DepartmentName: "External Affairs",
		},
		{
			NRP:            "5025231147",
			DepartmentName: "External Affairs",
		},
		{
			NRP:            "5025231116",
			DepartmentName: "External Affairs",
		},
		{
			NRP:            "5054231007",
			DepartmentName: "External Affairs",
		},
		{
			NRP:            "5025231031",
			DepartmentName: "External Affairs",
		},
		{
			NRP:            "5025231255",
			DepartmentName: "External Affairs",
		},
		{
			NRP:            "5025231021",
			DepartmentName: "External Affairs",
		},
		{
			NRP:            "5025231262",
			DepartmentName: "External Affairs",
		},
		{
			NRP:            "5025231245",
			DepartmentName: "External Affairs",
		},
		{
			NRP:            "5025221257",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5025231301",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5025231304",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5054231006",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5053231023",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5025231237",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5025231276",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5025231267",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5025231194",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5025231204",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5025231248",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5053231019",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5025231311",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5025231086",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5053231018",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5025231100",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5025231089",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5025231183",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5053231012",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5025231269",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5025231151",
			DepartmentName: "Internal Affairs",
		},
		{
			NRP:            "5025221044",
			DepartmentName: "Board Of Directors",
		},
		{
			NRP:            "5025221177",
			DepartmentName: "Board Of Directors",
		},
		{
			NRP:            "5025221159",
			DepartmentName: "Board Of Directors",
		},
		{
			NRP:            "5025221005",
			DepartmentName: "Board Of Directors",
		},
		{
			NRP:            "5025231058",
			DepartmentName: "Board Of Directors",
		},
		{
			NRP:            "5025221056",
			DepartmentName: "Board Of Directors",
		},
		{
			NRP:            "5053231003",
			DepartmentName: "Board Of Directors",
		},
		{
			NRP:            "5025231182",
			DepartmentName: "Entrepreneurship Development Departement",
		},
		{
			NRP:            "5025231253",
			DepartmentName: "Entrepreneurship Development Departement",
		},
		{
			NRP:            "5025231080",
			DepartmentName: "Entrepreneurship Development Departement",
		},
		{
			NRP:            "5025231005",
			DepartmentName: "Entrepreneurship Development Departement",
		},
		{
			NRP:            "5025231212",
			DepartmentName: "Entrepreneurship Development Departement",
		},
		{
			NRP:            "5025231177",
			DepartmentName: "Entrepreneurship Development Departement",
		},
		{
			NRP:            "5054231009",
			DepartmentName: "Entrepreneurship Development Departement",
		},
		{
			NRP:            "5025231315",
			DepartmentName: "Entrepreneurship Development Departement",
		},
		{
			NRP:            "5025231313",
			DepartmentName: "Entrepreneurship Development Departement",
		},
		{
			NRP:            "5025231191",
			DepartmentName: "Entrepreneurship Development Departement",
		},
		{
			NRP:            "5025231106",
			DepartmentName: "Entrepreneurship Development Departement",
		},
		{
			NRP:            "5025231074",
			DepartmentName: "Entrepreneurship Development Departement",
		},
		{
			NRP:            "5025221141",
			DepartmentName: "Student Resources Development",
		},
		{
			NRP:            "5025231039",
			DepartmentName: "Student Resources Development",
		},
		{
			NRP:            "5025221294",
			DepartmentName: "Student Resources Development",
		},
		{
			NRP:            "5025221055",
			DepartmentName: "Student Resources Development",
		},
		{
			NRP:            "5025221274",
			DepartmentName: "Student Resources Development",
		},
		{
			NRP:            "5025231142",
			DepartmentName: "Student Resources Development",
		},
		{
			NRP:            "5025231107",
			DepartmentName: "Student Resources Development",
		},
		{
			NRP:            "5025231186",
			DepartmentName: "Student Resources Development",
		},
		{
			NRP:            "5025231062",
			DepartmentName: "Student Resources Development",
		},
		{
			NRP:            "5053231011",
			DepartmentName: "Student Resources Development",
		},
		{
			NRP:            "5025231206",
			DepartmentName: "Student Resources Development",
		},
		{
			NRP:            "5025231188",
			DepartmentName: "Student Resources Development",
		},
		{
			NRP:            "5025231143",
			DepartmentName: "Student Resources Development",
		},
		{
			NRP:            "5025231101",
			DepartmentName: "Student Resources Development",
		}, {
			NRP:            "5025231307",
			DepartmentName: "Student Social Development",
		},
		{
			NRP:            "5025231041",
			DepartmentName: "Student Social Development",
		},
		{
			NRP:            "5025221236",
			DepartmentName: "Student Social Development",
		},
		{
			NRP:            "5025221280",
			DepartmentName: "Student Social Development",
		},
		{
			NRP:            "5025221315",
			DepartmentName: "Student Social Development",
		},
		{
			NRP:            "5025231111",
			DepartmentName: "Student Social Development",
		},
		{
			NRP:            "5025231215",
			DepartmentName: "Student Social Development",
		},
		{
			NRP:            "5025231205",
			DepartmentName: "Student Social Development",
		},
		{
			NRP:            "5025231259",
			DepartmentName: "Student Social Development",
		},
		{
			NRP:            "5025231128",
			DepartmentName: "Student Social Development",
		},
		{
			NRP:            "5025231312",
			DepartmentName: "Student Social Development",
		},
		{
			NRP:            "5025231063",
			DepartmentName: "Student Social Development",
		},
		{
			NRP:            "5025231292",
			DepartmentName: "Student Social Development",
		},
		{
			NRP:            "5025231219",
			DepartmentName: "Student Social Development",
		},
		{
			NRP:            "5025221012",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5025231178",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5025221054",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5025231075",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5025221307",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5025221288",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5025231309",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5054231001",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5025231236",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5054231016",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5025231046",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5025231156",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5025231243",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5025231303",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5025231249",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5025231029",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5025231223",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5025231016",
			DepartmentName: "Student Talenta and Interests",
		},
		{
			NRP:            "5025221101",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025231235",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025221218",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025231032",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025221106",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5053231004",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025231095",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025231088",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025231043",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025231051",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025231268",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025231275",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025231040",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025231096",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025231220",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025231256",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5053231021",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025231013",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025231208",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025231209",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
		{
			NRP:            "5025231265",
			DepartmentName: "Student Welfare, Research, and Technology",
		},
	}
	var seeded, skipped, failed int

	for _, userData := range usersToSeed {
		deptId, ok := departmentMap[userData.DepartmentName]
		if !ok {
			fmt.Printf("Warning: Department '%s' not found, skipping user '%s'\n",
				userData.DepartmentName, userData.NRP)
			skipped++
			continue
		}

		var existingUser model.User
		result := db.Where("nrp = ?", userData.NRP).First(&existingUser)

		if result.Error == nil {
			fmt.Printf("User already exists with NRP '%s' (skipping)\n", userData.NRP)
			skipped++
			continue
		}

		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Printf("Error checking for existing user with NRP '%s': %v\n", userData.NRP, result.Error)
			failed++
			continue
		}

		user := model.User{
			Id:            uuid.New(),
			NRP:           userData.NRP,
			DepartementId: &deptId,
			Timestamp: model.Timestamp{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		createResult := db.Create(&user)
		if createResult.Error != nil {
			fmt.Printf("Failed to create user '%s': %v\n", user.NRP, createResult.Error)
			failed++
		} else {
			fmt.Printf("Created new user: %s\n", user.NRP)
			seeded++
		}
	}

	fmt.Printf("Users: %d created, %d skipped, %d failed\n", seeded, skipped, failed)

	if failed > 0 && seeded == 0 && len(usersToSeed) > 0 {
		return fmt.Errorf("all user creation attempts failed")
	}

	return nil
}
