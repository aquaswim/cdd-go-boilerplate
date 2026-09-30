package api

import (
	"cdd-go-boilerplate/internal/config"
	globalLogger "cdd-go-boilerplate/internal/pkg/global_logger"
	"cdd-go-boilerplate/internal/pkg/utils"
	"context"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-playground/validator/v10"
	"github.com/golobby/container/v3"
	"github.com/labstack/echo/v5"
	echoMiddleware "github.com/labstack/echo/v5/middleware"
	oapiMiddleware "github.com/oapi-codegen/echo-v5-middleware"
	"github.com/rs/zerolog/log"
)

type Server interface {
	Start(ctx context.Context) error
}

type echoServer struct {
	echo     *echo.Echo
	cfg      *config.Config      `container:"type"`
	api      ServerInterface     `container:"type"`
	validate *validator.Validate `container:"type"`
}

func FillEchoServer(c container.Container) (Server, error) {
	svr, err := utils.Fill[echoServer](c)
	if err != nil {
		return nil, err
	}

	// setup echo server
	svr.echo = echo.New()
	svr.echo.Validator = svr.newValidator()
	svr.echo.HTTPErrorHandler = ErrorHandler()
	svr.echo.Logger = globalLogger.SlogAdapter()

	svr.echo.Use(echoMiddleware.RequestID())
	svr.echo.Use(echoMiddleware.Recover())
	svr.echo.Use(LoggerMiddleware(&log.Logger))

	// setup oapi handlers
	swagger, err := GetSpec()
	if err != nil {
		panic(err)
	}
	// remove server value from swagger
	swagger.Servers = nil
	svr.echo.Use(oapiMiddleware.OapiRequestValidatorWithOptions(swagger, &oapiMiddleware.Options{
		Options: openapi3filter.Options{
			// exclude all validation since we want to use go-validator
			ExcludeRequestBody:        true,
			ExcludeRequestQueryParams: true,
			// note: register auth function here
			//AuthenticationFunc: authentication(authModule),
		},
	}))
	RegisterHandlers(svr.echo, svr.api)

	return svr, nil
}

func (e echoServer) Start(ctx context.Context) error {
	sc := echo.StartConfig{
		Address: e.cfg.ListenAddr,
	}

	return sc.Start(ctx, e.echo)
}
