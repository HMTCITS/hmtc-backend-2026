package seeding

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/gorm"
)

type SeedMode string

const (
	Skip    SeedMode = "skip"
	Append  SeedMode = "append"
	Replace SeedMode = "replace"
)

type SeedOptions struct {
	Mode SeedMode
}

func DefaultSeedOptions() SeedOptions {
	mode := Skip

	envMode := strings.ToLower(os.Getenv("SEED_MODE"))
	switch envMode {
	case "skip":
		mode = Skip
	case "append":
		mode = Append
	case "replace":
		mode = Replace
	}

	return SeedOptions{
		Mode: mode,
	}
}

type Seeder interface {
	Seed(*gorm.DB, SeedOptions) error
	GetName() string
}

func RunAllSeeders(db *gorm.DB, options SeedOptions, seeders ...Seeder) error {
	for _, seeder := range seeders {
		log.Printf("Running seeder: %s (mode: %s)", seeder.GetName(), options.Mode)
		if err := seeder.Seed(db, options); err != nil {
			return fmt.Errorf("error in %s seeder: %w", seeder.GetName(), err)
		}
		log.Printf("Seeder %s completed successfully", seeder.GetName())
	}
	return nil
}

func GetSeeders() []Seeder {
	return []Seeder{
		NewDepartementSeeder(),
		NewUserSeeder(),
		NewAdminSeeder(),
	}
}
