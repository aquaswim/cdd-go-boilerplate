package internal

import (
	"cdd-go-boilerplate/internal/api"
	"cdd-go-boilerplate/internal/config"
	"cdd-go-boilerplate/internal/module"
	globalLogger "cdd-go-boilerplate/internal/pkg/global_logger"
	"cdd-go-boilerplate/internal/pkg/utils"
	"cdd-go-boilerplate/internal/pkg/validation"

	"github.com/golobby/container/v3"
)

func InitContainer() container.Container {
	// register all module constructor here and let the container wire it
	container.MustSingleton(container.Global, config.Get)
	container.MustCall(container.Global, func(cfg *config.Config) {
		globalLogger.Setup(&globalLogger.Config{
			LogPretty: cfg.LogPretty,
			LogLevel:  cfg.LogLevel,
		})
	})
	container.MustSingleton(container.Global, validation.NewValidator)

	utils.SingletonWithAutoInject(container.Global, module.FillDummyModule)

	utils.SingletonWithAutoInject(container.Global, api.FillApiServer)

	container.MustSingleton(container.Global, api.NewEchoServer)

	return container.Global
}
