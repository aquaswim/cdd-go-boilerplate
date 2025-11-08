package module

import (
	appErrors "cdd-go-boilerplate/internal/app_errors"
	"cdd-go-boilerplate/internal/pkg/utils"
	"context"

	"github.com/golobby/container/v3"
	"github.com/joomcode/errorx"
	"github.com/rs/zerolog"
)

type DummyModule interface {
	Dummy(ctx context.Context, paramType string) (interface{}, error)
}

type dummyModule struct {
}

func FillDummyModule(c container.Container) (DummyModule, error) {
	return utils.Fill[dummyModule](c)
}

func (d dummyModule) Dummy(ctx context.Context, paramType string) (interface{}, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msgf("Dummy endpoint called with type: %s", paramType)

	switch paramType {
	case "400":
		return nil, appErrors.ErrTypeValidation.New("dummy error").WithProperty(errorx.PropertyPayload(), map[string]any{
			"param_type": paramType,
		})
	case "401":
		return nil, appErrors.ErrTypeUnauthorized.New("dummy error")
	case "403":
		return nil, appErrors.ErrTypeForbidden.New("dummy error")
	case "404":
		return nil, appErrors.ErrTypeNotFound.New("dummy error")
	case "500":
		return nil, appErrors.ErrTypeInternal.New("dummy error")
	default:
		return map[string]any{
			"stuff": "lorem ipsum",
		}, nil
	}
}
