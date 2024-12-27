package services

import (
	"github.com/nanda03dev/go-ms-template/src/common"
	"github.com/nanda03dev/go-ms-template/src/core/domain/aggregates"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/repositories"
	"github.com/nanda03dev/go-ms-template/src/core/interface/dto"
)

type TemplateEntityService interface {
	Create(createDTO dto.CreateTemplateEntityDTO) (*aggregates.TemplateEntity, error)
	GetById(id string) (*aggregates.TemplateEntity, error)
	FindWithFilter(filterQuery common.FilterQuery) ([]*aggregates.TemplateEntity, error)
	Update(id string, updateDTO dto.UpdateTemplateEntityDTO) (*aggregates.TemplateEntity, error)
	Delete(id string) error
}

type templateEntityService struct {
	templateEntityRepo repositories.TemplateEntityRepository
}

func NewTemplateEntityService(templateEntityRepo repositories.TemplateEntityRepository) TemplateEntityService {
	return &templateEntityService{
		templateEntityRepo: templateEntityRepo,
	}
}

func (s *templateEntityService) Create(createDTO dto.CreateTemplateEntityDTO) (*aggregates.TemplateEntity, error) {
	newData := aggregates.NewTemplateEntity(createDTO)
	return s.templateEntityRepo.Create(newData)
}

func (s *templateEntityService) GetById(id string) (*aggregates.TemplateEntity, error) {
	return s.templateEntityRepo.FindById(id)
}

func (s *templateEntityService) FindWithFilter(filterQuery common.FilterQuery) ([]*aggregates.TemplateEntity, error) {
	return s.templateEntityRepo.FindWithFilter(filterQuery)
}

func (s *templateEntityService) Update(id string, updateDTO dto.UpdateTemplateEntityDTO) (*aggregates.TemplateEntity, error) {
	updatedData := aggregates.UpdateTemplateEntity(id, updateDTO)
	return s.templateEntityRepo.Update(updatedData)
}

func (s *templateEntityService) Delete(id string) error {
	return s.templateEntityRepo.Delete(id)
}
