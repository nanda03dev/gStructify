package repositories

import (
	"fmt"

	"github.com/nanda03dev/go-ms-template/src/core/domain/aggregates"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/entity"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/worker_channels"
	"gorm.io/gorm"
)

// TemplateEntityRepositoryImpl implements the TemplateEntityRepository interface.
type TemplateEntityRepositoryImpl struct {
	*BaseRepository[entity.TemplateEntity] // Embeds BaseRepository for CRUD operations
}

// NewTemplateEntityRepository initializes a new TemplateEntityRepositoryImpl instance.
func NewTemplateEntityRepository() aggregates.TemplateEntityRepository {
	var databases = db.ConnectAll()
	return &TemplateEntityRepositoryImpl{
		BaseRepository: NewBaseRepository[entity.TemplateEntity](databases.DB.DB), // Initialize BaseRepository with the entity.TemplateEntity type
	}
}

// Create inserts a new templateEntity.
func (r *TemplateEntityRepositoryImpl) Create(templateEntity *aggregates.TemplateEntity) (*aggregates.TemplateEntity, error) {
	entityTemplateEntity := r.toEntity(templateEntity)
	createdTemplateEntity, err := r.BaseRepository.Create(entityTemplateEntity)

	eventChannel := worker_channels.GetCRUDEventChannel()
	eventChannel <- entityTemplateEntity.GetCreatedEvent()

	return r.toDomain(createdTemplateEntity), err
}

// FindById retrieves a templateEntity by its ID.
func (r *TemplateEntityRepositoryImpl) FindById(id string) (*aggregates.TemplateEntity, error) {
	entityTemplateEntity, err := r.BaseRepository.FindById(id)
	return r.toDomain(entityTemplateEntity), err
}

// FindWithFilter retrieves a templateEntity by .
func (r *TemplateEntityRepositoryImpl) FindWithFilter(filterQueryDTO common.FilterQueryDTO) ([]*aggregates.TemplateEntity, error) {

	templateEntitys, err := r.BaseRepository.FindWithFilter(filterQueryDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to find templateEntitys by name: %w", err)
	}

	// Convert entity templateEntitys to aggregate templateEntitys and return
	var result []*aggregates.TemplateEntity
	for _, templateEntity := range templateEntitys {
		result = append(result, r.toDomain(templateEntity))
	}

	return result, nil
}

// Update modifies an existing templateEntity.
func (r *TemplateEntityRepositoryImpl) Update(templateEntity *aggregates.TemplateEntity) (*aggregates.TemplateEntity, error) {
	entityTemplateEntity := r.toEntity(templateEntity)
	updatedTemplateEntity, err := r.BaseRepository.Update(entityTemplateEntity)

	eventChannel := worker_channels.GetCRUDEventChannel()
	eventChannel <- updatedTemplateEntity.GetUpdatedEvent()

	return r.toDomain(updatedTemplateEntity), err
}

// Delete removes a templateEntity by its ID.
func (r *TemplateEntityRepositoryImpl) Delete(id string) error {
	err := r.BaseRepository.Delete(id)
	if err != nil {
		entity := entity.TemplateEntity{ID: id}
		eventChannel := worker_channels.GetCRUDEventChannel()
		eventChannel <- entity.GetUpdatedEvent()
	}
	return err
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
