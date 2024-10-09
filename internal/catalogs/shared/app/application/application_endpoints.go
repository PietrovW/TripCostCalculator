package application

import (
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/contracts"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func (a *Application) MapEndpoints() {
	endpoints := a.Container.Get("routes").([]contracts.Endpoint)
	for _, endpoint := range endpoints {
		endpoint.MapEndpoint()
	}

	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Catalogs Write-Service Api"
	docs.SwaggerInfo.Description = "Catalogs Write-Service Api."

	a.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
