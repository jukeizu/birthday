package commands

import (
	"github.com/go-kit/kit/log"

	"github.com/jukeizu/birthday/api"
)

type command struct {
	logger log.Logger
	Client api.Client
}

type Command interface {
	Set() Set
}

func NewCommand(logger log.Logger, client api.Client) Command {
	return &command{logger, client}
}
