package repository

import (
	"fmt"
	"strings"

	"github.com/nanda03dev/go-ms-template/src/common"
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

// Bulk inserts a new record into the database.
func (r *BaseRepository[T]) BulkCreate(entities []*T) ([]*T, error) {
	if err := r.db.Create(entities).Error; err != nil {
		return nil, fmt.Errorf("failed to create record: %w", err)
	}

	return entities, nil
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

// FindById retrieves a record by its ID.
func (r *BaseRepository[T]) FindByIdWithRelation(id string, preloadRelations []string) (*T, error) {
	var entity T
	query := r.db.Model(&entity)

	// Loop through the preloadRelations slice and apply Preload for each relation dynamically
	for _, relation := range preloadRelations {
		query = query.Preload(relation)
	}

	// Now execute the First query to fetch the entity with preloaded relations
	if err := query.First(&entity, "id = ?", id).Error; err != nil {
		// Handle error (e.g., record not found, or other DB errors)
		return nil, fmt.Errorf("failed to fetch entity with relations: %w", err)
	}

	// Return the entity with preloaded relations
	return &entity, nil
}

// FindWithFilter retrieves records based on filters, sorting, limit, and skip for pagination
func (r *BaseRepository[T]) FindWithFilter(filterQuery common.FilterQuery) ([]*T, error) {
	var results []*T

	// Fetch valid columns for the table corresponding to model T
	validColumns, err := GetValidColumnsForTable[T](r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve valid columns for table: %w", err)
	}

	// Apply filters
	query := r.db.Model(new(T))
	// Determine whether to use AND or OR
	operationFunc := query.Where // Default to AND
	if strings.ToUpper(filterQuery.Logic) == "OR" {
		operationFunc = query.Or
	}

	// Loop through filters and dynamically apply conditions
	for _, filter := range filterQuery.Conditions {

		if !validColumns[filter.Key] {
			return nil, fmt.Errorf("invalid column name: %s", filter.Key)
		}

		var condition string
		switch filter.Operator {
		case common.CONDITION_EQ:
			condition = fmt.Sprintf("%s = ?", filter.Key)
		case common.CONDITION_IN:
			condition = fmt.Sprintf("%s IN (?)", filter.Key)
		case common.CONDITION_GT:
			condition = fmt.Sprintf("%s > ?", filter.Key)
		case common.CONDITION_LT:
			condition = fmt.Sprintf("%s < ?", filter.Key)
		case common.CONDITION_GTE:
			condition = fmt.Sprintf("%s >= ?", filter.Key)
		case common.CONDITION_LTE:
			condition = fmt.Sprintf("%s <= ?", filter.Key)
		case common.CONDITION_NQ:
			condition = fmt.Sprintf("%s != ?", filter.Key)
		default:
			condition = fmt.Sprintf("%s = ?", filter.Key)
		}

		query = operationFunc(condition, filter.Value)
	}

	// Apply sorting
	for _, order := range filterQuery.Sorts {
		// Check order type (ASC or DESC)
		if order.Type == common.SORT_ASC {
			query = query.Order(clause.OrderByColumn{
				Column: clause.Column{Name: order.Key},
				Desc:   false,
			})
		} else if order.Type == common.SORT_DESC {
			query = query.Order(clause.OrderByColumn{
				Column: clause.Column{Name: order.Key},
				Desc:   true,
			})
		} else {
			return nil, fmt.Errorf("invalid order type %s", order.Type)
		}
	}

	// Apply pagination (limit and skip) with default values if not provided
	if filterQuery.MaxResults > 0 {
		query = query.Limit(int(filterQuery.MaxResults))
	} else {
		// Default to 10 if limit is not provided
		query = query.Limit(1000)
	}
	if filterQuery.Offset > 0 {
		query = query.Offset(int(filterQuery.Offset))
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

// getValidColumns retrieves valid columns for the specific table associated with the model
func GetValidColumnsForTable[T any](db *gorm.DB) (map[string]bool, error) {
	// Parse the model to get its schema
	stmt := &gorm.Statement{DB: db}
	err := stmt.Parse(new(T))
	if err != nil {
		return nil, err
	}

	// Retrieve column information only for the specific table
	columns := make(map[string]bool)
	for _, field := range stmt.Schema.Fields {
		columns[field.DBName] = true
	}

	return columns, nil
}
