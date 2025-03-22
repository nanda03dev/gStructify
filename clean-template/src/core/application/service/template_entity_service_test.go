package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type mockTemplateEntityRepository struct {
	mock.Mock
}

// BulkCreate implements repository.TemplateEntityRepository.
func (m *mockTemplateEntityRepository) BulkCreate(templateEntity []*aggregate.TemplateEntity) ([]*aggregate.TemplateEntity, error) {
	panic("unimplemented")
}

func (m *mockTemplateEntityRepository) Create(templateEntity *aggregate.TemplateEntity) (*aggregate.TemplateEntity, error) {
	args := m.Called(templateEntity)
	return args.Get(0).(*aggregate.TemplateEntity), args.Error(1)
}

func (m *mockTemplateEntityRepository) FindById(id string) (*aggregate.TemplateEntity, error) {
	args := m.Called(id)
	return args.Get(0).(*aggregate.TemplateEntity), args.Error(1)
}

func (m *mockTemplateEntityRepository) FindWithFilter(filter common.FilterQuery) ([]*aggregate.TemplateEntity, error) {
	args := m.Called(filter)
	return args.Get(0).([]*aggregate.TemplateEntity), args.Error(1)
}

func (m *mockTemplateEntityRepository) Update(templateEntity *aggregate.TemplateEntity) (*aggregate.TemplateEntity, error) {
	args := m.Called(templateEntity)
	return args.Get(0).(*aggregate.TemplateEntity), args.Error(1)
}

func (m *mockTemplateEntityRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateTemplateEntity(t *testing.T) {
	mockRepo := new(mockTemplateEntityRepository)
	service := NewTemplateEntityService(mockRepo)

	createDTO := dto.CreateTemplateEntityDTO{}
	templateEntity := &aggregate.TemplateEntity{ID: "1234567890", CreatedAt: time.Now(), UpdatedAt: time.Now()}

	mockRepo.On("Create", mock.Anything).Return(templateEntity, nil)

	createdTemplateEntity, err := service.Create(createDTO)
	assert.NoError(t, err)
	assert.Equal(t, templateEntity, createdTemplateEntity)
	mockRepo.AssertExpectations(t)
}

func TestGetById(t *testing.T) {
	mockRepo := new(mockTemplateEntityRepository)
	service := NewTemplateEntityService(mockRepo)

	templateEntity := &aggregate.TemplateEntity{ID: "123"}
	mockRepo.On("FindById", "123").Return(templateEntity, nil)

	foundTemplateEntity, err := service.GetById("123")
	assert.NoError(t, err)
	assert.Equal(t, templateEntity, foundTemplateEntity)
	mockRepo.AssertExpectations(t)
}

func TestFindWithFilter(t *testing.T) {
	mockRepo := new(mockTemplateEntityRepository)
	service := NewTemplateEntityService(mockRepo)

	filter := common.FilterQuery{}
	templateEntitys := []*aggregate.TemplateEntity{{ID: "123"}}

	mockRepo.On("FindWithFilter", filter).Return(templateEntitys, nil)

	foundTemplateEntitys, err := service.FindWithFilter(filter)
	assert.NoError(t, err)
	assert.Equal(t, templateEntitys, foundTemplateEntitys)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTemplateEntity(t *testing.T) {
	mockRepo := new(mockTemplateEntityRepository)
	service := NewTemplateEntityService(mockRepo)

	updateDTO := dto.UpdateTemplateEntityDTO{}
	templateEntity := &aggregate.TemplateEntity{ID: "123"}

	mockRepo.On("Update", mock.Anything).Return(templateEntity, nil)

	updatedTemplateEntity, err := service.Update("123", updateDTO)
	assert.NoError(t, err)
	assert.Equal(t, templateEntity, updatedTemplateEntity)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTemplateEntity(t *testing.T) {
	mockRepo := new(mockTemplateEntityRepository)
	service := NewTemplateEntityService(mockRepo)

	mockRepo.On("Delete", "123").Return(nil)

	err := service.Delete("123")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
