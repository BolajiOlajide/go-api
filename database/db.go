package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitializeDB method for initializing the database
func InitializeDB() (*mongo.Client, context.Context, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, _ := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_URL")))
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, nil, err
	}
	return client, ctx, nil
}
