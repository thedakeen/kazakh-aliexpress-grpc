package product

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/exp/slog"
	authv1 "product/gen/auth"
	"product/internal/domain/entities"
	"product/internal/lib/logger/sl"
	"product/pkg/storage"
)

type Product struct {
	log              *slog.Logger
	productProvider  ProductProvider //repository
	categoryProvider CategoryProvider
	authService      authv1.AuthServer
}

type CategoryProvider interface {
	GetAllCategories(ctx context.Context) ([]entities.Category, error)
}

type ProductProvider interface {
	GetProduct(ctx context.Context, productID string) (entities.Product, error)
	GetProductsByCategory(ctx context.Context, categoryName string) ([]entities.Product, error)
}

var (
	ErrNoCategories = errors.New("no categories found")
)

func New(
	log *slog.Logger,
	productProvider ProductProvider,
	categoryProvider CategoryProvider) *Product {
	return &Product{
		log:              log,
		productProvider:  productProvider,
		categoryProvider: categoryProvider,
	}
}

func (a *Product) Categories(ctx context.Context) ([]entities.Category, error) {
	const op = "product.Categories"

	log := a.log.With(
		slog.String("op", op),
	)

	log.Info("attempting to get all categories")

	categories, err := a.categoryProvider.GetAllCategories(ctx)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrNoRecordFound):
			log.Warn("no categories found", sl.Err(err))
			return nil, fmt.Errorf("%s:%w", err, ErrNoCategories)
		default:
			a.log.Error("failed to get categories", sl.Err(err))
			return nil, fmt.Errorf("%s:%w", op, err)
		}
	}

	log.Info("get all categories successfully")

	return categories, nil
}
