package services

import (
	"github.com/nanda03dev/go-ms-template/src/common"
	"github.com/nanda03dev/go-ms-template/src/core/domain/aggregates"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/repositories"
	"github.com/nanda03dev/go-ms-template/src/core/interface/dto"
)

type EventService interface {
	Create(createEventDTO common.Event) (*aggregates.Event, error)
	GetById(id string) (*aggregates.Event, error)
	FindWithFilter(filterQueryDTO common.FilterQueryDTO) ([]*aggregates.Event, error)
	Update(id string, updateEventDTO common.Event) (*aggregates.Event, error)
	Delete(id string) error
}

type eventService struct {
	eventRepo aggregates.EventRepository
}

func NewEventService() EventService {
	return &eventService{
		eventRepo: repositories.NewEventRepository(),
	}
}

func (s *eventService) Create(createEventDTO common.Event) (*aggregates.Event, error) {
	newData := aggregates.NewEvent(createEventDTO)
	return s.eventRepo.Create(newData)
}

func (s *eventService) GetById(id string) (*aggregates.Event, error) {
	return s.eventRepo.FindById(id)
}

func (s *eventService) FindWithFilter(filterQueryDTO common.FilterQueryDTO) ([]*aggregates.Event, error) {
	return s.eventRepo.FindWithFilter(filterQueryDTO)
}

func (s *eventService) Update(id string, updateEventDTO common.Event) (*aggregates.Event, error) {
	updatedData := aggregates.UpdateEvent(id, updateEventDTO)
	return s.eventRepo.Update(updatedData)
}

func (s *eventService) Delete(id string) error {
	return s.eventRepo.Delete(id)
}
