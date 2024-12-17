package repositories

import (
	"errors"
	"fmt"

	"github.com/nanda03dev/go-ms-template/src/core/domain/aggregates"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/entity"
	"gorm.io/gorm"
)

type TemplateEntityRepositoryImpl struct {
	db *gorm.DB
}

// NewTemplateEntityRepository initializes a new TemplateEntityRepositoryImpl
func NewTemplateEntityRepository(db *gorm.DB) aggregates.TemplateEntityRepository {
	return &TemplateEntityRepositoryImpl{
		db: db,
	}
}

// Create inserts a new TemplateEntity record into the database
func (r *TemplateEntityRepositoryImpl) Create(templateEntity *aggregates.TemplateEntity) (*aggregates.TemplateEntity, error) {
	entityTemplateEntity := r.toEntity(templateEntity)
	if err := r.db.Create(&entityTemplateEntity).Error; err != nil {
		return nil, fmt.Errorf("failed to create templateEntity: %w", err)
	}
	return r.toDomain(entityTemplateEntity), nil
}

// FindById retrieves an TemplateEntity record by its ID
func (r *TemplateEntityRepositoryImpl) FindById(id string) (*aggregates.TemplateEntity, error) {
	var entityTemplateEntity entity.TemplateEntity
	if err := r.db.First(&entityTemplateEntity, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("templateEntity not found")
		}
		return nil, fmt.Errorf("failed to find templateEntity by ID: %w", err)
	}
	return r.toDomain(&entityTemplateEntity), nil
}

// Update updates an existing TemplateEntity record in the database
func (r *TemplateEntityRepositoryImpl) Update(templateEntity *aggregates.TemplateEntity) (*aggregates.TemplateEntity, error) {
	entityTemplateEntity := r.toEntity(templateEntity)
	if err := r.db.Save(&entityTemplateEntity).Error; err != nil {
		return nil, fmt.Errorf("failed to update templateEntity: %w", err)
	}
	return r.toDomain(entityTemplateEntity), nil
}

// Delete removes an TemplateEntity record from the database by its ID
func (r *TemplateEntityRepositoryImpl) Delete(id string) error {
	if err := r.db.Delete(&entity.TemplateEntity{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete templateEntity: %w", err)
	}
	return nil
}

// Helper function: Converts an aggregate TemplateEntity to an entity TemplateEntity
func (r *TemplateEntityRepositoryImpl) toEntity(templateEntity *aggregates.TemplateEntity) *entity.TemplateEntity {
	return &entity.TemplateEntity{
		ID:        templateEntity.ID,
		CreatedAt: templateEntity.CreatedAt,
		UpdatedAt: templateEntity.UpdatedAt,
	}
}

// Helper function: Converts an entity TemplateEntity to an aggregate TemplateEntity
func (r *TemplateEntityRepositoryImpl) toDomain(templateEntity *entity.TemplateEntity) *aggregates.TemplateEntity {
	return &aggregates.TemplateEntity{
		ID:        templateEntity.ID,
		CreatedAt: templateEntity.CreatedAt,
		UpdatedAt: templateEntity.UpdatedAt,
	}
}
