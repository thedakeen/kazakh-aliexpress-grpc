package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *Storage) GetProduct(ctx context.Context, productID string) (*entities.Product, error) {
	const op = "repository.product.GetProduct"
	var product *entities.Product

	objID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, err
	}

	err = s.database.Collection("items").FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	return product, nil
}

func (s *Storage) GetProductsByCategory(ctx context.Context, categoryID string, limit int64, offset int64, sortOrder string) ([]*entities.Product, error) {
	const op = "repository.product.GetProductsByCategory"

	var products []*entities.Product

	objID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	filter := bson.D{
		{"categories", bson.D{
			{"$elemMatch", bson.D{{"_id", objID}}},
		}},
	}

	var sort string
	if sortOrder == "desc" {
		sort = "-item_name"
	} else {
		sort = "item_name"
	}

	opts := options.Find().SetSort(sort).SetLimit(limit).SetSkip(offset)

	cursor, err := s.database.Collection("products").Find(ctx, filter, opts)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, fmt.Errorf("%s:%w", op, storage.ErrNoRecordFound)
		default:
			return nil, fmt.Errorf("%s:%w", op, err)
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
		var product entities.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, fmt.Errorf("%s:%w", op, err)
		}
		products = append(products, &product)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("%s: %w", op, storage.ErrNoRecordFound)
	}

	return products, nil
}
