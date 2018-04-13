package commands

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	shellwords "github.com/mattn/go-shellwords"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/en"

	"github.com/jukeizu/birthday/api"
	"github.com/jukeizu/handler"
)

type Set interface {
	handler.Command
}

type set struct {
	command *command
	logger  log.Logger
}

func (c *command) Set() Set {
	logger := log.With(c.logger, "command", "birthday.set")

	return &set{c, logger}
}

func (s *set) IsCommand(request handler.Request) (bool, error) {
	return regexp.MatchString(`!setbirthday (.*?)`, request.Content)
}

func (s *set) Handle(request handler.Request) (handler.Results, error) {
	args, err := shellwords.Parse(request.Content)

	if err != nil {
		return handler.Results{}, err
	}

	input := strings.Join(args[1:], " ")

	parsed, err := s.parseTime(input)

	if err != nil || parsed == nil {
		errorMessage := handler.Result{
			Content: fmt.Sprintf("`%s` is not a valid date format.", input),
		}

		return handler.Results{errorMessage}, nil
	}

	month := parsed.Time.Month().String()
	day := parsed.Time.Day()

	setRequest := api.SetBirthdayRequest{
		UserId:    request.Author.Id,
		ChannelId: request.ChannelId,
		Month:     month,
		Day:       day,
	}

	_, err = s.command.Client.Birthday().Set(setRequest)

	if err != nil {
		return handler.Results{}, err
	}

	result := handler.Result{
		Content: fmt.Sprintf("Your birthday has been set to %s, %d. :birthday:", month, day),
	}

	return handler.Results{result}, nil
}

func (s *set) parseTime(content string) (*when.Result, error) {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Log("panic", r)
		}
	}()

	w := when.New(nil)
	w.Add(en.All...)

	r, err := w.Parse(content, time.Now())

	return r, err
}
