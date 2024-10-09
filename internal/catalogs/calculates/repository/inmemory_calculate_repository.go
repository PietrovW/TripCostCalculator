package repository

import (
	"context"

	uuid "github.com/gofrs/uuid"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/contracts"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type InMemoryCalculateRepository struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewInMemoryCalculateRepository(db *gorm.DB, logger *zap.SugaredLogger) contracts.CalculateRepository {
	return &InMemoryCalculateRepository{db: db, logger: logger}
}

func (p *InMemoryCalculateRepository) CreateCalculate(ctx context.Context, calculate *models.Calculate) (*models.Calculate, error) {
	err := p.db.WithContext(ctx).Create(calculate).Error
	if err != nil {
		return nil, err
	}
	return calculate, nil
}

func (p *InMemoryCalculateRepository) GetCalculateById(ctx context.Context, uuid uuid.UUID) (*models.Calculate, error) {
	var calculate *models.Calculate
	err := p.db.WithContext(ctx).First(&calculate, "calculate_id = ?", uuid).Error
	if err != nil {
		return nil, err
	}
	return calculate, nil
}
