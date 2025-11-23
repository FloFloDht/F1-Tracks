package db

import (
	"FT26/internal/tracks/models"
	"log"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	// Auto-migrations : création des tables si absentes
	db.AutoMigrate(&models.Track{})

	// Vérifie si la table est vide
	var count int64
	db.Model(&models.Track{}).Count(&count)
	if count > 0 {
		log.Println("Database already seeded")
		return
	}

	// Données de départ
	tracks := []models.Track{
		{Name: "Monza", Country: "Italie", LengthKm: 5.793},
		{Name: "Silverstone", Country: "Royaume-Uni", LengthKm: 5.891},
		{Name: "Spa-Francorchamps", Country: "Belgique", LengthKm: 7.004},
	}

	if err := db.Create(&tracks).Error; err != nil {
		log.Fatalf("Error seeding database: %v", err)
	}

	log.Println("Seeding completed")
}
