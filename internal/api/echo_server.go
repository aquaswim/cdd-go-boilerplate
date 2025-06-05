package api

import (
	"context"
	"errors"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	oapiMiddleware "github.com/oapi-codegen/echo-middleware"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type Server interface {
	Start() error
	Stop() error
}

type echoServer struct {
	echo *echo.Echo
}

func NewEchoServer(api ServerInterface) Server {
	svr := &echoServer{}

	// setup echo server
	svr.echo = echo.New()
	svr.echo.HideBanner = true
	svr.echo.HTTPErrorHandler = ErrorHandler()

	svr.echo.Use(echoMiddleware.RequestID())
	svr.echo.Use(LoggerMiddleware(&log.Logger))
	svr.echo.Use(RecoverMiddleware())

	// setup oapi handlers
	swagger, err := GetSwagger()
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

func (e echoServer) Start() error {
	err := e.echo.Start(":3000")

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func (e echoServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return e.echo.Shutdown(ctx)
}
