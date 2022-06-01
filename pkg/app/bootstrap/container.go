package bootstrap

import (
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Container struct {
	inner *dig.Container
}

func NewContainer() *Container {
	return &Container{
		inner: dig.New(),
	}
}

func (c *Container) MustProvide(constructor interface{}, opts ...dig.ProvideOption) {
	err := c.inner.Provide(constructor, opts...)
	err = errors.WithStack(err)
	if err != nil {
		panic(err)
	}
}

func (c *Container) MustInvoke(function interface{}, opts ...dig.InvokeOption) {
	err := c.inner.Invoke(function, opts...)
	err = errors.WithStack(err)
	if err != nil {
		panic(err)
	}
}

// func (c *Container) Invoke(function interface{}, opts ...dig.InvokeOption) error {
// 	err := c.Invoke(function, opts...)
// 	return errors.WithStack(err)
// }
