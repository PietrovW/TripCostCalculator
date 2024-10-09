package database

import (
	"github.com/pietrovw/trip-cost-calculator/internal/pkg/database/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(options *options.GormOptions) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(options.Dns()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
