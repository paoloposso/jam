package musicalevents

type Service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return Service{repository: repository}
}

func (service *Service) CreateEvent(musicalEvent MusicalEvent) (musicalEventId string, err error) {
	return service.repository.Create(musicalEvent)
}
