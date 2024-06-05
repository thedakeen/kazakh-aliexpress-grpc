package product

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	productv1 "product/gen/product"
	_ "product/internal/clients/auth/grpc"
	"product/internal/domain/entities"
	"product/internal/grpc/structs"
	"product/internal/services/product"
)

type Product interface {
	Categories(ctx context.Context) ([]entities.Category, error)
	// TODO: products
}

type serverAPI struct {
	productv1.UnimplementedProductServer
	product Product
	v       *validator.Validate
}

func Register(gRPC *grpc.Server, product Product) {
	productv1.RegisterProductServer(gRPC, &serverAPI{
		product: product,
		v:       validator.New(),
	})
}

var (
	ErrNoTokenInRequest = errors.New("no token in request")
	JWTSecret           = []byte("s7Ndh+pPznbHbS*+9Pk8qGWhTzbpa@tw")
)

func ExtractAndVerifyToken(ctx context.Context) (*jwt.Token, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrNoTokenInRequest
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, ErrNoTokenInRequest
	}

	tokenString := authHeader[0][7:] // "Bearer " prefix length is 7
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func (s *serverAPI) Categories(ctx context.Context, req *productv1.CategoryRequest) (*productv1.CategoryResponse, error) {
	_, err := ExtractAndVerifyToken(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "unauthorized: "+err.Error())
	}

	categoryRequest := structs.CategoryRequest{}

	err = s.v.Struct(categoryRequest)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	categories, err := s.product.Categories(ctx)
	if err != nil {
		switch {
		case errors.Is(err, product.ErrNoCategories):
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
