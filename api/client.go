package api

import (
	"github.com/jukeizu/treediagram-client-base"
)

type Client interface {
	Birthday() *Birthday
}

type client struct {
	ClientConfig treediagramclient.ClientConfig
}

func NewClient(config treediagramclient.ClientConfig) Client {
	return &client{config}
}