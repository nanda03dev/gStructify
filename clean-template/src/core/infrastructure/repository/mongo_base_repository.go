package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoBaseRepository[T any] struct {
	collection *mongo.Collection
}

// NewMongoBaseRepository initializes a new MongoBaseRepository.
func NewMongoBaseRepository[T any](db *mongo.Database, collectionName string) *MongoBaseRepository[T] {
	return &MongoBaseRepository[T]{
		collection: db.Collection(collectionName),
	}
}

// Create inserts a new document into MongoDB.
func (r *MongoBaseRepository[T]) Create(ctx context.Context, entity *T) (*T, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, entity)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %w", err)
	}

	return entity, nil
}

// BulkCreate inserts multiple documents into MongoDB.
func (r *MongoBaseRepository[T]) BulkCreate(ctx context.Context, entities []*T) ([]*T, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var docs []interface{}
	for _, entity := range entities {
		docs = append(docs, entity)
	}

	_, err := r.collection.InsertMany(ctx, docs)
	if err != nil {
		return nil, fmt.Errorf("failed to insert multiple documents: %w", err)
	}
	return entities, nil
}

// FindById retrieves a document by its ID.
func (r *MongoBaseRepository[T]) FindById(ctx context.Context, id string) (*T, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var result T
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ObjectID format")
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("document not found")
		}
		return nil, fmt.Errorf("failed to find document by ID: %w", err)
	}
	return &result, nil
}

// FindWithFilter retrieves documents based on filters, sorting, limit, and offset.
func (r *MongoBaseRepository[T]) FindWithFilter(ctx context.Context, filterQuery common.FilterQuery) ([]*T, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	filter := bson.M{}
	for _, condition := range filterQuery.Conditions {
		switch condition.Operator {
		case common.CONDITION_EQ:
			filter[condition.Key] = condition.Value
		case common.CONDITION_IN:
			filter[condition.Key] = bson.M{"$in": condition.Value}
		case common.CONDITION_GT:
			filter[condition.Key] = bson.M{"$gt": condition.Value}
		case common.CONDITION_LT:
			filter[condition.Key] = bson.M{"$lt": condition.Value}
		case common.CONDITION_GTE:
			filter[condition.Key] = bson.M{"$gte": condition.Value}
		case common.CONDITION_LTE:
			filter[condition.Key] = bson.M{"$lte": condition.Value}
		case common.CONDITION_NQ:
			filter[condition.Key] = bson.M{"$ne": condition.Value}
		default:
			filter[condition.Key] = condition.Value
		}
	}
	fmt.Printf("\n filter %v", filter)

	opts := options.Find()
	if filterQuery.MaxResults > 0 {
		opts.SetLimit(int64(filterQuery.MaxResults))
	}
	if filterQuery.Offset > 0 {
		opts.SetSkip(int64(filterQuery.Offset))
	}

	sortOrder := bson.D{}
	for _, sort := range filterQuery.Sorts {
		order := 1
		if sort.Type == common.SORT_DESC {
			order = -1
		}
		sortOrder = append(sortOrder, bson.E{Key: sort.Key, Value: order})
	}
	opts.SetSort(sortOrder)

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find documents: %w", err)
	}
	defer cursor.Close(ctx)

	var results []*T
	if err := cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode documents: %w", err)
	}

	return results, nil
}

// Update modifies an existing document by its ID.
func (r *MongoBaseRepository[T]) Update(ctx context.Context, id string, updateData bson.M) (*T, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ObjectID format")
	}

	update := bson.M{"$set": updateData}
	result := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": objectID}, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		return nil, fmt.Errorf("failed to update document: %w", result.Err())
	}

	var updatedEntity T
	if err := result.Decode(&updatedEntity); err != nil {
		return nil, fmt.Errorf("failed to decode updated document: %w", err)
	}

	return &updatedEntity, nil
}

// Delete removes a document by its ID.
func (r *MongoBaseRepository[T]) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ObjectID format")
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return fmt.Errorf("failed to delete document: %w", err)
	}
	if result.DeletedCount == 0 {
		return errors.New("no document found to delete")
	}
	return nil
}
