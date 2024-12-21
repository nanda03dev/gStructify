package entity

import (
	"time"

	"github.com/nanda03dev/go-ms-template/src/common"
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

func (e *Event) GetEntityName() common.EntityName {
	return EventEntityName
}
