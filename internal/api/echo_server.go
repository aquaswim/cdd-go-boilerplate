package api

import (
	"context"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v5"
	echoMiddleware "github.com/labstack/echo/v5/middleware"
	oapiMiddleware "github.com/oapi-codegen/echo-v5-middleware"
	"github.com/rs/zerolog/log"
)

type Server interface {
	Start(ctx context.Context) error
}

type echoServer struct {
	echo *echo.Echo
}

func NewEchoServer(api ServerInterface) Server {
	svr := &echoServer{}

	// setup echo server
	svr.echo = echo.New()
	//svr.echo.HideBanner = true
	svr.echo.HTTPErrorHandler = ErrorHandler()

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
	RegisterHandlers(svr.echo, api)

	return svr
}

func (e echoServer) Start(ctx context.Context) error {
	sc := echo.StartConfig{
		Address:    ":3000",
		HideBanner: true,
	}

	return sc.Start(ctx, e.echo)
}
