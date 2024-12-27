package entity

import (
	"time"

	"github.com/nanda03dev/go-ms-template/src/common"
	"github.com/nanda03dev/go-ms-template/src/helper"
	"gorm.io/gorm"
)

const TemplateEntityEntityName common.EntityName = "TemplateEntity"

type TemplateEntity struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	#@$Field$ $FieldType$#@
}

// Helper function: Converts an aggregate TemplateEntity to an entity TemplateEntity
func NewTemplateEntity(templateEntity *aggregates.TemplateEntity) *TemplateEntity {
	return &TemplateEntity{
		ID:        templateEntity.ID,
		CreatedAt: templateEntity.CreatedAt,
		UpdatedAt: templateEntity.UpdatedAt,
		#@$Field$: templateEntity.$Field$,#@
	}
}

func (e *TemplateEntity) GetEntityName() common.EntityName {
	return TemplateEntityEntityName
}

func (e *TemplateEntity) GetCreatedEvent() common.Event {
	return e.GetEvent(common.ENTITY_CREATED)
}
func (e *TemplateEntity) GetUpdatedEvent() common.Event {
	return e.GetEvent(common.ENTITY_CREATED)
}
func (e *TemplateEntity) GetDeletedEvent() common.Event {
	return e.GetEvent(common.ENTITY_CREATED)
}

func (e *TemplateEntity) GetEvent(operationType common.EventType) common.Event {
	return common.Event{
		ID:         helper.Generate16DigitUUID(),
		EntityId:   e.ID,
		EntityName: e.GetEntityName(),
		Type:       operationType,
		Config: common.EntityConfig{
			EventStore: true,
		},
	}
}

// Helper function: Converts an entity TemplateEntity to an aggregate TemplateEntity
func (templateEntity *TemplateEntity) ToDomain() *aggregates.TemplateEntity {
	return &aggregates.TemplateEntity{
		ID:        templateEntity.ID,
		CreatedAt: templateEntity.CreatedAt,
		UpdatedAt: templateEntity.UpdatedAt,
		#@$Field$: templateEntity.$Field$,#@
	}
}
