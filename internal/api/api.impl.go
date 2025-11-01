package api

import (
	"cdd-go-boilerplate/internal/entity"
	"cdd-go-boilerplate/internal/module"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golobby/container/v3"
	"github.com/labstack/echo/v4"
)

type apiServer struct {
	validate    *validator.Validate `container:"type"`
	dummyModule module.DummyModule  `container:"type"`
}

func FillApiServer(c container.Container) (ServerInterface, error) {
	o := new(apiServer)
	err := c.Fill(o)
	return o, err
}

func (a apiServer) HealthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, entity.HealthCheckResponse{
		Healthy: true,
	})
}

func (a apiServer) DummyEndpoint(ctx echo.Context, params entity.DummyEndpointParams) error {
	err := validateStruct(a.validate, params)
	if err != nil {
		return err
	}

	res, err := a.dummyModule.Dummy(ctx.Request().Context(), *params.Type)
	if err != nil {
		return err
	}
	return sendSuccessResponse(ctx, res)
}

func (a apiServer) DummyEndpointPost(ctx echo.Context) error {
	param, err := bindAndValidate[entity.DummyEndpointPostJSONBody](ctx, a.validate)
	if err != nil {
		return err
	}
	dummy, err := a.dummyModule.Dummy(ctx.Request().Context(), *param.Type)
	if err != nil {
		return err
	}
	return sendSuccessResponse(ctx, dummy)
}
