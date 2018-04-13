package birthday

import (
	"strconv"

	"github.com/jukeizu/clients/scheduler/scheduler"
	"github.com/rs/xid"
)

type Config struct {
	NotificationChannel string
	Hour                string
	Minute              string
}

type Service interface {
	Set(SetBirthdayRequest) error
	Remove(userId string, channelId string) error
}

type service struct {
	Storage          BirthdayStorage
	SchedulingClient schedulingclient.SchedulingClient
	Config           Config
}

func NewService(birthdayStorage BirthdayStorage, client schedulingclient.SchedulingClient, config Config) Service {
	return &service{birthdayStorage, client, config}
}

func (s *service) Set(request SetBirthdayRequest) error {
	err := s.Remove(request.UserId, request.ServerId)

	if err != nil {
		return err
	}

	jobRequest := schedulingclient.CreateJobRequest{}
	jobRequest.Type = "birthday"
	jobRequest.User = request.UserId
	jobRequest.Destination = s.Config.NotificationChannel
	jobRequest.Schedule.Month = request.Month
	jobRequest.Schedule.DayOfMonth = strconv.Itoa(request.Day)
	jobRequest.Schedule.Hour = s.Config.Hour
	jobRequest.Schedule.Minute = s.Config.Minute

	response, err := s.SchedulingClient.Create(jobRequest)

	if err != nil {
		return err
	}

	b := Birthday{}

	b.Id = xid.New().String()
	b.Month = request.Month
	b.Day = request.Day
	b.UserId = request.UserId
	b.ChannelId = request.ChannelId
	b.ServerId = request.ServerId
	b.JobId = response.Job.Id
	b.Enabled = true

	return s.Storage.Save(b)
}

func (s *service) Remove(userId string, serverId string) error {
	existing, _ := s.Storage.Birthday(userId, serverId)

	if existing.Id == "" {
		return nil
	}

	_, err := s.SchedulingClient.Disable(schedulingclient.DisableJobRequest{Id: existing.JobId})

	if err != nil {
		return err
	}

	return s.Storage.Disable(existing.Id)
}
