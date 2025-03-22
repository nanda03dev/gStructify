package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Test for TemplateEntityRepository Create method
func TestCreateTemplateEntity(t *testing.T) {
	mockRepo := new(MockMongoBaseRepository[entity.TemplateEntity])
	templateEntityRepo := &templateEntityRepository{baseRepo: mockRepo}

	objectId := primitive.NewObjectID()
	objectIdHex := objectId.Hex()

	// TODO: add other fields
	inputTemplateEntity := &aggregate.TemplateEntity{ID: objectIdHex}
	entityTemplateEntity := &entity.TemplateEntity{ID: objectId}

	// Define expected behavior of the mock
	mockRepo.On("Create", mock.Anything).Return(entityTemplateEntity, nil)

	// Call the actual method
	createdTemplateEntity, err := templateEntityRepo.Create(inputTemplateEntity)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, createdTemplateEntity)
	assert.Equal(t, inputTemplateEntity.ID, createdTemplateEntity.ID)
	assert.Equal(t, inputTemplateEntity.Name, createdTemplateEntity.Name)

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}

// Test FindById Method
func TestTemplateEntityRepository_FindById(t *testing.T) {
	mockRepo := new(MockMongoBaseRepository[entity.TemplateEntity])
	templateEntityRepo := &templateEntityRepository{baseRepo: mockRepo}

	objectId := primitive.NewObjectID()
	objectIdHex := objectId.Hex()

	// TODO: add other fields
	entityTemplateEntity := &entity.TemplateEntity{ID: objectId}
	aggregateTemplateEntity := entityTemplateEntity.ToDomain()

	// Mocking the repository call
	mockRepo.On("FindById", objectIdHex).Return(entityTemplateEntity, nil)

	foundTemplateEntity, err := templateEntityRepo.FindById(objectIdHex)

	assert.NoError(t, err)
	assert.Equal(t, aggregateTemplateEntity, foundTemplateEntity)

	mockRepo.AssertExpectations(t)
}

// Test FindWithFilter Method
func TestTemplateEntityRepository_FindWithFilter(t *testing.T) {
	mockRepo := new(MockMongoBaseRepository[entity.TemplateEntity])
	templateEntityRepo := &templateEntityRepository{baseRepo: mockRepo}

	objectId := primitive.NewObjectID()
	objectIdHex := objectId.Hex()

	filterQuery := common.FilterQuery{}
	// TODO: add other fields
	entityTemplateEntitys := []*entity.TemplateEntity{
		{ID: objectId},
	}

	mockRepo.On("FindWithFilter", filterQuery).Return(entityTemplateEntitys, nil)

	foundTemplateEntitys, err := templateEntityRepo.FindWithFilter(filterQuery)

	assert.NoError(t, err)
	assert.Len(t, foundTemplateEntitys, 1)
	assert.Equal(t, objectIdHex, foundTemplateEntitys[0].ID)

	mockRepo.AssertExpectations(t)
}

// Test Update Method
func TestTemplateEntityRepository_Update(t *testing.T) {
	mockRepo := new(MockMongoBaseRepository[entity.TemplateEntity])
	templateEntityRepo := &templateEntityRepository{baseRepo: mockRepo}

	objectId := primitive.NewObjectID()
	objectIdHex := objectId.Hex()

	// TODO: add other fields
	templateEntity := &aggregate.TemplateEntity{ID: objectIdHex}
	entityTemplateEntity := entity.NewTemplateEntity(templateEntity)

	mockRepo.On("Update", entityTemplateEntity).Return(entityTemplateEntity, nil)

	updatedTemplateEntity, err := templateEntityRepo.Update(templateEntity)

	assert.NoError(t, err)
	assert.Equal(t, templateEntity.ID, updatedTemplateEntity.ID)

	mockRepo.AssertExpectations(t)
}

// Test Delete Method
func TestTemplateEntityRepository_Delete(t *testing.T) {
	mockRepo := new(MockMongoBaseRepository[entity.TemplateEntity])
	templateEntityRepo := &templateEntityRepository{baseRepo: mockRepo}

	objectId := primitive.NewObjectID()
	objectIdHex := objectId.Hex()

	mockRepo.On("Delete", objectIdHex).Return(nil)

	err := templateEntityRepo.Delete(objectIdHex)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
