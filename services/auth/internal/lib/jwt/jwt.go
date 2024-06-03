package jwt

import (
	"auth/internal/domain/entities"
	"github.com/golang-jwt/jwt"
	"time"
)

func NewToken(user entities.User, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte("s7Ndh+pPznbHbS*+9Pk8qGWhTzbpa@tw"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
