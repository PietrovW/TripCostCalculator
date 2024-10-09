package products

import (
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/dtos"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/models"
)

func MapCalculateToCalculateDto(calculate *models.Calculate) *dtos.CalculateDto {
	return &dtos.CalculateDto{
		CalculateID:     calculate.CalculateID,
		Distance:        calculate.Distance,
		FuelConsumption: calculate.FuelConsumption,
		FuelPrice:       calculate.FuelPrice,
		Passengers:      calculate.Passengers,
		UpdatedAt:       calculate.UpdatedAt,
		CreatedAt:       calculate.CreatedAt,
	}
}
