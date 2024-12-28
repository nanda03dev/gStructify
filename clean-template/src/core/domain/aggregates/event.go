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

func NewEvent(createDTO common.Event) *Event {
	return &Event{
		ID:         helper.Generate16DigitUUID(), // Generate unique ID (UUID or similar)
		EntityId:   createDTO.EntityId,
		EntityName: string(createDTO.EntityName),
		Type:       string(createDTO.Type),
	}
}

func UpdateEvent(id string, updateDTO common.Event) *Event {
	return &Event{
		ID:         id,
		EntityId:   updateDTO.EntityId,
		EntityName: string(updateDTO.EntityName),
		Type:       string(updateDTO.Type),
	}
}

type EventRepository interface {
	Create(event *Event) (*Event, error)
	FindById(id string) (*Event, error)
	FindWithFilter(filterQuery common.FilterQuery) ([]*Event, error)
	Update(event *Event) (*Event, error)
	Delete(id string) error
}
