package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BolajiOlajide/go-api/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lirstname string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client
var ctx context.Context

func createPerson(response http.ResponseWriter, request http.Request) {
	response.Header().Add("content-type", "application/json")
	var person Person
	json.NewDecoder(request.Body).Decode(person)
	// collection = client.Database()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	fmt.Println("Starting the application...")
	client, ctx, err = database.InitializeDB()
	if err != nil {
		log.Fatal("Error initializing the database")
		return
	}

	router := mux.NewRouter()
	port := ":" + os.Getenv("PORT")
	http.ListenAndServe(port, router)
}
