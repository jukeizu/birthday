package birthday

import (
	"fmt"

	"github.com/jukeizu/handler"
)

type Config struct {
	Message string
}

type Job interface {
	handler.Job
}

type job struct {
	Config Config
}

func NewJob(config Config) Job {
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
