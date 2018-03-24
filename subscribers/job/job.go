package job

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/jukeizu/handler"
)

type Config struct {
	HandlerConfig handler.HandlerConfig
	Message       string
}

type Job interface {
	handler.Job
}

type job struct {
	Logger log.Logger
	Config Config
}

func NewJob(logger log.Logger, config Config) Job {
	logger = log.With(logger, "component", "birthday.job")
	return &job{logger, config}
}

func (j *job) IsJob(request handler.JobRequest) (bool, error) {
	return request.Type == "birthday", nil
}

func (j *job) Handle(request handler.JobRequest) (handler.Results, error) {
	message := fmt.Sprintf("<@!%s> %s", request.User, j.Config.Message)

	result := handler.Result{
		Content: message,
	}

	return handler.Results{result}, nil
}

func (j *job) Start() error {
	handler, err := handler.NewJobHandler(j.Logger, j.Config.HandlerConfig)

	if err != nil {
		return err
	}

	handler.Start(j)

	return nil
}
