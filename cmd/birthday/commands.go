package main

import (
	"sync"

	"github.com/go-kit/kit/log"
	"github.com/jukeizu/birthday/api"
	"github.com/jukeizu/birthday/subscribers/commands"
	"github.com/jukeizu/handler"
)

func StartCommands(wg *sync.WaitGroup, logger log.Logger, config CommandConfig) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		client := api.NewClient(config.BirthdayClient)
		c := commands.NewCommand(logger, client)

		logger = log.With(logger, "command", "bithday.set")

		handler, err := handler.NewCommandHandler(logger, config.SetHandlerConfig)

		if err != nil {
			panic(err)
		}

		handler.Start(c.Set())
	}()
}
