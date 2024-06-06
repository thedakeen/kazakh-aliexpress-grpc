package product

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"os"
	productv1 "product/gen/product"
	_ "product/internal/clients/auth/grpc"
	authgrpc "product/internal/clients/auth/grpc"
	"product/internal/domain/entities"
	"product/internal/grpc/structs"
	"product/internal/lib/logger/sl"
	productService "product/internal/services/product"
	"strings"
	"time"
)

type Product interface {
	Categories(ctx context.Context) ([]entities.Category, error)
	GetProduct(ctx context.Context, productID string) (*entities.Product, error)
	GetProductsByCategory(ctx context.Context, categoryID string, limit int64, offset int64, sortOrder string) ([]*entities.Product, error)
	CreateCategory(ctx context.Context, category *entities.Category) (*entities.Category, error)
	UpdateCategory(ctx context.Context, categoryID string, categoryName string) (*entities.Category, error)
	DeleteCategory(ctx context.Context, categoryID string) (string, error)
}

type serverAPI struct {
	productv1.UnimplementedProductServer
	product Product
	v       *validator.Validate
	log     *slog.Logger
}

func Register(gRPC *grpc.Server, product Product, log *slog.Logger) {
	productv1.RegisterProductServer(gRPC, &serverAPI{
		product: product,
		v:       validator.New(),
		log:     log,
	})
}

func (s *serverAPI) authorizeAdmin(ctx context.Context) (jwt.MapClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "no metadata in request")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, status.Error(codes.Unauthenticated, "no authorization header")
	}

	tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")

	duration := 10 * time.Second

	authClient, err := authgrpc.New(
		context.Background(),
		s.log,
		"auth:50000",
		duration,
		3,
	)

	if err != nil {
		s.log.Error("failed to init auth client", sl.Err(err))
		os.Exit(1)
	}

	_, err = authClient.IsTokenValid(context.Background(), tokenString)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "invalid token claims")
	}

	role, ok := claims["role"].(string)
	if !ok || role != "admin" {
		return nil, status.Error(codes.PermissionDenied, "no access for this resource")
	}

	return claims, nil
}

func (s *serverAPI) Categories(ctx context.Context, req *productv1.CategoryRequest) (*productv1.CategoryResponse, error) {

	categoryRequest := structs.CategoryRequest{}

	err := s.v.Struct(categoryRequest)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	categories, err := s.product.Categories(ctx)
	if err != nil {
		switch {
		case errors.Is(err, productService.ErrNoCategories):
			return nil, status.Error(codes.NotFound, "no categories found")
		default:
			return nil, status.Error(codes.Internal, "failed to login")
		}
	}

	var categoryPointers []*productv1.Category
	for _, category := range categories {
		categoryPointer := &productv1.Category{
			Id:   category.ID.Hex(),
			Name: category.CategoryName,
		}
		categoryPointers = append(categoryPointers, categoryPointer)
	}

	return &productv1.CategoryResponse{
		Categories: categoryPointers,
	}, nil
}

func (s *serverAPI) GetProduct(ctx context.Context, req *productv1.ProductRequest) (*productv1.ProductResponse, error) {
	productRequest := structs.ProductRequest{
		ProductID: req.Id,
	}

	err := s.v.Struct(productRequest)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	product, err := s.product.GetProduct(ctx, req.GetId())
	if err != nil {
		switch {
		case errors.Is(err, productService.ErrNoProduct):
			return nil, status.Error(codes.InvalidArgument, "invalid product id")
		default:
			return nil, status.Error(codes.Internal, "failed to get product")
		}
	}

	var infosPointers []*productv1.ProductInfo
	for _, info := range product.Infos {
		infosPointers = append(infosPointers, &productv1.ProductInfo{
			InfoContent: info.Content,
			InfoTitle:   info.Title,
		})
	}

	var variantPointers []*productv1.ProductVariant
	for _, variant := range product.Variants {
		variantPointers = append(variantPointers, &productv1.ProductVariant{
			VariantTitle:   variant.Title,
			VariantOptions: variant.Variants,
		})
	}

	var categoryPointers []*productv1.Category
	for _, category := range product.Categories {
		categoryPointer := &productv1.Category{
			Id:   category.ID.Hex(),
			Name: category.CategoryName,
		}
		categoryPointers = append(categoryPointers, categoryPointer)
	}

	prod := &productv1.ProductEntry{
		Id:         product.ID.Hex(),
		Name:       product.ProductName,
		Price:      product.Price,
		Infos:      infosPointers,
		ImageUrls:  product.Images,
		Options:    variantPointers,
		Categories: categoryPointers,
	}

	return &productv1.ProductResponse{
		Product: prod,
	}, nil
}

func (s *serverAPI) GetProductsByCategory(ctx context.Context, req *productv1.ProductsByCategoryRequest) (*productv1.ProductsByCategoryResponse, error) {
	productsByCategoryRequest := structs.ProductsByCategoryRequest{
		CategoryID: req.CategoryId,
		Limit:      req.Limit,
		Offset:     req.Offset,
		SortOrder:  req.SortOrder,
	}

	err := s.v.Struct(productsByCategoryRequest)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	sortOrder := req.SortOrder
	if strings.TrimSpace(sortOrder) == "" {
		sortOrder = "asc"
	} else if sortOrder != "asc" && sortOrder != "desc" {
		return nil, status.Error(codes.InvalidArgument, "invalid sort order")
	}

	products, err := s.product.GetProductsByCategory(ctx, req.GetCategoryId(), req.Limit, req.Offset, sortOrder)
	if err != nil {
		switch {
		case errors.Is(err, productService.ErrNoProduct):
			return nil, status.Error(codes.InvalidArgument, "invalid category id")
		default:
			return nil, status.Error(codes.Internal, "failed to get products by category id")
		}
	}

	var productsPointers []*productv1.ProductEntry
	for _, product := range products {
		productsPointer := &productv1.ProductEntry{
			Id:        product.ID.Hex(),
			Name:      product.ProductName,
			Price:     product.Price,
			ImageUrls: product.Images,
		}
		productsPointers = append(productsPointers, productsPointer)
	}

	return &productv1.ProductsByCategoryResponse{
		Products: productsPointers,
	}, nil
}

func (s *serverAPI) CreateCategory(ctx context.Context, req *productv1.CreateCategoryRequest) (*productv1.CreateCategoryResponse, error) {
	//_, err := s.authorizeAdmin(ctx)
	//if err != nil {
	//	return nil, err
	//}

	createCategoryRequest := structs.CreateCategoryRequest{
		CategoryName: req.Category.Name,
	}

	err := s.v.Struct(createCategoryRequest)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	newCategory := entities.Category{
		CategoryName: req.Category.GetName(),
	}

	createdCategory, err := s.product.CreateCategory(ctx, &newCategory)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create category")
	}

	cat := &productv1.Category{
		Name: createdCategory.CategoryName,
	}

	return &productv1.CreateCategoryResponse{
		Category: cat,
	}, nil

}

func (s *serverAPI) UpdateCategory(ctx context.Context, req *productv1.UpdateCategoryRequest) (*productv1.UpdateCategoryResponse, error) {
	//_, err := s.authorizeAdmin(ctx)
	//if err != nil {
	//	return nil, err
	//}

	updateCategoryRequest := structs.UpdateCategoryRequest{
		CategoryName: req.Name,
		CategoryID:   req.Id,
	}

	err := s.v.Struct(updateCategoryRequest)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	updatedCategory, err := s.product.UpdateCategory(ctx, req.GetId(), req.GetName())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update category")
	}

	category := &productv1.Category{
		Id:   updatedCategory.ID.Hex(),
		Name: updatedCategory.CategoryName,
	}

	return &productv1.UpdateCategoryResponse{
		Category: category,
	}, nil

}

func (s *serverAPI) DeleteCategory(ctx context.Context, req *productv1.DeleteCategoryRequest) (*productv1.DeleteCategoryResponse, error) {
	deleteCategoryRequest := structs.DeleteCategoryRequest{
		CategoryID: req.Id,
	}

	err := s.v.Struct(deleteCategoryRequest)

	resp, err := s.product.DeleteCategory(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete category")
	}

	return &productv1.DeleteCategoryResponse{
		Response: resp,
	}, nil
}
