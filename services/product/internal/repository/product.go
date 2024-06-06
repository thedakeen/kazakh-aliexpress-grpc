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

	filter := bson.D{
		{"$match", bson.D{
			{"categories", bson.D{
				{"$elemMatch", bson.D{{"_id", categoryID}}},
			}},
		}},
	}

	sortStage := bson.D{}
	if sortOrder == "desc" {
		sortStage = bson.D{{"$sort", bson.D{{"item_name", -1}}}}
	} else {
		sortStage = bson.D{{"$sort", bson.D{{"item_name", 1}}}}
	}

	skipStage := bson.D{{"$skip", offset}}
	limitStage := bson.D{{"$limit", limit}}

	cursor, err := s.database.Collection("items").Aggregate(ctx, mongo.Pipeline{filter, sortStage, skipStage, limitStage})
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

	var products []*entities.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, storage.ErrNoRecordFound
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("%s: %w", op, storage.ErrNoRecordFound)
	}

	return products, nil

}

//////////////////////// ADMIN //////////////////////////////

func (s *Storage) CreateCategory(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	const op = "repository.product.CreateCategory"

	existingUser := s.database.Collection("categories").FindOne(ctx, bson.M{"category_name": category.CategoryName})
	if existingUser.Err() == nil {
		return nil, fmt.Errorf("%s: %w", op, storage.ErrCategoryExists)
	}

	_, err := s.database.Collection("categories").InsertOne(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	return category, nil
}

func (s *Storage) UpdateCategory(ctx context.Context, categoryID string, categoryName string) (*entities.Category, error) {
	const op = "repository.category.UpdateCategory"

	objID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "category_name", Value: categoryName},
		}},
	}

	var updatedCategory entities.Category
	err = s.database.Collection("categories").FindOneAndUpdate(ctx, filter, update).Decode(&updatedCategory)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, fmt.Errorf("%s:%w", op, storage.ErrNoRecordFound)
		default:
			return nil, fmt.Errorf("%s:%w", op, err)
		}
	}

	return &updatedCategory, nil
}

func (s *Storage) DeleteCategory(ctx context.Context, categoryID string) (string, error) {
	const op = "repository.category.DeleteCategory"

	objID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return "", fmt.Errorf("%s:%w", op, err)
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := s.database.Collection("categories").DeleteOne(ctx, filter)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return "", fmt.Errorf("%s:%w", op, storage.ErrNoRecordFound)
		default:
			return "", fmt.Errorf("%s:%w", op, err)
		}
	}

	if result.DeletedCount == 0 {
		return "", fmt.Errorf("%s:%w", op, storage.ErrNoRecordFound)
	}

	return "deleted successfully", nil

}
