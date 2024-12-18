package repositories

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// BaseRepository defines CRUD operations for any entity type.
type BaseRepository[T any] struct {
	db *gorm.DB
}

// NewBaseRepository initializes a new BaseRepository.
func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

// Create inserts a new record into the database.
func (r *BaseRepository[T]) Create(entity *T) (*T, error) {
	if err := r.db.Create(entity).Error; err != nil {
		return nil, fmt.Errorf("failed to create record: %w", err)
	}
	return entity, nil
}

// FindById retrieves a record by its ID.
func (r *BaseRepository[T]) FindById(id string) (*T, error) {
	var entity T
	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("record not found")
		}
		return nil, fmt.Errorf("failed to find record by ID: %w", err)
	}
	return &entity, nil
}

// FindWithFilter retrieves records based on filters, sorting, limit, and skip for pagination
func (r *BaseRepository[T]) FindWithFilter(filterQueryDTO common.FilterQueryDTO) ([]*T, error) {
	var results []*T

	// Apply filters
	query := r.db.Model(new(T))
	for _, filter := range filterQueryDTO.Filters {
		switch filter.Operation {
		case common.OperationEQ:
			query = query.Where(fmt.Sprintf("%s = ?", filter.Key), filter.Value)
		case common.OperationIN:
			query = query.Where(fmt.Sprintf("%s IN (?)", filter.Key), filter.Value)
		case common.OperationGT:
			query = query.Where(fmt.Sprintf("%s > ?", filter.Key), filter.Value)
		case common.OperationLT:
			query = query.Where(fmt.Sprintf("%s < ?", filter.Key), filter.Value)
		case common.OperationGTE:
			query = query.Where(fmt.Sprintf("%s >= ?", filter.Key), filter.Value)
		case common.OperationLTE:
			query = query.Where(fmt.Sprintf("%s <= ?", filter.Key), filter.Value)
		case common.OperationNQ:
			query = query.Where(fmt.Sprintf("%s != ?", filter.Key), filter.Value)
		default:
			query = query.Where(fmt.Sprintf("%s = ?", filter.Key), filter.Value)
		}
	}

	// Apply sorting
	for _, order := range filterQueryDTO.Orders {
		// Check order type (ASC or DESC)
		if order.Type == common.OrderASC {
			query = query.Order(clause.OrderByColumn{
				Column: clause.Column{Name: order.Key},
				Desc:   false,
			})
		} else if order.Type == common.OrderDESC {
			query = query.Order(clause.OrderByColumn{
				Column: clause.Column{Name: order.Key},
				Desc:   true,
			})
		} else {
			return nil, fmt.Errorf("invalid order type %s", order.Type)
		}
	}

	// Apply pagination (limit and skip) with default values if not provided
	if filterQueryDTO.Limit > 0 {
		query = query.Limit(int(filterQueryDTO.Limit))
	} else {
		// Default to 10 if limit is not provided
		query = query.Limit(10)
	}
	if filterQueryDTO.OffSet > 0 {
		query = query.Offset(int(filterQueryDTO.OffSet))
	}

	// Execute the query
	if err := query.Find(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to find records with filters: %w", err)
	}

	return results, nil
}

// Update modifies an existing record in the database.
func (r *BaseRepository[T]) Update(entity *T) (*T, error) {
	if err := r.db.Save(entity).Error; err != nil {
		return nil, fmt.Errorf("failed to update record: %w", err)
	}
	return entity, nil
}

// Delete removes a record by its ID.
func (r *BaseRepository[T]) Delete(id string) error {
	var entity T
	if err := r.db.Delete(&entity, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete record: %w", err)
	}
	return nil
}
