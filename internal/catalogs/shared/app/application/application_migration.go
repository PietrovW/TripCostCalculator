package application

import (
	"time"

	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func (a *Application) MigrateDatabase() error {
	db := a.Container.Get("gorm").(*gorm.DB)
	// Auto-migrate the Person model
	err := db.AutoMigrate(&models.Calculate{})
	if err != nil {
		return err
	}

	db.Create(&models.Calculate{Distance: 0, FuelConsumption: 0.0, FuelPrice: 0, Passengers: 1, CreatedAt: time.Now(), CalculateID: uuid.NewV4()})

	return nil
}
