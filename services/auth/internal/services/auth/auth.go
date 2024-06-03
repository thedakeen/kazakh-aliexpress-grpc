package auth

import (
	"auth_service/internal/domain/entities"
	"auth_service/internal/lib/jwt"
	"auth_service/internal/lib/logger/sl"
	"auth_service/pkg/storage"
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
	IsAdmin(ctx context.Context, userID string) (bool, error)
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

	//err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
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

func (a *Auth) IsAdmin(ctx context.Context, userID string) (bool, error) {
	const op = "auth.IsAdmin"

	log := a.log.With(
		slog.String("op", op),
		slog.String("user_id", userID),
	)

	log.Info("checking if user is admin")

	isAdmin, err := a.userProvider.IsAdmin(ctx, userID)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("checked if user is admin", slog.Bool("is_admin", isAdmin))

	return isAdmin, nil
}
