package entity

import (
	"time"

	"github.com/nanda03dev/go-ms-template/src/common"
	"github.com/nanda03dev/go-ms-template/src/core/domain/aggregate"
	"gorm.io/gorm"
)

const EventEntityName common.EntityName = "Event"

type Event struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	EntityId   string
	EntityName string
	Type       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Helper function: Converts an aggregate Event to an entity Event
func NewEvent(event *aggregate.Event) *Event {
	return &Event{
		ID:         event.ID,
		EntityId:   event.EntityId,
		EntityName: event.EntityName,
		Type:       event.Type,
		CreatedAt:  event.CreatedAt,
		UpdatedAt:  event.UpdatedAt,
	}
}

func (e *Event) GetEntityName() common.EntityName {
	return EventEntityName
}

// Helper function: Converts an entity Event to an aggregate Event
func (e *Event) ToDomain() *aggregate.Event {
	return &aggregate.Event{
		ID:         e.ID,
		EntityId:   e.EntityId,
		EntityName: e.EntityName,
		Type:       e.Type,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
	}
}
