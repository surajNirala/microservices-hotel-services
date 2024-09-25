package databases

import (
	"log"

	"github.com/surajNirala/hotel_services/app/config"
	"github.com/surajNirala/hotel_services/app/models"
)

func DatabaseUp() {
	DB := config.DB
	err := DB.AutoMigrate(
		&models.Hotel{},
	)
	if err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	}
}
