package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/BolajiOlajide/go-api/database"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreatePerson endpoint for creating a person
func CreatePerson(response http.ResponseWriter, request *http.Request) {
	collection, ctx := database.GetDB()
	response.Header().Add("content-type", "application/json")
	var person Person
	err := json.NewDecoder(request.Body).Decode(&person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", person)
	result, err := collection.InsertOne(ctx, person)
	if err != nil {
		log.Fatal(err, "Encountered error!")
		return
	}
	json.NewEncoder(response).Encode(result)
	log.Print("Done creating persoN!")
}

// GetPeople fetch everyone in the DB
func GetPeople(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var people []Person
	collection, ctx := database.GetDB()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

// GetPerson fetch a single person details
func GetPerson(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person Person

	collection, ctx := database.GetDB()
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

// Person structure for a person instance
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}
