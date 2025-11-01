package module

import (
	"cdd-go-boilerplate/internal/pkg/errorx"
	"context"

	"github.com/golobby/container/v3"
	"github.com/rs/zerolog"
)

type DummyModule interface {
	Dummy(ctx context.Context, paramType string) (interface{}, error)
}

type dummyModule struct {
}

func FillDummyModule(c container.Container) (DummyModule, error) {
	o := new(dummyModule)
	err := c.Fill(o)
	return o, err
}

func (d dummyModule) Dummy(ctx context.Context, paramType string) (interface{}, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msgf("Dummy endpoint called with type: %s", paramType)

	switch paramType {
	case "400":
		return nil, errorx.New(errorx.ErrCodeBadRequest, "Dummy Error")
	case "401":
		return nil, errorx.New(errorx.ErrCodeUnauthorized, "Dummy Error")
	case "403":
		return nil, errorx.New(errorx.ErrCodeForbidden, "Dummy Error")
	case "404":
		return nil, errorx.New(errorx.ErrCodeNotFound, "Dummy Error")
	case "500":
		return nil, errorx.New(errorx.ErrCodeInternal, "Dummy Error")
	default:
		return map[string]any{
			"stuff": "lorem ipsum",
		}, nil
	}
}
