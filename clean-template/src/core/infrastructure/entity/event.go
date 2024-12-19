package entity

import (
	"time"

	"github.com/nanda03dev/go-ms-template/src/common"
	"gorm.io/gorm"
)

const EventEntityName common.ENTITY_NAME = "Event"

type Event struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	EntityId   string
	EntityName string
	Type       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (e *Event) GetEntityName() common.ENTITY_NAME {
	return EventEntityName
}
