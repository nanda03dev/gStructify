package aggregate

import (
	"time"

	"github.com/nanda03dev/go-ms-template/src/core/interface/dto"
	"github.com/nanda03dev/go-ms-template/src/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TemplateEntity struct {
	ID        string
	#@$Field$ $FieldType$#@
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTemplateEntity(createDTO dto.CreateTemplateEntityDTO) *TemplateEntity {
	return &TemplateEntity{
		ID: primitive.NewObjectID().Hex(),
		#@$Field$: createDTO.$Field$,#@
	}
}

func UpdateTemplateEntity(id string, updateDTO dto.UpdateTemplateEntityDTO) *TemplateEntity {
	return &TemplateEntity{
		ID: id,
		#@$Field$: updateDTO.$Field$,#@
	}
}
