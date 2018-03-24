package birthday

type Service interface {
	Set(SetBirthdayRequest) error
}

type service struct {
	Storage BirthdayStorage
}

func NewService(birthdayStorage BirthdayStorage) Service {
	return &service{birthdayStorage}
}

func (s *service) Set(request SetBirthdayRequest) error {
	return nil
}
