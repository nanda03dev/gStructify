package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const TemplateEntityEntityName common.EntityName = "TemplateEntity"

type TemplateEntity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"` // MongoDB ObjectID
	#@$Field$ $FieldType$#@      
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

// Helper function: Converts an aggregate TemplateEntity to a MongoDB entity
func NewTemplateEntity(templateEntity *aggregate.TemplateEntity) *TemplateEntity {
	objectID, _ := primitive.ObjectIDFromHex(templateEntity.ID)
	return &TemplateEntity{
		ID:        objectID,
		#@$Field$: templateEntity.$Field$,#@
		CreatedAt: templateEntity.CreatedAt,
		UpdatedAt: templateEntity.UpdatedAt,
	}
}

func (e *TemplateEntity) GetEntityName() common.EntityName {
	return TemplateEntityEntityName
}

func (e *TemplateEntity) GetCreatedEvent() common.Event {
	return e.GetEvent(common.ENTITY_CREATED)
}
func (e *TemplateEntity) GetUpdatedEvent() common.Event {
	return e.GetEvent(common.ENTITY_UPDATED)
}
func (e *TemplateEntity) GetDeletedEvent() common.Event {
	return e.GetEvent(common.ENTITY_DELETED)
}

func (e *TemplateEntity) GetEvent(operationType common.EventType) common.Event {
	return common.Event{
		ID:         helper.Generate16DigitUUID(),
		EntityId:   e.ID.Hex(), // Convert ObjectID to string
		EntityName: e.GetEntityName(),
		Type:       operationType,
		Config: common.EntityConfig{
			EventStore: true,
		},
	}
}

// Helper function: Converts a MongoDB entity TemplateEntity to an aggregate TemplateEntity
func (templateEntity *TemplateEntity) ToDomain() *aggregate.TemplateEntity {
	return &aggregate.TemplateEntity{
		ID:        templateEntity.ID.Hex(), // Convert ObjectID to string
		#@$Field$: templateEntity.$Field$,#@
		CreatedAt: templateEntity.CreatedAt,
		UpdatedAt: templateEntity.UpdatedAt,
	}
}
