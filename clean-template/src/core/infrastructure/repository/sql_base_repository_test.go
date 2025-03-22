package repository

import (
	"github.com/stretchr/testify/mock"
)

// MockSqlBaseRepository mocks SqlBaseRepository for Product
type MockSqlBaseRepository[T any] struct {
	mock.Mock
}

// Mock Create method for Product
func (m *MockSqlBaseRepository[T]) Create(product *T) (*T, error) {
	args := m.Called(product)

	// Corrected: Type assertion using a concrete type
	result, ok := args.Get(0).(*T)
	if !ok {
		return nil, args.Error(1)
	}

	return result, args.Error(1)
}

// Mock Create method for Product
func (m *MockSqlBaseRepository[T]) Update(product *T) (*T, error) {
	args := m.Called(product)

	// Corrected: Type assertion using a concrete type
	result, ok := args.Get(0).(*T)
	if !ok {
		return nil, args.Error(1)
	}

	return result, args.Error(1)
}

// Mock BulkCreate method
func (m *MockSqlBaseRepository[T]) BulkCreate(entities []*T) ([]*T, error) {
	args := m.Called(entities)

	// Type assertion
	result, ok := args.Get(0).([]*T)
	if !ok {
		return nil, args.Error(1)
	}

	return result, args.Error(1)
}

// Mock FindById method
func (m *MockSqlBaseRepository[T]) FindById(id string) (*T, error) {
	args := m.Called(id)

	// Type assertion
	result, ok := args.Get(0).(*T)
	if !ok {
		return nil, args.Error(1)
	}

	return result, args.Error(1)
}

// Mock FindByIdWithRelation method
func (m *MockSqlBaseRepository[T]) FindByIdWithRelation(id string, relations []string) (*T, error) {
	args := m.Called(id, relations)

	// Type assertion
	result, ok := args.Get(0).(*T)
	if !ok {
		return nil, args.Error(1)
	}

	return result, args.Error(1)
}

// Mock FindWithFilter method
func (m *MockSqlBaseRepository[T]) FindWithFilter(filter common.FilterQuery) ([]*T, error) {
	args := m.Called(filter)

	// Type assertion
	result, ok := args.Get(0).([]*T)
	if !ok {
		return nil, args.Error(1)
	}

	return result, args.Error(1)
}

// Mock Delete method
func (m *MockSqlBaseRepository[T]) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
