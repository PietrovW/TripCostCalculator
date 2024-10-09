package database

import (
	"github.com/pietrovw/trip-cost-calculator/internal/pkg/database/options"
	"github.com/sarulabs/di"
)

func AddGorm(container *di.Builder) error {
	gormOptionDep := di.Def{
		Name:  "gormOptions",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return options.ProvideConfig()
		},
	}

	gormDep := di.Def{
		Name:  "gorm",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			opt := ctn.Get("gormOptions").(*options.GormOptions)
			return NewGormDB(opt)
		},
	}

	err := container.Add(gormOptionDep)
	err = container.Add(gormDep)

	return err
}
