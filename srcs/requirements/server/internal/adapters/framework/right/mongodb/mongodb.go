package mongodb

import (
	"context"
	"data_impact/srcs/requirements/server/internal/domain/models"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	UserQuery() UserQuery
}

type repository struct{}

var (
	Client *mongo.Client
)

func NewRepository(client *mongo.Client) Repository {
	Client = client
	return &repository{}
}

func NewDB(dbUrl string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
	if err != nil {
		return nil, fmt.Errorf("client couldn't connect with context: %v", err)
	}

	// err = client.Ping(ctx, nil)
	// if err != nil {
	// 	return nil, fmt.Errorf("client couldn't ping database with context: %v", err)
	// }

	// TODO: use os.Getenv() for URI database name
	dataImpactDb := client.Database("dataimpactdb")
	userCollection = dataImpactDb.Collection(models.UserCollectionName)

	log.Println("successfully connected to MongoDB")

	return client, nil
}

func CloseDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := Client.Disconnect(ctx)
	if err != nil {
		log.Fatalf("db close failure: %v\n", err)
	}
}

func (r *repository) UserQuery() UserQuery {
	return &userQuery{}
}
