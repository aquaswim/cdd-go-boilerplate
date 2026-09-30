package api

import (
	"cdd-go-boilerplate/internal/entity"
	"cdd-go-boilerplate/internal/module"
	"cdd-go-boilerplate/internal/pkg/utils"
	"net/http"

	"github.com/golobby/container/v3"
	"github.com/labstack/echo/v5"
)

//go:generate go tool oapi-codegen -config ../../api/generate-server.config.yaml -o api.gen.go ../../api/api.yaml
type apiServer struct {
	dummyModule module.DummyModule `container:"type"`
}

func FillApiServer(c container.Container) (ServerInterface, error) {
	return utils.Fill[apiServer](c)
}

func (a apiServer) HealthCheck(ctx *echo.Context) error {
	return ctx.JSON(http.StatusOK, entity.HealthCheckResponse{
		Healthy: true,
	})
}

func (a apiServer) DummyEndpoint(ctx *echo.Context, params entity.DummyEndpointParams) error {
	err := ctx.Validate(params)
	if err != nil {
		return err
	}

	res, err := a.dummyModule.Dummy(ctx.Request().Context(), *params.Type)
	if err != nil {
		return err
	}
	return sendSuccessResponse(ctx, res)
}

func (a apiServer) DummyEndpointPost(ctx *echo.Context) error {
	param, err := bindAndValidate[entity.DummyEndpointPostJSONBody](ctx)
	if err != nil {
		return err
	}
	dummy, err := a.dummyModule.Dummy(ctx.Request().Context(), *param.Type)
	if err != nil {
		return err
	}
	return sendSuccessResponse(ctx, dummy)
}
