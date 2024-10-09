package params

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type CalculateRouteParams struct {
	Logger          *zap.SugaredLogger
	CalculatesGroup *echo.Group
	Validator       *validator.Validate
}
