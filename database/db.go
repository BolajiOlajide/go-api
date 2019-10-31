package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx context.Context
var collection *mongo.Collection

// InitializeDB method for initializing the database
func InitializeDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, _ := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_URL")))
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal("Error initializing DB!")
		return
	}
	collection = client.Database("hapi-auth").Collection("people")
	log.Println("Successfully connected to DB!")
}

// GetDB get db collection
func GetDB() (*mongo.Collection, context.Context) {
	return collection, ctx
}
