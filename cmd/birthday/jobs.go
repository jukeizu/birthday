package main

import (
	"sync"

	"github.com/go-kit/kit/log"

	"github.com/jukeizu/birthday/subscribers/job"
)

func StartJobs(wg *sync.WaitGroup, logger log.Logger, config Config) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		j := job.NewJob(logger, config.JobConfig)

		err := j.Start()

		if err != nil {
			panic(err)
		}
	}()
}
