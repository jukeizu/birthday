package birthday

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"

	"github.com/jukeizu/treediagram-services-scheduler/storage"
	"github.com/shawntoffel/services-core/transport"
)

type SetBirthdayRequest struct {
	User     string
	Birthday string
}

type SetBirthdayResponse struct {
	Job storage.Job `json:"job"`
}

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
