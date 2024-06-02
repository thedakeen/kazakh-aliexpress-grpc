package repository

import (
	"auth_service/internal/domain/entities"
	"auth_service/pkg/storage"
	mongodb "auth_service/pkg/storage/mongo"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	client   *mongo.Client
	database *mongo.Database
}

func New(storagePath string) (*Storage, error) {
	const op = "repository.user.New"
	client, err := mongodb.GetClient()
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	database := client.Database("kazakh-aliexpress")
	return &Storage{client: client, database: database}, nil
}

func (s *Storage) SaveUser(ctx context.Context, email string, name string, passHash []byte) (string, error) {
	const op = "repository.user.SaveUser"

	existingUser := s.database.Collection("users").FindOne(ctx, bson.M{"email": email})
	if existingUser.Err() == nil {
		return "", fmt.Errorf("%s: %w", op, storage.ErrUserExists)
	}

	user := bson.M{
		"email":     email,
		"name":      name,
		"pass_hash": passHash,
		"role":      "buyer",
	}

	result, err := s.database.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (s *Storage) User(ctx context.Context, email string) (entities.User, error) {
	const op = "repository.user.User"

	var user entities.User

	err := s.database.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return entities.User{}, fmt.Errorf("%s:%w", op, storage.ErrNoRecordFound)
		default:
			return entities.User{}, fmt.Errorf("%s:%w", op, err)
		}
	}

	return user, nil
}

func (s *Storage) IsAdmin(ctx context.Context, userID string) (bool, error) {
	const op = "storage.sqlite.IsAdmin"

	return false, nil
}
