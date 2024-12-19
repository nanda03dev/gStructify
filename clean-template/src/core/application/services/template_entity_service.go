package services

import (
	"github.com/nanda03dev/go-ms-template/src/core/domain/aggregates"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/repositories"
	"github.com/nanda03dev/go-ms-template/src/core/interface/dto"
)

type TemplateEntityService interface {
	Create(createTemplateEntityDTO dto.CreateTemplateEntityDTO) (*aggregates.TemplateEntity, error)
	GetById(id string) (*aggregates.TemplateEntity, error)
	FindWithFilter(filterQueryDTO common.FilterQueryDTO) ([]*aggregates.TemplateEntity, error)
	Update(id string, updateTemplateEntityDTO dto.UpdateTemplateEntityDTO) (*aggregates.TemplateEntity, error)
	Delete(id string) error
}

type templateEntityService struct {
	templateEntityRepo aggregates.TemplateEntityRepository
}

func NewTemplateEntityService() TemplateEntityService {
	return &templateEntityService{
		templateEntityRepo: repositories.NewTemplateEntityRepository(),
	}
}

func (s *templateEntityService) Create(createTemplateEntityDTO dto.CreateTemplateEntityDTO) (*aggregates.TemplateEntity, error) {
	newData := aggregates.NewTemplateEntity(createTemplateEntityDTO)
	return s.templateEntityRepo.Create(newData)
}

func (s *templateEntityService) GetById(id string) (*aggregates.TemplateEntity, error) {
	return s.templateEntityRepo.FindById(id)
}

func (s *templateEntityService) FindWithFilter(filterQueryDTO common.FilterQueryDTO) ([]*aggregates.TemplateEntity, error) {
	return s.templateEntityRepo.FindWithFilter(filterQueryDTO)
}

func (s *templateEntityService) Update(id string, updateTemplateEntityDTO dto.UpdateTemplateEntityDTO) (*aggregates.TemplateEntity, error) {
	updatedData := aggregates.UpdateTemplateEntity(id, updateTemplateEntityDTO)
	return s.templateEntityRepo.Update(updatedData)
}

func (s *templateEntityService) Delete(id string) error {
	return s.templateEntityRepo.Delete(id)
}
