package repository

import (
	"github.com/stretchr/testify/mock"
)

// MockMongoBaseRepository mocks MongoBaseRepository for Product
type MockMongoBaseRepository[T any] struct {
	mock.Mock
}

// Mock Create method for Product
func (m *MockMongoBaseRepository[T]) Create(product *T) (*T, error) {
	args := m.Called(product)

	// Corrected: Type assertion using a concrete type
	result, ok := args.Get(0).(*T)
	if !ok {
		return nil, args.Error(1)
	}

	return result, args.Error(1)
}

// Mock Create method for Product
func (m *MockMongoBaseRepository[T]) Update(id string, product *T) (*T, error) {
	args := m.Called(product)

	// Corrected: Type assertion using a concrete type
	result, ok := args.Get(0).(*T)
	if !ok {
		return nil, args.Error(1)
	}

	return result, args.Error(1)
}

// Mock BulkCreate method
func (m *MockMongoBaseRepository[T]) BulkCreate(entities []*T) ([]*T, error) {
	args := m.Called(entities)

	// Type assertion
	result, ok := args.Get(0).([]*T)
	if !ok {
		return nil, args.Error(1)
	}

	return result, args.Error(1)
}

// Mock FindById method
func (m *MockMongoBaseRepository[T]) FindById(id string) (*T, error) {
	args := m.Called(id)

	// Type assertion
	result, ok := args.Get(0).(*T)
	if !ok {
		return nil, args.Error(1)
	}

	return result, args.Error(1)
}

// Mock FindWithFilter method
func (m *MockMongoBaseRepository[T]) FindWithFilter(filter common.FilterQuery) ([]*T, error) {
	args := m.Called(filter)

	// Type assertion
	result, ok := args.Get(0).([]*T)
	if !ok {
		return nil, args.Error(1)
	}

	return result, args.Error(1)
}

// Mock Delete method
func (m *MockMongoBaseRepository[T]) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
