package main

import (
	"net/http"
	"sync"

	"github.com/go-kit/kit/log"
	"github.com/jukeizu/birthday/services/birthday"
	"github.com/jukeizu/clients/scheduler/scheduler"
	"github.com/shawntoffel/services-core/runner"
)

func StartServices(wg *sync.WaitGroup, logger log.Logger, config Config) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		storage, err := birthday.NewStorage(config.DbConfig)
		defer storage.Close()

		schedulingClient := schedulingclient.NewClient(config.SchedulingClient)

		service := birthday.NewService(storage, schedulingClient)
		service = birthday.NewLoggingService(logger, service)

		if err != nil {
			panic(err)
		}

		httpLogger := log.With(logger, "component", "http")

		mux := http.NewServeMux()

		birthdayHandler := birthday.MakeHandler(service, httpLogger)
		mux.Handle("/birthday", birthdayHandler)

		runner.StartService(mux, logger, config.ServiceConfig)
	}()

}
