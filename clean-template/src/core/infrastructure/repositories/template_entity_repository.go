package repositories

import (
	"fmt"

	"github.com/nanda03dev/go-ms-template/src/common"
	"github.com/nanda03dev/go-ms-template/src/core/domain/aggregates"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/entity"
	"github.com/nanda03dev/go-ms-template/src/core/infrastructure/worker_channels"
)

type TemplateEntityRepository interface {
	Create(templateEntity *aggregates.TemplateEntity) (*aggregates.TemplateEntity, error)
	BulkCreate(templateEntity []*aggregates.TemplateEntity) ([]*aggregates.TemplateEntity, error)
	FindById(id string) (*aggregates.TemplateEntity, error)
	FindWithFilter(filterQuery common.FilterQuery) ([]*aggregates.TemplateEntity, error)
	Update(templateEntity *aggregates.TemplateEntity) (*aggregates.TemplateEntity, error)
	Delete(id string) error
}

// templateEntityRepository implements the TemplateEntityRepository interface.
type templateEntityRepository struct {
	*BaseRepository[entity.TemplateEntity] // Embeds BaseRepository for CRUD operations
}

// NewTemplateEntityRepository initializes a new templateEntityRepository instance.
func NewTemplateEntityRepository(databases *db.Databases) TemplateEntityRepository {
	return &templateEntityRepository{
		BaseRepository: NewBaseRepository[entity.TemplateEntity](databases.Postgres.DB), // Initialize BaseRepository with the entity.TemplateEntity type
	}
}

// Create inserts a new templateEntity.
func (r *templateEntityRepository) Create(templateEntity *aggregates.TemplateEntity) (*aggregates.TemplateEntity, error) {
	entityTemplateEntity := entity.NewTemplateEntity(templateEntity)
	createdTemplateEntity, err := r.BaseRepository.Create(entityTemplateEntity)

	if err != nil {
		return nil, err
	}

	eventChannel := worker_channels.GetCRUDEventChannel()
	eventChannel <- entityTemplateEntity.GetCreatedEvent()

	return createdTemplateEntity.ToDomain(), nil
}

// Bulk inserts a new templateEntity.
func (r *templateEntityRepository) BulkCreate(aggregateList []*aggregates.TemplateEntity) ([]*aggregates.TemplateEntity, error) {

	var entityList = make([]*entity.TemplateEntity, 0, len(aggregateList))

	for _, each := range aggregateList {
		entityList = append(entityList, entity.NewTemplateEntity(each))
	}

	createdList, err := r.BaseRepository.BulkCreate(entityList)

	if err != nil {
		return nil, err
	}

	for _, each := range createdList {
		eventChannel := worker_channels.GetCRUDEventChannel()
		eventChannel <- each.GetCreatedEvent()

	}
	return aggregateList, nil
}

// FindById retrieves a templateEntity by its ID.
func (r *templateEntityRepository) FindById(id string) (*aggregates.TemplateEntity, error) {
	entityTemplateEntity, err := r.BaseRepository.FindById(id)
	return entityTemplateEntity.ToDomain(), err
}

// FindWithFilter retrieves a templateEntity by .
func (r *templateEntityRepository) FindWithFilter(filterQuery common.FilterQuery) ([]*aggregates.TemplateEntity, error) {

	templateEntitys, err := r.BaseRepository.FindWithFilter(filterQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to find templateEntitys by name: %w", err)
	}

	// Convert entity templateEntitys to aggregate templateEntitys and return
	var result []*aggregates.TemplateEntity
	for _, templateEntity := range templateEntitys {
		result = append(result, templateEntity.ToDomain())
	}

	return result, nil
}

// Update modifies an existing templateEntity.
func (r *templateEntityRepository) Update(templateEntity *aggregates.TemplateEntity) (*aggregates.TemplateEntity, error) {
	entityTemplateEntity := entity.NewTemplateEntity(templateEntity)
	updatedTemplateEntity, err := r.BaseRepository.Update(entityTemplateEntity)

	if err != nil {
		return nil, err
	}

	eventChannel := worker_channels.GetCRUDEventChannel()
	eventChannel <- updatedTemplateEntity.GetUpdatedEvent()

	return updatedTemplateEntity.ToDomain(), err
}

// Delete removes a templateEntity by its ID.
func (r *templateEntityRepository) Delete(id string) error {
	err := r.BaseRepository.Delete(id)
	if err != nil {
		entity := entity.TemplateEntity{ID: id}
		eventChannel := worker_channels.GetCRUDEventChannel()
		eventChannel <- entity.GetUpdatedEvent()
	}
	return err
}
