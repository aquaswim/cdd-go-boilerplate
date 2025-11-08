package utils

import (
	"github.com/golobby/container/v3"
)

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

type AutoInjectConstructor[T any] = func(c container.Container) (T, error)

func SingletonWithAutoInject[T any](c container.Container, fn AutoInjectConstructor[T]) {
	container.MustSingleton(c, func() (T, error) {
		return fn(c)
	})
}

func Fill[T any](c container.Container) (*T, error) {
	t := new(T)
	err := c.Fill(t)
	if err != nil {
		return nil, err
	}
	return t, nil
}
