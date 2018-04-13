package birthday

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"

	"github.com/shawntoffel/services-core/transport"
)

type SetBirthdayRequest struct {
	UserId    string `json:"userId"`
	ChannelId string `json:"channelId"`
	ServerId  string `json:"serverId"`
	Month     string `json:"month"`
	Day       int    `json:"day"`
}

type SetBirthdayResponse struct{}

func DecodeSetBirthdayRequest(_ context.Context, r *http.Request) (interface{}, error) {
	request := SetBirthdayRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)

	return request, err
}

func MakeHandler(s Service, logger log.Logger) http.Handler {
	setBirthdayHandler := transport.NewDefaultServer(
		logger,
		MakeSetBirthdayEndpoint(s),
		DecodeSetBirthdayRequest,
	)

	router := mux.NewRouter()

	router.Handle("/birthday", setBirthdayHandler).Methods("POST")

	return router
}
