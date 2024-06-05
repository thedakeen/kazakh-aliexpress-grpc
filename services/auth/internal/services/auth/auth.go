package auth

import (
	"auth/internal/domain/entities"
	"auth/internal/lib/jwt"
	"auth/internal/lib/logger/sl"
	"auth/pkg/storage"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slog"
	"time"
)

type Auth struct {
	log          *slog.Logger
	userSaver    UserSaver
	userProvider UserProvider
	tokenTTL     time.Duration
}

type UserSaver interface {
	SaveUser(ctx context.Context, email string, name string, passHash []byte) (userID string, err error)
}

type UserProvider interface {
	GetUser(ctx context.Context, email string) (entities.User, error)
	//IsLoggedIn(ctx context.Context, token string) (bool, error)
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
)

func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	tokenTTL time.Duration) *Auth {
	return &Auth{
		userSaver:    userSaver,
		userProvider: userProvider,
		log:          log,
		tokenTTL:     tokenTTL,
	}
}

func (a *Auth) Login(ctx context.Context, email string, password string) (string, error) {
	const op = "auth.Login"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("attempting to login user")

	user, err := a.userProvider.GetUser(ctx, email)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrNoRecordFound):
			log.Warn("user not found", sl.Err(err))
			return "", fmt.Errorf("%s:%w", op, ErrInvalidCredentials)
		default:
			a.log.Error("failed to get user", sl.Err(err))
			return "", fmt.Errorf("%s: %w", op, err)
		}
	}

	err = entities.Matches(user, password)
	if err != nil {
		log.Info("invalid credentials", sl.Err(err))
		return "", fmt.Errorf("%s:%w", op, ErrInvalidCredentials)
	}

	log.Info("logged in successfully")

	token, err := jwt.NewToken(user, a.tokenTTL)
	if err != nil {
		a.log.Error("failed to generate token", sl.Err(err))

		return "", fmt.Errorf("%s:%w", op, err)
	}

	return token, nil
}

func (a *Auth) RegisterNewUser(ctx context.Context, email string, name string, password string) (string, error) {
	const op = "auth.RegisterNewUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("registering user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	id, err := a.userSaver.SaveUser(ctx, email, name, passHash)
	if err != nil {
		log.Error("failed to save user", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

//func (a *Auth) IsLoggedIn(ctx context.Context, token string) (bool, error) {
//	const op = "auth.IsLoggedIn"
//
//	log := a.log.With(
//		slog.String("op", op),
//		slog.String("token", token),
//	)
//
//	log.Info("checking if user is logged in")
//
//	isLoggedIn, err := a.userProvider.IsLoggedIn(ctx, token)
//	if err != nil {
//		return false, fmt.Errorf("%s: %w", op, err)
//	}
//
//	log.Info("checked if user is logged in", slog.Bool("is_logged_in", isLoggedIn))
//
//	return isLoggedIn, nil
//}
