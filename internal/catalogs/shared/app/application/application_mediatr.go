package application

import (
	"github.com/mehdihadeli/go-mediatr"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/contracts"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/features/creating_calculate/commands"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/calculates/features/creating_calculate/dtos"
	"github.com/pietrovw/trip-cost-calculator/internal/catalogs/shared/behaviours"
)

func (a *Application) ConfigMediator() error {
	calculateRepository := a.Container.Get("calculateRepository").(contracts.CalculateRepository)

	loggerPipeline := &behaviours.RequestLoggerBehaviour{}
	err := mediatr.RegisterRequestPipelineBehaviors(loggerPipeline)

	if err != nil {
		return err
	}

	createCalculateCommandHandler := commands.NewCreateProductCommandHandler(calculateRepository)
	err = mediatr.RegisterRequestHandler[*commands.CreateCalculateCommand, *dtos.CreateCalculateCommandResponse](createCalculateCommandHandler)
	if err != nil {
		return err
	}

	//getByIdQueryHandler := queries.NewGetCalculateByIdHandler(calculateRepository)
	////	err = mediatr.RegisterRequestHandler[*queries.GetCalculateByIdQuery, *dtos2.GetCalculateByIdQueryResponse](getByIdQueryHandler)
	//if err != nil {
	//		return err
	//	}

	//notificationHandler := events.NewCalculateCreatedEventHandler()
	//	err = mediatr.RegisterNotificationHandler[*events.CalculateCreatedEvent](notificationHandler)
	//	if err != nil {
	//		return err
	//	}
	return nil
}
