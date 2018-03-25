package api

import restclient "github.com/shawntoffel/go-restclient"

type Birthday struct {
	client *client
}

type SetBirthdayRequest struct {
	UserId    string
	ChannelId string
	Month     string
	Day       int
}

type SetBirthdayResponse struct{}

func (c *client) Birthday() *Birthday {
	return &Birthday{c}
}

func (b *Birthday) Set(request SetBirthdayRequest) (SetBirthdayResponse, error) {
	response := SetBirthdayResponse{}

	err := restclient.Post(b.client.ClientConfig.BaseUrl+"/birthday", request, &response)

	return response, err
}
