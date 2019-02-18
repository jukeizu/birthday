package main

import (
	"context"
	"errors"
	"net"

	"github.com/jukeizu/birthday/api/protobuf-spec/birthdaypb"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Server struct {
	logger     zerolog.Logger
	repository Repository
	grpcServer *grpc.Server
}

func NewServer(logger zerolog.Logger, repository Repository, grpcServer *grpc.Server) Server {
	return Server{logger, repository, grpcServer}
}

func (s Server) SetBirthday(ctx context.Context, req *birthdaypb.SetBirthdayRequest) (*birthdaypb.SetBirthdayReply, error) {
	birthday, err := s.repository.SetBirthday(req)
	if err != nil {
		return nil, errors.New("couldn't set birthday: " + err.Error())
	}

	return &birthdaypb.SetBirthdayReply{Birthday: birthday}, nil
}

func (s Server) Query(ctx context.Context, req *birthdaypb.QueryBirthdaysRequest) (*birthdaypb.QueryBirthdaysReply, error) {
	birthdays, err := s.repository.Query(req)
	if err != nil {
		return nil, err
	}

	return &birthdaypb.QueryBirthdaysReply{Birthdays: birthdays}, nil
}

func (s Server) Start(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s.logger.Info().
		Str("transport", "grpc").
		Str("addr", addr).
		Msg("listening")

	return s.grpcServer.Serve(listener)
}

func (s Server) Stop() {
	if s.grpcServer == nil {
		return
	}

	s.logger.Info().
		Str("transport", "grpc").
		Msg("stopping")

	s.grpcServer.GracefulStop()
}
