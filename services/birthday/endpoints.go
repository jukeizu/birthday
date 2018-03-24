package birthday

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeSetBirthdayEndpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, receivedRequest interface{}) (interface{}, error) {
		request := receivedRequest.(SetBirthdayRequest)

		err := service.Set(request)

		return nil, err
	}
}
