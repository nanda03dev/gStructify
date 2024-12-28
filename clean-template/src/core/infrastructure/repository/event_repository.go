package repository

import (
	"fmt"

	"github.com/nanda03dev/go-ms-template/src/common"
	"github.com/nanda03dev/go-ms-template/src/core/domain/aggregate"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/entity"
)

type EventRepository interface {
	Create(event *aggregate.Event) (*aggregate.Event, error)
	FindById(id string) (*aggregate.Event, error)
	FindWithFilter(filterQuery common.FilterQuery) ([]*aggregate.Event, error)
	Update(event *aggregate.Event) (*aggregate.Event, error)
	Delete(id string) error
}

// eventRepository implements the EventRepository interface.
type eventRepository struct {
	*BaseRepository[entity.Event] // Embeds BaseRepository for CRUD operations
}

// NewEventRepository initializes a new eventRepository instance.
func NewEventRepository(databases *db.Databases) EventRepository {
	return &eventRepository{
		BaseRepository: NewBaseRepository[entity.Event](databases.SqlDB.DB), // Initialize BaseRepository with the entity.Event type
	}
}

// Create inserts a new event.
func (r *eventRepository) Create(event *aggregate.Event) (*aggregate.Event, error) {
	entityEvent := r.toEntity(event)
	createdEvent, err := r.BaseRepository.Create(entityEvent)

	return r.toDomain(createdEvent), err
}

// FindById retrieves a event by its ID.
func (r *eventRepository) FindById(id string) (*aggregate.Event, error) {
	entityEvent, err := r.BaseRepository.FindById(id)
	return r.toDomain(entityEvent), err
}

// FindWithFilter retrieves a event by .
func (r *eventRepository) FindWithFilter(filterQuery common.FilterQuery) ([]*aggregate.Event, error) {

	events, err := r.BaseRepository.FindWithFilter(filterQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to find events by name: %w", err)
	}

	// Convert entity events to aggregate events and return
	var result []*aggregate.Event
	for _, event := range events {
		result = append(result, r.toDomain(event))
	}

	return result, nil
}

// Update modifies an existing event.
func (r *eventRepository) Update(event *aggregate.Event) (*aggregate.Event, error) {
	entityEvent := r.toEntity(event)
	updatedEvent, err := r.BaseRepository.Update(entityEvent)

	return r.toDomain(updatedEvent), err
}

// Delete removes a event by its ID.
func (r *eventRepository) Delete(id string) error {
	err := r.BaseRepository.Delete(id)
	return err
}

// Helper function: Converts an aggregate Event to an entity Event
func (r *eventRepository) toEntity(event *aggregate.Event) *entity.Event {
	return &entity.Event{
		ID:         event.ID,
		EntityId:   event.EntityId,
		EntityName: event.EntityName,
		Type:       event.Type,
		CreatedAt:  event.CreatedAt,
		UpdatedAt:  event.UpdatedAt,
	}
}

// Helper function: Converts an entity Event to an aggregate Event
func (r *eventRepository) toDomain(event *entity.Event) *aggregate.Event {
	return &aggregate.Event{
		ID:         event.ID,
		EntityId:   event.EntityId,
		EntityName: event.EntityName,
		Type:       event.Type,
		CreatedAt:  event.CreatedAt,
		UpdatedAt:  event.UpdatedAt,
	}
}
