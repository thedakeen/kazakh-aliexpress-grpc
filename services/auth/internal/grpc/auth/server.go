package auth

import (
	authv1 "auth/gen/auth"
	"auth/internal/grpc/structs"
	"auth/internal/services/auth"
	"auth/pkg/storage"
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(ctx context.Context, email string, password string) (token string, err error)
	RegisterNewUser(ctx context.Context, email string, name string, password string) (userID string, err error)
}

type serverAPI struct {
	authv1.UnimplementedAuthServer
	auth Auth
	v    *validator.Validate
}

func Register(gRPC *grpc.Server, auth Auth) {
	authv1.RegisterAuthServer(gRPC, &serverAPI{
		auth: auth,
		v:    validator.New(),
	})
}

func (s *serverAPI) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	loginRequest := structs.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	err := s.v.Struct(loginRequest)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrInvalidCredentials):
			return nil, status.Error(codes.InvalidArgument, "invalid email or password")
		default:
			return nil, status.Error(codes.Internal, "failed to login")
		}
	}

	return &authv1.LoginResponse{
		Token: token,
	}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	registerRequest := structs.RegisterRequest{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	err := s.v.Struct(registerRequest)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userID, err := s.auth.RegisterNewUser(ctx, req.GetEmail(), req.GetName(), req.GetPassword())
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrUserExists):
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &authv1.RegisterResponse{
		UserId: userID,
	}, nil
}

//
//func (s *serverAPI) IsLoggedIn(ctx context.Context, req *authv1.IsLoggedInRequest) (*authv1.IsLoggedInResponse, error) {
//	isLoggedInRequest := structs.IsLoggedInRequest{
//		Token: req.Token,
//	}
//
//	err := s.v.Struct(isLoggedInRequest)
//	if err != nil {
//		return nil, status.Error(codes.InvalidArgument, err.Error())
//	}
//
//	isLoggedIn, err := s.auth.IsLoggedIn(ctx, req.GetToken())
//	if err != nil {
//		switch {
//		case errors.Is(err, storage.ErrNoRecordFound):
//			return nil, status.Error(codes.NotFound, "token not found")
//		default:
//			return nil, status.Error(codes.Internal, "internal server error")
//		}
//	}
//
//	return &authv1.IsLoggedInResponse{
//		IsLoggedIn: nil,
//	}, nil
//}

func (s *serverAPI) IsTokenValid(ctx context.Context, req *authv1.IsTokenValidRequest) (*authv1.IsTokenValidResponse, error) {
	const op = "auth.server.IsTokenValid"
	isTokenValidRequest := structs.IsTokenValidRequest{
		Token: req.Token,
	}

	err := s.v.Struct(isTokenValidRequest)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token, err := jwt.Parse(req.GetToken(), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s:%w", op, err)
		}
		// TODO: hide secret key
		return []byte("s7Ndh+pPznbHbS*+9Pk8qGWhTzbpa@tw"), nil
	})

	if err != nil {
		return &authv1.IsTokenValidResponse{TokenValid: false}, nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("Token is valid. Claims: %v\n", claims)
		return &authv1.IsTokenValidResponse{TokenValid: true}, nil
	}

	return &authv1.IsTokenValidResponse{
		TokenValid: false,
	}, nil
}
