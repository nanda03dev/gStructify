package repositories

import (
	"github.com/nanda03dev/go-ms-template/src/core/domain/aggregates"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/entity"
)

type TemplateEntityRepositoryImpl struct {
}

func NewTemplateEntityRepository(dbs *db.Databases) aggregates.TemplateEntityRepository {
	return &TemplateEntityRepositoryImpl{}
}

func (r *TemplateEntityRepositoryImpl) Create(templateEntity *aggregates.TemplateEntity) (*aggregates.TemplateEntity, error) {
	// Implement find by id logic here
	return templateEntity, nil
}

func (r *TemplateEntityRepositoryImpl) FindById(id string) (*aggregates.TemplateEntity, error) {
	var templateEntity entity.TemplateEntity
	// Implement find by id logic here
	return r.toDomain(&templateEntity), nil
}

func (r *TemplateEntityRepositoryImpl) Update(templateEntity *aggregates.TemplateEntity) (*aggregates.TemplateEntity, error) {
	// Implement update entity logic here
	return templateEntity, nil
}
func (r *TemplateEntityRepositoryImpl) Delete(id string) error {
	// Implement delete logic here
	return nil
}

// Convert entity.TemplateEntity to domain.TemplateEntity
func (r *TemplateEntityRepositoryImpl) toDomain(templateEntity *entity.TemplateEntity) *aggregates.TemplateEntity {
	return &aggregates.TemplateEntity{
		ID: templateEntity.ID,
	}
}

// Convert domain.TemplateEntity to entity.TemplateEntity for MongoDB storage
func (r *TemplateEntityRepositoryImpl) ToEntity(templateEntity *aggregates.TemplateEntity) *entity.TemplateEntity {
	return &entity.TemplateEntity{
		ID: templateEntity.ID,
	}
}
