package config

import (
	"github.com/pietrovw/trip-cost-calculator/internal/pkg/config/environemnt"
	"github.com/sarulabs/di"
)

func AddEnv(container *di.Builder) error {
	envDep := di.Def{
		Name:  "env",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return environemnt.ConfigAppEnv(), nil
		},
	}
	return container.Add(envDep)
}
