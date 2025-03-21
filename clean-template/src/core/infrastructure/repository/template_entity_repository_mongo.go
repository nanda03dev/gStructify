package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TemplateEntityRepository interface {
	Create(templateEntity *aggregate.TemplateEntity) (*aggregate.TemplateEntity, error)
	BulkCreate(templateEntitys []*aggregate.TemplateEntity) ([]*aggregate.TemplateEntity, error)
	FindById(id string) (*aggregate.TemplateEntity, error)
	FindWithFilter(filterQuery common.FilterQuery) ([]*aggregate.TemplateEntity, error)
	Update(templateEntity *aggregate.TemplateEntity) (*aggregate.TemplateEntity, error)
	Delete(id string) error
}

// templateEntityRepository implements the TemplateEntityRepository interface using MongoBaseRepository.
type templateEntityRepository struct {
	*MongoBaseRepository[entity.TemplateEntity] // Embeds MongoBaseRepository for CRUD operations
}

// NewTemplateEntityRepository initializes a new templateEntityRepository instance.
func NewTemplateEntityRepository(databases *db.Databases) TemplateEntityRepository {
	return &templateEntityRepository{
		MongoBaseRepository: NewMongoBaseRepository[entity.TemplateEntity](databases.MongoDB.DB, "templateEntitys"), // MongoDB collection name
	}
}

// Create inserts a new templateEntity.
func (r *templateEntityRepository) Create(templateEntity *aggregate.TemplateEntity) (*aggregate.TemplateEntity, error) {
	entityTemplateEntity := entity.NewTemplateEntity(templateEntity)
	createdTemplateEntity, err := r.MongoBaseRepository.Create(context.TODO(), entityTemplateEntity)

	if err != nil {
		return nil, err
	}

	eventChannel := worker_channel.GetCRUDEventChannel()
	eventChannel <- entityTemplateEntity.GetCreatedEvent()

	return createdTemplateEntity.ToDomain(), nil
}

// BulkCreate inserts multiple templateEntitys at once.
func (r *templateEntityRepository) BulkCreate(aggregateList []*aggregate.TemplateEntity) ([]*aggregate.TemplateEntity, error) {
	var entityList = make([]*entity.TemplateEntity, 0, len(aggregateList))

	for _, each := range aggregateList {
		entityList = append(entityList, entity.NewTemplateEntity(each))
	}

	createdList, err := r.MongoBaseRepository.BulkCreate(context.TODO(), entityList)
	if err != nil {
		return nil, err
	}

	for _, each := range createdList {
		eventChannel := worker_channel.GetCRUDEventChannel()
		eventChannel <- each.GetCreatedEvent()
	}

	return aggregateList, nil
}

// FindById retrieves a templateEntity by its ID.
func (r *templateEntityRepository) FindById(id string) (*aggregate.TemplateEntity, error) {
	entityTemplateEntity, err := r.MongoBaseRepository.FindById(context.TODO(), id)
	if err != nil {
		return nil, err
	}

	return entityTemplateEntity.ToDomain(), nil
}

// FindWithFilter retrieves templateEntitys based on filter conditions.
func (r *templateEntityRepository) FindWithFilter(filterQuery common.FilterQuery) ([]*aggregate.TemplateEntity, error) {
	templateEntitys, err := r.MongoBaseRepository.FindWithFilter(context.TODO(), filterQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to find templateEntitys: %w", err)
	}

	// Convert entity templateEntitys to aggregate templateEntitys
	var result []*aggregate.TemplateEntity
	for _, templateEntity := range templateEntitys {
		result = append(result, templateEntity.ToDomain())
	}

	return result, nil
}

// Update modifies an existing templateEntity.
func (r *templateEntityRepository) Update(templateEntity *aggregate.TemplateEntity) (*aggregate.TemplateEntity, error) {
	entityTemplateEntity := entity.NewTemplateEntity(templateEntity)

	// Convert entityTemplateEntity struct to bson.M
	bsonData, err := bson.Marshal(entityTemplateEntity)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal entity to bson: %w", err)
	}

	var updateData bson.M
	err = bson.Unmarshal(bsonData, &updateData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal bson to bson.M: %w", err)
	}

	// Pass bson.M to the base repository update method
	updatedTemplateEntity, err := r.MongoBaseRepository.Update(context.TODO(), templateEntity.ID, updateData)
	if err != nil {
		return nil, err
	}

	// Send an event after updating
	eventChannel := worker_channel.GetCRUDEventChannel()
	eventChannel <- updatedTemplateEntity.GetUpdatedEvent()

	return updatedTemplateEntity.ToDomain(), nil
}

// Delete removes a templateEntity by its ID.
func (r *templateEntityRepository) Delete(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %w", err)
	}

	err = r.MongoBaseRepository.Delete(context.TODO(), id)
	if err == nil {
		entity := entity.TemplateEntity{ID: objectID}
		eventChannel := worker_channel.GetCRUDEventChannel()
		eventChannel <- entity.GetDeletedEvent()
	}
	return err
}
