package commands

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreateCalculateCommand struct {
	CaculateID      uuid.UUID `validate:"required"`
	Distance        float64   `validate:"required"`
	FuelConsumption float64   `validate:"required"`
	FuelPrice       float64   `validate:"required"`
	Passengers      int       `validate:"required"`
	CreatedAt       time.Time `validate:"required"`
}

func NewCreateCalculateCommand(distance float64, fuelConsumption float64, fuelPrice float64, passengers int) *CreateCalculateCommand {
	return &CreateCalculateCommand{CaculateID: uuid.NewV4(), Distance: distance, FuelConsumption: fuelConsumption, FuelPrice: fuelPrice, Passengers: passengers, CreatedAt: time.Now()}
}
