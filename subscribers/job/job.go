package job

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/jukeizu/handler"
)

type Config struct {
	HandlerConfig     handler.HandlerConfig
	BirthdayJobConfig JobConfig
}

type JobConfig struct {
	Message string
}

type Job interface {
	handler.Job
}

type job struct {
	Config JobConfig
}

func NewJob(config JobConfig) Job {
	return &job{config}
}

func (j *job) IsJob(request handler.JobRequest) (bool, error) {
	return request.Type == "birthday", nil
}

func (j *job) Run(request handler.JobRequest) (handler.Results, error) {
	message := fmt.Sprintf("<@!%s> %s", request.User, j.Config.Message)

	result := handler.Result{
		Content: message,
	}

	return handler.Results{result}, nil
}

func Start(logger log.Logger, config Config) error {
	logger = log.With(logger, "component", "birthday.job")

	handler, err := handler.NewJobHandler(logger, config.HandlerConfig)

	if err != nil {
		return err
	}

	birthdayJob := NewJob(config.BirthdayJobConfig)

	handler.Start(birthdayJob)

	return nil
}
