package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"product/internal/domain/entities"
	"product/pkg/storage"
	mongodb "product/pkg/storage/mongo"
)

type Storage struct {
	client   *mongo.Client
	database *mongo.Database
}

func New(storagePath string) (*Storage, error) {
	const op = "repository.product.New"
	client, err := mongodb.GetClient()

	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	database := client.Database("Qazaq-Aliexpress")
	return &Storage{client: client, database: database}, nil
}

func (s *Storage) GetAllCategories(ctx context.Context) ([]entities.Category, error) {
	const op = "repository.product.GetAllCategories"

	var categories []entities.Category

	opts := options.Find().SetSort(bson.D{{"category_name", 1}})

	cursor, err := s.database.Collection("categories").Find(ctx, bson.D{}, opts)

	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return []entities.Category{}, fmt.Errorf("%s:%w", op, storage.ErrNoRecordFound)
		default:
			return []entities.Category{}, fmt.Errorf("%s:%w", op, err)
		}
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			err = storage.ErrNoRecordFound
			return
		}
	}(cursor, ctx)

	for cursor.Next(ctx) {
		var category entities.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, fmt.Errorf("%s:%w", op, err)
		}
		categories = append(categories, category)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	if len(categories) == 0 {
		return nil, fmt.Errorf("%s: %w", op, storage.ErrNoRecordFound)
	}

	return categories, nil
}

func (s *Storage) GetProduct(ctx context.Context, productID string) (entities.Product, error) {
	return entities.Product{}, nil
}
func (s *Storage) GetProductsByCategory(ctx context.Context, categoryName string) ([]entities.Product, error) {
	return []entities.Product{}, nil
}
