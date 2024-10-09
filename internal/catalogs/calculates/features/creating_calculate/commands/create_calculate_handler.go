package commands

import (
	"context"

	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/contracts"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/features/creating_calculate/dtos"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/models"
)

type CreateCalculateCommandHandler struct {
	calculateRepository contracts.CalculateRepository
}

func NewCreateProductCommandHandler(calculateRepository contracts.CalculateRepository) *CreateCalculateCommandHandler {
	return &CreateCalculateCommandHandler{calculateRepository: calculateRepository}
}

func (c *CreateCalculateCommandHandler) Handle(ctx context.Context, command *CreateCalculateCommand) (*dtos.CreateCalculateCommandResponse, error) {

	calculate := &models.Calculate{
		CalculateID:     command.CaculateID,
		Distance:        command.Distance,
		FuelConsumption: command.FuelConsumption,
		FuelPrice:       command.FuelPrice,
		Passengers:      command.Passengers,
		CreatedAt:       command.CreatedAt,
	}

	createdCalculate, err := c.calculateRepository.CreateCalculate(ctx, calculate)
	if err != nil {
		return nil, err
	}

	response := &dtos.CreateCalculateCommandResponse{CalculateID: createdCalculate.CalculateID}

	// Publish notification event to the mediatr for dispatching to the notification handlers

	//productCreatedEvent := events.NewProductCreatedEvent(createdCalculate.CalculateID, createdCalculate.Name, createdCalculate.Description, createdProduct.Price, createdCalculate.CreatedAt)
	//err = mediatr.Publish[*events.ProductCreatedEvent](ctx, productCreatedEvent)
	//if err != nil {
	//	return nil, err
	//}

	return response, nil
}
