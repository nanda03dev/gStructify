package services

import (
	"github.com/nanda03dev/go-ms-template/src/domain/aggregates"
	"github.com/nanda03dev/go-ms-template/src/interface/dto"
)

type TemplateEntityService interface {
	Create(createTemplateEntityDTO dto.CreateTemplateEntityDTO) (*aggregates.TemplateEntity, error)
	GetById(id string) (*aggregates.TemplateEntity, error)
	Update(id string, updateTemplateEntityDTO dto.UpdateTemplateEntityDTO) (*aggregates.TemplateEntity, error)
	Delete(id string) error
}

type templateEntityService struct {
	templateEntityRepo aggregates.TemplateEntityRepository
}

func NewTemplateEntityService(templateEntityRepo aggregates.TemplateEntityRepository) TemplateEntityService {
	return &templateEntityService{
		templateEntityRepo: templateEntityRepo,
	}
}

func (s *templateEntityService) Create(createTemplateEntityDTO dto.CreateTemplateEntityDTO) (*aggregates.TemplateEntity, error) {
	newData := aggregates.NewTemplateEntity(createTemplateEntityDTO)
	return s.templateEntityRepo.Create(newData)
}

func (s *templateEntityService) GetById(id string) (*aggregates.TemplateEntity, error) {
	return s.templateEntityRepo.FindById(id)
}

func (s *templateEntityService) Update(id string, updateTemplateEntityDTO dto.UpdateTemplateEntityDTO) (*aggregates.TemplateEntity, error) {
	updatedData := aggregates.UpdateTemplateEntity(id, updateTemplateEntityDTO)
	return s.templateEntityRepo.Update(updatedData)
}

func (s *templateEntityService) Delete(id string) error {
	return s.templateEntityRepo.Delete(id)
}
