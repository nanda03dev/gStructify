package entity

import (
	"github.com/nanda03dev/go-ms-template/src/common"
	"gorm.io/gorm"
)

const EventEntityName common.ENTITY_NAME = "Event"

type Event struct {
	gorm.Model
	ID            string `gorm:"primaryKey"`
	EntityId      string
	EntityName    string
	OperationType string
}

func (e *Event) GetEntityName() common.ENTITY_NAME {
	return EventEntityName
}
