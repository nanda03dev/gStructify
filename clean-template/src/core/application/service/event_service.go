package service

import (
	"github.com/nanda03dev/go-ms-template/src/common"
	"github.com/nanda03dev/go-ms-template/src/core/domain/aggregate"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/repository"
)

type EventService interface {
	Create(createDTO common.Event) (*aggregate.Event, error)
	GetById(id string) (*aggregate.Event, error)
	FindWithFilter(filterQuery common.FilterQuery) ([]*aggregate.Event, error)
	Update(id string, updateDTO common.Event) (*aggregate.Event, error)
	Delete(id string) error
}

type eventService struct {
	eventRepo repository.EventRepository
}

func NewEventService(eventRepository repository.EventRepository) EventService {
	return &eventService{
		eventRepo: eventRepository,
	}
}

func (s *eventService) Create(createDTO common.Event) (*aggregate.Event, error) {
	newData := aggregate.NewEvent(createDTO)
	return s.eventRepo.Create(newData)
}

func (s *eventService) GetById(id string) (*aggregate.Event, error) {
	return s.eventRepo.FindById(id)
}

func (s *eventService) FindWithFilter(filterQuery common.FilterQuery) ([]*aggregate.Event, error) {
	return s.eventRepo.FindWithFilter(filterQuery)
}

func (s *eventService) Update(id string, updateDTO common.Event) (*aggregate.Event, error) {
	updatedData := aggregate.UpdateEvent(id, updateDTO)
	return s.eventRepo.Update(updatedData)
}

func (s *eventService) Delete(id string) error {
	return s.eventRepo.Delete(id)
}
