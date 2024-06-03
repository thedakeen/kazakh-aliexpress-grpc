package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	opt    *options.ClientOptions
)

func MustStart(uri string, opts ...*options.ClientOptions) {
	if len(opts) == 0 {
		opt = options.Client().ApplyURI(uri)
	} else {
		opt = opts[0]
	}

	err := connectDB()
	if err != nil {
		panic(err)
	}
}

func GetClient() (*mongo.Client, error) {
	if client != nil {
		return client, nil
	}

	err := connectDB()
	return client, err
}
func Close(ctx context.Context) (err error) {
	if client == nil {
		return
	}

	err = client.Disconnect(ctx)
	return
}

func connectDB() error {
	const op = "storage.mongo.connectDB"
	cli, err := mongo.Connect(context.Background(), opt)

	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}

	client = cli
	return err
}
