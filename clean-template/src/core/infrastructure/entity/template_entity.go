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
