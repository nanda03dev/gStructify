package service

import (
	"github.com/nanda03dev/go-ms-template/src/common"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/repository"
	"github.com/nanda03dev/go-ms-template/src/core/interface/dto"
)

type TemplateEntityService interface {
	Create(createDTO dto.CreateTemplateEntityDTO) (*aggregate.TemplateEntity, error)
	GetById(id string) (*aggregate.TemplateEntity, error)
	FindWithFilter(filterQuery common.FilterQuery) ([]*aggregate.TemplateEntity, error)
	Update(id string, updateDTO dto.UpdateTemplateEntityDTO) (*aggregate.TemplateEntity, error)
	Delete(id string) error
}

type templateEntityService struct {
	templateEntityRepo repository.TemplateEntityRepository
}

func NewTemplateEntityService(templateEntityRepo repository.TemplateEntityRepository) TemplateEntityService {
	return &templateEntityService{
		templateEntityRepo: templateEntityRepo,
	}
}

func (s *templateEntityService) Create(createDTO dto.CreateTemplateEntityDTO) (*aggregate.TemplateEntity, error) {
	newData := aggregate.NewTemplateEntity(createDTO)
	return s.templateEntityRepo.Create(newData)
}

func (s *templateEntityService) GetById(id string) (*aggregate.TemplateEntity, error) {
	return s.templateEntityRepo.FindById(id)
}

func (s *templateEntityService) FindWithFilter(filterQuery common.FilterQuery) ([]*aggregate.TemplateEntity, error) {
	return s.templateEntityRepo.FindWithFilter(filterQuery)
}

func (s *templateEntityService) Update(id string, updateDTO dto.UpdateTemplateEntityDTO) (*aggregate.TemplateEntity, error) {
	updatedData := aggregate.UpdateTemplateEntity(id, updateDTO)
	return s.templateEntityRepo.Update(updatedData)
}

func (s *templateEntityService) Delete(id string) error {
	return s.templateEntityRepo.Delete(id)
}
