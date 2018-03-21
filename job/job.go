package job

import (
	"github.com/go-kit/kit/log"
	"github.com/jukeizu/birthday/job/birthday"
	"github.com/jukeizu/handler"
)

type Config struct {
	HandlerConfig     handler.HandlerConfig
	BirthdayJobConfig birthday.Config
}

func Start(logger log.Logger, config Config) error {
	logger = log.With(logger, "component", "birthday.job")

	handler, err := handler.NewJobHandler(logger, config.HandlerConfig)

	if err != nil {
		return err
	}

	birthdayJob := birthday.NewJob(config.BirthdayJobConfig)

	handler.Start(birthdayJob)

	return nil
}
