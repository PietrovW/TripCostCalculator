package applicationbuilder

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/pietrovw/trip-cost-calculator/config"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/contracts"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/contracts/params"
	endpoints2 "github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/features/creating_calculate/endpoints"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/repository"
	config2 "github.com/pietrovw/trip-cost-calculator/internal/pkg/config"
	"github.com/pietrovw/trip-cost-calculator/internal/pkg/config/environemnt"
	"github.com/pietrovw/trip-cost-calculator/internal/pkg/database"
	"github.com/sarulabs/di"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (b *ApplicationBuilder) AddCore() {

	logDep := di.Def{
		Name:  "zap",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return b.Logger, nil
		},
	}

	configDep := di.Def{
		Name:  "config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			env := ctn.Get("env").(environemnt.Environment)
			return config.NewConfig(env)
		},
	}

	err := config2.AddEnv(b.Services)
	if err != nil {
		b.Logger.Fatal(err)
	}

	err = b.Services.Add(logDep)
	if err != nil {
		b.Logger.Fatal(err)
	}

	err = b.Services.Add(configDep)
	if err != nil {
		b.Logger.Fatal(err)
	}
}

func (b *ApplicationBuilder) AddInfrastructure() {
	err := database.AddGorm(b.Services)
	if err != nil {
		b.Logger.Fatal(err)
	}

	err = addEcho(b.Services)
	if err != nil {
		b.Logger.Fatal(err)
	}
}

func (b *ApplicationBuilder) AddRoutes() {
	routesDep := di.Def{
		Name:  "routes",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			calculateRouteParams := ctn.Get("productRouteGroup").(*params.CalculateRouteParams)

			createProductEndpoint := endpoints2.NewCreteCalculateEndpoint(calculateRouteParams)
			//	getProductById := endpoints3.NewGetCalculateByIdEndpoint(calculateRouteParams)
			endpoints := []contracts.Endpoint{createProductEndpoint} //, getProductById}

			return endpoints, nil
		},
	}

	err := b.Services.Add(routesDep)
	if err != nil {
		b.Logger.Fatal(err)
	}
}

func (b *ApplicationBuilder) AddRepositories() {
	calculateRepositoryDep := di.Def{
		Name:  "calculateRepository",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			g := ctn.Get("gorm").(*gorm.DB)
			l := ctn.Get("zap").(*zap.SugaredLogger)
			return repository.NewInMemoryCalculateRepository(g, l), nil
		},
	}

	err := b.Services.Add(calculateRepositoryDep)
	if err != nil {
		b.Logger.Fatal(err)
	}
}

func addEcho(container *di.Builder) error {
	echoDep := di.Def{
		Name:  "echo",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return echo.New(), nil
		},
	}

	productGroupDep := di.Def{
		Name:  "productRouteGroup",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			echo := ctn.Get("echo").(*echo.Echo)
			logger := ctn.Get("zap").(*zap.SugaredLogger)

			v1 := echo.Group("/api/v1")
			calculates := v1.Group("/calculates")

			calculatesRouteParams := &params.CalculateRouteParams{
				Logger:          logger,
				Validator:       validator.New(),
				CalculatesGroup: calculates,
			}

			return calculatesRouteParams, nil
		},
	}

	err := container.Add(echoDep)
	if err != nil {
		return err
	}

	err = container.Add(productGroupDep)
	if err != nil {
		return err
	}

	return nil
}
