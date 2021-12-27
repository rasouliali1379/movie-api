package mongodb

import (
	"context"
	"github.com/rasouliali1379/movie-api/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type userDb struct {
	db *mongo.Collection
}

func InitDatabase() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewUserDatabase(client *mongo.Client) (repository.IUserDb, error) {
	collection := client.Database("movie-db").Collection("users")

	return &userDb{db: collection}, nil
}

func DisconnectDatabase(client *mongo.Client, ctx context.Context) {
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatalf("Unable to disconnect from database: %s", err)
	}

}
