package aggregates

import (
	"time"

	"github.com/nanda03dev/go-ms-template/src/core/interface/dto"
	"github.com/nanda03dev/go-ms-template/src/helper"
)

type TemplateEntity struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTemplateEntity(createTemplateEntityDTO dto.CreateTemplateEntityDTO) *TemplateEntity {
	return &TemplateEntity{
		ID: helper.Generate16DigitUUID(), // Generate unique ID (UUID or similar)
		// add other fields
	}
}
func UpdateTemplateEntity(id string, updateTemplateEntityDTO dto.UpdateTemplateEntityDTO) *TemplateEntity {
	return &TemplateEntity{
		ID: updateTemplateEntityDTO.ID,
		// add other fields
	}
}

type TemplateEntityRepository interface {
	Create(templateEntity *TemplateEntity) (*TemplateEntity, error)
	FindById(id string) (*TemplateEntity, error)
	Update(templateEntity *TemplateEntity) (*TemplateEntity, error)
	Delete(id string) error
}
