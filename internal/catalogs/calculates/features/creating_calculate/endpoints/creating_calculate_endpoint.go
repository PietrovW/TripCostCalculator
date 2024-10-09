package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/contracts"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/contracts/params"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/features/creating_calculate/commands"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/features/creating_calculate/dtos"
)

type createCalculateEndpoint struct {
	*params.CalculateRouteParams
}

func NewCreteCalculateEndpoint(params *params.CalculateRouteParams) contracts.Endpoint {
	return &createCalculateEndpoint{CalculateRouteParams: params}
}

func (ep *createCalculateEndpoint) MapEndpoint() {
	ep.CalculatesGroup.POST("", ep.handler())
}

// CreateProduct
// @Tags Products
// @Summary Create product
// @Description Create new product item
// @Accept json
// @Produce json
// @Param CreateCalculateRequestDto body creatingProductsDtos.CreateProductRequestDto true "Product data"
// @Success 201 {object} creatingProductsDtos.CreateProductResponseDto
// @Router /api/v1/products [post]
func (ep *createCalculateEndpoint) handler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := &dtos.CreateCalculateRequestDto{}
		if err := ctx.Bind(request); err != nil {
			return err
		}

		if err := ep.Validator.StructCtx(ctx.Request().Context(), request); err != nil {
			return err
		}

		command := commands.NewCreateCalculateCommand(request.Distance, request.FuelConsumption, request.FuelPrice, request.Passengers)
		result, err := mediatr.Send[*commands.CreateCalculateCommand, *dtos.CreateCalculateCommandResponse](ctx.Request().Context(), command)

		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusCreated, result)
	}
}
