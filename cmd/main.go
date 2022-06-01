package main

import (
	"log"

	"context"

	"github.com/SlamJam/goyafaces/pkg/app/bootstrap"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	ctx := context.TODO()
	c := bootstrap.AppContainer(ctx)

	var command *log.Logger
	c.MustInvoke(func(l *log.Logger) {
		l.Print("You've been invoked")
		command = l
	})

	spew.Dump(command)
}
