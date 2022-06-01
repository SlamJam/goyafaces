package bootstrap

import (
	"context"
	"log"
	"os"

	"github.com/SlamJam/goyafaces/pkg/app/bootstrap/providers"
	"github.com/SlamJam/goyafaces/pkg/app/config"
	"go.uber.org/dig"
)

type Invoker interface {
	// Invoke(function interface{}, opts ...dig.InvokeOption) error
	MustInvoke(function interface{}, opts ...dig.InvokeOption)
}

func AppContainer(ctx context.Context) Invoker {
	c := NewContainer()

	// Root context for all
	c.MustProvide(func() context.Context {
		return ctx
	})
	c.MustProvide(providers.ProvideConfig)
	c.MustProvide(func(cfg *config.Config) *log.Logger {
		return log.New(os.Stdout, cfg.Foo.Bar, 0)
	})

	return c
}
