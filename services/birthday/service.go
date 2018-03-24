package birthday

import (
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
	jobRequest.Schedule.DayOfMonth = string(request.Day)

	job, err := s.SchedulingClient.Create(jobRequest)

	if err != nil {
		return err
	}

	b := Birthday{}

	b.Id = xid.New().String()
	b.Month = request.Month
	b.Day = request.Day
	b.UserId = request.UserId
	b.ChannelId = request.ChannelId
	b.JobId = job.Job.Id
	b.Enabled = true

	return s.Storage.Save(b)
}

func (s *service) Remove(userId string, channelId string) error {
	existing, err := s.Storage.Birthday(userId, channelId)

	if err != nil {
		return err
	}

	if existing.Id == "" {
		return nil
	}

	_, err = s.SchedulingClient.Disable(schedulingclient.DisableJobRequest{existing.JobId})

	if err != nil {
		return err
	}

	return s.Storage.Disable(existing.Id)
}
