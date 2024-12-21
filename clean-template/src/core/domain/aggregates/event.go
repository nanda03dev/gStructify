package aggregates

import (
	"time"

	"github.com/nanda03dev/go-ms-template/src/common"
	"github.com/nanda03dev/go-ms-template/src/helper"
)

type Event struct {
	ID         string
	EntityId   string
	EntityName string
	Type       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewEvent(createEventDTO common.Event) *Event {
	return &Event{
		ID:         helper.Generate16DigitUUID(), // Generate unique ID (UUID or similar)
		EntityId:   createEventDTO.EntityId,
		EntityName: string(createEventDTO.EntityName),
		Type:       string(createEventDTO.Type),
	}
}

func UpdateEvent(id string, updateEventDTO common.Event) *Event {
	return &Event{
		ID:         id,
		EntityId:   updateEventDTO.EntityId,
		EntityName: string(updateEventDTO.EntityName),
		Type:       string(updateEventDTO.Type),
	}
}

type EventRepository interface {
	Create(event *Event) (*Event, error)
	FindById(id string) (*Event, error)
	FindWithFilter(filterQuery common.FilterQuery) ([]*Event, error)
	Update(event *Event) (*Event, error)
	Delete(id string) error
}
