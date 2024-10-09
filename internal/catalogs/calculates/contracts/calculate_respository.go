package contracts

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/models"
)

type CalculateRepository interface {
	CreateCalculate(ctx context.Context, product *models.Calculate) (*models.Calculate, error)
	GetCalculateById(ctx context.Context, uuid uuid.UUID) (*models.Calculate, error)
}
