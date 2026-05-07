package seeding

import (
	"errors"
	"fmt"
	"time"

	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GallerySeeder struct {
	galleries []model.Gallery
}

func NewGallerySeeder() *GallerySeeder {
	layout := "2006-01-02"
	t, err := time.Parse(layout, "2026-03-15")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &GallerySeeder{
		galleries: []model.Gallery{
			{
				Id:           uuid.New(),
				Title:        "Dokumentasi Inagurasi 2026",
				Description:  "Kumpulan foto dan video malam puncak inagurasi HMTC.",
				GDriveLink:   "https://drive.google.com/drive/folders/...",
				ThumbnailUrl: "https://storage.hmtc.com/thumb/inagurasi.jpg",
				EventDate: model.Timestamp{
					CreatedAt: t,
					UpdatedAt: t,
				},
				Timestamp: model.Timestamp{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
		},
	}

}

func (s *GallerySeeder) GetName() string {
	return "GallerySeeder"
}

func (s *GallerySeeder) Seed(db *gorm.DB, options SeedOptions) error {
	var count int64
	db.Model(&model.Gallery{}).Count(&count)

	switch options.Mode {
	case Skip:
		if count > 0 {
			fmt.Println("Galleries already exist, skipping...")
			return nil
		}

	case Replace:
		if count > 0 {
			fmt.Println("Clearing existing galleries...")
			if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Gallery{}).Error; err != nil {
				return fmt.Errorf("failed to clear galleries: %w", err)
			}
		}

	case Append:
		fmt.Println("Appending galleries...")
	}

	var seeded, skipped, failed int

	for _, gallery := range s.galleries {
		var existingGallery model.Gallery
		result := db.Where("title = ?", gallery.Title).First(&existingGallery)

		if result.Error == nil {
			fmt.Printf("Gallery already exists: %s (skipping)\n", gallery.Title)
			skipped++
			continue
		}

		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Printf("Error checking for existing gallery '%s': %v\n", gallery.Title, result.Error)
			failed++
			continue
		}

		createResult := db.Create(&gallery)
		if createResult.Error != nil {
			fmt.Printf("Failed to create gallery '%s': %v\n", gallery.Title, createResult.Error)
			failed++
		} else {
			fmt.Printf("Created new gallery: %s\n", gallery.Title)
			seeded++
		}
	}

	fmt.Printf("Galleries: %d created, %d skipped, %d failed\n", seeded, skipped, failed)

	if failed > 0 && seeded == 0 && len(s.galleries) > 0 {
		return fmt.Errorf("all gallery creation attempts failed")
	}

	return nil
}
