package product

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
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
	// TODO: products
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

//var (
//	ErrNoTokenInRequest = errors.New("no token in request")
//	JWTSecret           = []byte("s7Ndh+pPznbHbS*+9Pk8qGWhTzbpa@tw")
//)

//func ExtractAndVerifyToken(ctx context.Context) (*jwt.Token, error) {
//	md, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		return nil, ErrNoTokenInRequest
//	}
//
//	authHeader, ok := md["authorization"]
//	if !ok || len(authHeader) == 0 {
//		return nil, ErrNoTokenInRequest
//	}
//
//	tokenString := authHeader[0][7:] // "Bearer " prefix length is 7
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return JWTSecret, nil
//	})
//
//	if err != nil {
//		return nil, err
//	}
//
//	if !token.Valid {
//		return nil, errors.New("invalid token")
//	}
//
//	return token, nil
//}

func (s *serverAPI) Categories(ctx context.Context, req *productv1.CategoryRequest) (*productv1.CategoryResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "no metadata in request")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, status.Error(codes.Unauthenticated, "no authorization header")
	}

	token := authHeader[0][7:]

	authClient, err := authgrpc.New(
		context.Background(),
		s.log,
		"localhost:5000",
		time.Duration(400000),
		3)

	if err != nil {
		s.log.Error("failed to init auth client", sl.Err(err))
		os.Exit(1)
	}

	_, err = authClient.IsTokenValid(context.Background(), token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	categoryRequest := structs.CategoryRequest{}

	err = s.v.Struct(categoryRequest)

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
