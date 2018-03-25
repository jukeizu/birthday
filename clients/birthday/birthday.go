package birthday

import (
	"github.com/jukeizu/treediagram-client-base"
	restclient "github.com/shawntoffel/go-restclient"
)

type SetBirthdayRequest struct {
	UserId    string
	ChannelId string
	Month     string
	Day       int
}

type SetBirthdayResponse struct{}

type BirthdayClient interface {
	Set(SetBirthdayRequest) (SetBirthdayResponse, error)
}

type client struct {
	ClientConfig treediagramclient.ClientConfig
}

func NewClient(config treediagramclient.ClientConfig) BirthdayClient {
	return &client{config}
}

func (c *client) Set(request SetBirthdayRequest) (SetBirthdayResponse, error) {
	response := SetBirthdayResponse{}

	err := restclient.Post(c.ClientConfig.BaseUrl+"/birthday", request, &response)

	return response, err
}
