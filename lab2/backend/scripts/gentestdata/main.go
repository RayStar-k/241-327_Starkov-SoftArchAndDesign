package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"guitarshop/internal/config"
	"guitarshop/internal/database"
	"guitarshop/internal/models"

	"github.com/jaswdr/faker"
)

func main() {
	log.Println("Starting test data generation...")

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if err := database.Connect(&cfg.Database); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fake := faker.New()
	rand.Seed(time.Now().UnixNano())

	guitars := generateGuitars(fake, 100)
	log.Printf("Generated %d guitars", len(guitars))

	log.Println("Test data generation completed successfully!")
}

func generateGuitars(fake faker.Faker, count int) []models.Guitar {
	brandNames := []string{
		"Fender", "Gibson", "Ibanez", "PRS", "Jackson",
		"ESP", "Schecter", "Gretsch", "Epiphone", "Washburn",
		"Yamaha", "Taylor", "Martin", "Guild", "Rickenbacker",
	}

	categories := []string{
		"Electric", "Acoustic", "Bass", "Classical",
		"Semi-Hollow", "Acoustic-Electric",
	}

	guitarModels := []string{
		"Stratocaster", "Telecaster", "Les Paul", "SG", "Flying V",
		"ES-335", "Jaguar", "Jazzmaster", "Thunderbird", "Precision",
		"Superstrat", "Dreadnought", "Grand Auditorium", "Jumbo",
	}

	colors := []string{
		"Sunburst", "Black", "White", "Red", "Blue",
		"Natural", "Cherry", "Vintage White", "Metallic Blue", "Purple",
	}

	var guitars []models.Guitar
	for i := 0; i < count; i++ {
		brand := brandNames[rand.Intn(len(brandNames))]
		category := categories[rand.Intn(len(categories))]

		guitar := models.Guitar{
			Model:            guitarModels[rand.Intn(len(guitarModels))],
			Brand:            brand,
			Category:         category,
			Price:            float64(rand.Intn(3000)+200) + 0.99,
			StringCount:      []int{4, 6, 7, 12}[rand.Intn(4)],
			Color:            colors[rand.Intn(len(colors))],
			SerialNumber:     fmt.Sprintf("%s-%d-%d", brand[:3], time.Now().Unix()+int64(i), rand.Intn(10000)),
			InStock:          rand.Float32() > 0.3,
			StockQuantity:    rand.Intn(20),
			YearManufactured: rand.Intn(25) + 2000,
			Description:      fake.Lorem().Sentence(15),
		}
		database.DB.Create(&guitar)
		guitars = append(guitars, guitar)
	}
	return guitars
}
