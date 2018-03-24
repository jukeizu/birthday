package birthday

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) Set(request SetBirthdayRequest) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Set",
			"request", request,
			"took", time.Since(begin),
		)

	}(time.Now())

	err = s.Service.Set(request)

	return
}
