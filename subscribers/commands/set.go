package commands

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-kit/kit/log"

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
	logger := log.With(c.logger, "component", "command.birthday.set")

	return &set{c, logger}
}

func (s *set) IsCommand(request handler.Request) (bool, error) {
	_, matches := parse(request.Content)

	return matches, nil
}

func (s *set) Handle(request handler.Request) (handler.Results, error) {
	input, match := parse(request.Content)

	if !match {
		return handler.Results{}, nil
	}

	t, err := time.Parse("January 2", input)

	if err != nil {
		errorMessage := handler.Result{
			Content: fmt.Sprintf("`%s` is not a valid input.", input),
		}

		return handler.Results{errorMessage}, nil
	}

	setRequest := api.SetBirthdayRequest{
		UserId:    request.Author.Id,
		ChannelId: request.ChannelId,
		Month:     t.Month().String(),
		Day:       t.Day(),
	}

	_, err = s.command.Client.Birthday().Set(setRequest)

	if err != nil {
		return handler.Results{}, err
	}

	result := handler.Result{
		Content: fmt.Sprintf("Ok. Your birthday is %s, %d.", t.Month().String(), t.Day()),
	}

	return handler.Results{result}, nil
}

func parse(content string) (string, bool) {
	re := regexp.MustCompile(`!setbirthday (\w+) (\d+)`)

	matches := re.FindAllStringSubmatch(content, 1)

	if len(matches) < 1 {
		return "", false
	}

	match := matches[0]

	if len(match) < 3 {
		return "", false
	}

	return match[1] + " " + match[2], true
}
