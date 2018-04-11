package birthday

import (
	"strconv"

	"github.com/jukeizu/clients/scheduler/scheduler"
	"github.com/rs/xid"
)

type Service interface {
	Set(SetBirthdayRequest) error
	Remove(userId string, channelId string) error
}

type service struct {
	Storage          BirthdayStorage
	SchedulingClient schedulingclient.SchedulingClient
}

func NewService(birthdayStorage BirthdayStorage, client schedulingclient.SchedulingClient) Service {
	return &service{birthdayStorage, client}
}

func (s *service) Set(request SetBirthdayRequest) error {
	err := s.Remove(request.UserId, request.ChannelId)

	if err != nil {
		return err
	}

	jobRequest := schedulingclient.CreateJobRequest{}
	jobRequest.Type = "birthday"
	jobRequest.User = request.UserId
	jobRequest.Destination = request.ChannelId
	jobRequest.Schedule.Month = request.Month
	jobRequest.Schedule.DayOfMonth = strconv.Itoa(request.Day)
	jobRequest.Schedule.Hour = "11"
	jobRequest.Schedule.Minute = "0"

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
	b.JobId = response.Job.Id
	b.Enabled = true

	return s.Storage.Save(b)
}

func (s *service) Remove(userId string, channelId string) error {
	existing, _ := s.Storage.Birthday(userId, channelId)

	if existing.Id == "" {
		return nil
	}

	_, err := s.SchedulingClient.Disable(schedulingclient.DisableJobRequest{Id: existing.JobId})

	if err != nil {
		return err
	}

	return s.Storage.Disable(existing.Id)
}
