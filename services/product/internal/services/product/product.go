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
	GetProduct(ctx context.Context, productID string) (*entities.Product, error)
	GetProductsByCategory(ctx context.Context, categoryID string, limit int64, offset int64, sortOrder string) ([]*entities.Product, error)
}

var (
	ErrNoCategories = errors.New("no categories found")
	ErrNoProduct    = errors.New("no product found")
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

func (a *Product) GetProduct(ctx context.Context, productID string) (*entities.Product, error) {
	const op = "product.Categories"

	log := a.log.With(
		"op", op,
	)

	log.Info("attempting to get product")

	product, err := a.productProvider.GetProduct(ctx, productID)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrNoRecordFound):
			log.Warn("no product found", sl.Err(err))
			return nil, fmt.Errorf("%s:%w", err, ErrNoProduct)
		default:
			a.log.Error("failed to get product", sl.Err(err))
			return nil, fmt.Errorf("%s:%w", op, err)
		}
	}

	log.Info("get product successfully")

	return product, nil
}

func (a *Product) GetProductsByCategory(ctx context.Context, categoryID string, limit int64, offset int64, sortOrder string) ([]*entities.Product, error) {
	const op = "product.GetProductsByCategory"

	log := a.log.With(
		"op", op,
		"category", categoryID,
	)

	log.Info("attempting to get products by category")

	products, err := a.productProvider.GetProductsByCategory(ctx, categoryID, limit, offset, sortOrder)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrNoRecordFound):
			log.Warn("no products found", sl.Err(err))
			return nil, fmt.Errorf("%s:%w", err, ErrNoProduct)
		default:
			a.log.Error("failed to get products", sl.Err(err))
			return nil, fmt.Errorf("%s:%w", op, err)
		}
	}

	log.Info("get products by category successfully")

	return products, nil
}
