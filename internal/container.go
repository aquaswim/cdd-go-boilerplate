package internal

import (
	"cdd-go-boilerplate/internal/api"
	"cdd-go-boilerplate/internal/config"
	"cdd-go-boilerplate/internal/module"
	globalLogger "cdd-go-boilerplate/internal/pkg/global_logger"
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

	container.MustSingleton(container.Global, module.NewDummyModule)

	container.MustSingleton(container.Global, api.NewApiServer)
	container.MustSingleton(container.Global, api.NewEchoServer)

	return container.Global
}

func Resolve[T any](c container.Container) T {
	var t T
	container.MustResolve(c, &t)
	return t
}

func ResolveNamed[T any](c container.Container, name string) T {
	var t T
	container.MustNamedResolve(c, &t, name)
	return t
}
