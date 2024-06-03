package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Email          string             `bson:"email"`
	Name           string             `bson:"name"`
	HashedPassword []byte             `bson:"pass_hash"`
	Created        time.Time          `bson:"created"`
	Role           string             `bson:"role"`
}

func Matches(user User, password string) error {
	err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
	if err != nil {
		return err
	}

	return nil
}
