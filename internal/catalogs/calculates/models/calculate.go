package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Calculate struct {
	CalculateID     uuid.UUID `json:"productId" gorm:"primaryKey"`
	Distance        float64   `json:"distance"`
	FuelConsumption float64   `json:"fuel_consumption"`
	FuelPrice       float64   `json:"fuel_price"`
	Passengers      int       `json:"passengers"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
