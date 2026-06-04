package db

import (
	"log"
	"weather-app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("weather.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Couldn't connect to the database:", err)
	}

	DB.AutoMigrate(&models.Weather{}, &models.Forecast{})

	seed()
}

func seed() {
	var count int64
	DB.Model(&models.Weather{}).Count(&count)

	if count == 0 {
		log.Println("Seeding database...")
		
		initialData := []models.Weather{
			{Location: "Wieliczka", Temperature: 22.5, Precipitation: 0.0, Description: "Pochmurno"},
			{Location: "Kielce", Temperature: 20.0, Precipitation: 8.0 ,Description: "Lekki wiatr"},
			{Location: "Kraków", Temperature: 22.1, Precipitation: 0.0 ,Description: "Słonecznie"},
		}

		for _, w := range initialData {
			DB.Create(&w)
		}
	} else {
		log.Println("Seeding skipped.")
	}
}