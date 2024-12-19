package repositories

import (
	"fmt"

	"github.com/nanda03dev/go-ms-template/src/core/domain/aggregates"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/entity"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/worker_channels"
	"gorm.io/gorm"
)

// EventRepositoryImpl implements the EventRepository interface.
type EventRepositoryImpl struct {
	*BaseRepository[entity.Event] // Embeds BaseRepository for CRUD operations
}

// NewEventRepository initializes a new EventRepositoryImpl instance.
func NewEventRepository() aggregates.EventRepository {
	var databases = db.ConnectAll()
	return &EventRepositoryImpl{
		BaseRepository: NewBaseRepository[entity.Event](databases.DB.DB), // Initialize BaseRepository with the entity.Event type
	}
}

// Create inserts a new event.
func (r *EventRepositoryImpl) Create(event *aggregates.Event) (*aggregates.Event, error) {
	entityEvent := r.toEntity(event)
	createdEvent, err := r.BaseRepository.Create(entityEvent)

	return r.toDomain(createdEvent), err
}

// FindById retrieves a event by its ID.
func (r *EventRepositoryImpl) FindById(id string) (*aggregates.Event, error) {
	entityEvent, err := r.BaseRepository.FindById(id)
	return r.toDomain(entityEvent), err
}

// FindWithFilter retrieves a event by .
func (r *EventRepositoryImpl) FindWithFilter(filterQueryDTO common.FilterQueryDTO) ([]*aggregates.Event, error) {

	events, err := r.BaseRepository.FindWithFilter(filterQueryDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to find events by name: %w", err)
	}

	// Convert entity events to aggregate events and return
	var result []*aggregates.Event
	for _, event := range events {
		result = append(result, r.toDomain(event))
	}

	return result, nil
}

// Update modifies an existing event.
func (r *EventRepositoryImpl) Update(event *aggregates.Event) (*aggregates.Event, error) {
	entityEvent := r.toEntity(event)
	updatedEvent, err := r.BaseRepository.Update(entityEvent)

	return r.toDomain(updatedEvent), err
}

// Delete removes a event by its ID.
func (r *EventRepositoryImpl) Delete(id string) error {
	err := r.BaseRepository.Delete(id)
	return err
}

// Helper function: Converts an aggregate Event to an entity Event
func (r *EventRepositoryImpl) toEntity(event *aggregates.Event) *entity.Event {
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
func (r *EventRepositoryImpl) toDomain(event *entity.Event) *aggregates.Event {
	return &aggregates.Event{
		ID:         event.ID,
		EntityId:   event.EntityId,
		EntityName: event.EntityName,
		Type:       event.Type,
		CreatedAt:  event.CreatedAt,
		UpdatedAt:  event.UpdatedAt,
	}
}
