package treediagram

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jukeizu/birthday/api/protobuf-spec/birthdaypb"
	"github.com/jukeizu/contract"
	shellwords "github.com/mattn/go-shellwords"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/en"
	"github.com/rs/zerolog"
)

type Handler interface {
	SetBirthday(contract.Request) (*contract.Response, error)
	SayHappyBirthday(contract.Job) (*contract.Response, error)

	Start() error
	Stop() error
}

type handler struct {
	logger     zerolog.Logger
	client     birthdaypb.BirthdayServiceClient
	httpServer *http.Server
}

func NewHandler(logger zerolog.Logger, client birthdaypb.BirthdayServiceClient, addr string) Handler {
	logger = logger.With().Str("component", "intent.endpoint.birthday").Logger()

	httpServer := http.Server{
		Addr: addr,
	}

	return &handler{logger, client, &httpServer}
}

func (h *handler) SetBirthday(request contract.Request) (*contract.Response, error) {
	args, err := shellwords.Parse(request.Content)
	if err != nil {
		return nil, err
	}

	input := strings.Join(args[1:], " ")

	parsed, err := h.parseTime(input)
	if err != nil || parsed == nil {
		errorMessage := contract.Message{
			Content: fmt.Sprintf("`%s` is not a valid date format.", input),
		}

		return &contract.Response{Messages: []*contract.Message{&errorMessage}}, nil
	}

	month := parsed.Time.Month()
	day := parsed.Time.Day()

	setBirthdayRequest := birthdaypb.SetBirthdayRequest{
		ServerId: request.ServerId,
		UserId:   request.Author.Id,
		Month:    int32(month),
		Day:      int32(day),
	}

	_, err = h.client.SetBirthday(context.Background(), &setBirthdayRequest)
	if err != nil {
		return nil, err
	}

	message := contract.Message{
		Content: fmt.Sprintf("Your birthday has been set to %s %d. :birthday:", month.String(), day),
	}

	return &contract.Response{Messages: []*contract.Message{&message}}, nil
}

func (h *handler) SayHappyBirthday(job contract.Job) (*contract.Response, error) {
	now := time.Now().UTC()

	queryBirthdaysRequest := birthdaypb.QueryBirthdaysRequest{
		ServerId: job.Destination,
		Month:    int32(now.Month()),
		Day:      int32(now.Day()),
	}

	queryReply, err := h.client.QueryBirthdays(context.Background(), &queryBirthdaysRequest)
	if err != nil {
		return nil, err
	}

	response := &contract.Response{}

	if len(queryReply.Birthdays) < 1 {
		return response, nil
	}

	for _, birthday := range queryReply.Birthdays {
		message := contract.Message{
			Content: fmt.Sprintf("<@!%s> お誕生日おめでとう Happy Birthday! :birthday:", birthday.UserId),
		}

		response.Messages = append(response.Messages, &message)

		h.logger.Info().
			Str("userId", birthday.UserId).
			Str("serverId", birthday.ServerId).
			Int32("month", birthday.Month).
			Int32("day", birthday.Day).
			Msg("added birthday message for user")
	}

	return response, nil
}

func (h *handler) Start() error {
	h.logger.Info().Msg("starting")

	mux := http.NewServeMux()
	mux.HandleFunc("/setbirthday", h.makeLoggingHttpHandlerFunc("setbirthday", h.SetBirthday))
	mux.HandleFunc("/sayhappybirthday", h.makeLoggingJobHttpHandlerFunc("sayhappybirthday", h.SayHappyBirthday))

	h.httpServer.Handler = mux

	h.logger.Info().
		Str("transport", "http").
		Str("addr", h.httpServer.Addr).
		Msg("listening")

	return h.httpServer.ListenAndServe()
}

func (h *handler) Stop() error {
	h.logger.Info().Msg("stopping")

	return h.httpServer.Shutdown(context.Background())
}

func (h *handler) parseTime(content string) (*when.Result, error) {
	defer func() {
		if r := recover(); r != nil {
			h.logger.Error().Msgf("caught panic while parsing time: %v", r)
		}
	}()

	w := when.New(nil)
	w.Add(en.All...)

	r, err := w.Parse(content, time.Now())
	if err != nil {
		h.logger.Error().Err(err).Caller().Msg("couldn't parse time")
		return nil, err
	}

	return r, err
}

func (h *handler) makeLoggingHttpHandlerFunc(name string, f func(contract.Request) (*contract.Response, error)) http.HandlerFunc {
	contractHandlerFunc := contract.MakeRequestHttpHandlerFunc(f)

	return func(w http.ResponseWriter, r *http.Request) {
		defer func(begin time.Time) {
			h.logger.Info().
				Str("intent", name).
				Str("took", time.Since(begin).String()).
				Msg("called")
		}(time.Now())

		contractHandlerFunc.ServeHTTP(w, r)
	}
}

func (h *handler) makeLoggingJobHttpHandlerFunc(name string, f func(contract.Job) (*contract.Response, error)) http.HandlerFunc {
	contractHandlerFunc := contract.MakeJobHttpHandlerFunc(f)

	return func(w http.ResponseWriter, r *http.Request) {
		defer func(begin time.Time) {
			h.logger.Info().
				Str("job", name).
				Str("took", time.Since(begin).String()).
				Msg("called")
		}(time.Now())

		contractHandlerFunc.ServeHTTP(w, r)
	}
}
